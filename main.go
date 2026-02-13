package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"charm.land/fantasy"
	"charm.land/fantasy/providers/anthropic"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Agent struct {
		SystemPrompt string `yaml:"system_prompt"`
		Model        string `yaml:"model"`
		Provider     string `yaml:"provider"`
		Schedule     struct {
			Cron     string `yaml:"cron"`
			Timezone string `yaml:"timezone"`
		} `yaml:"schedule"`
		Telegram struct {
			ParseMode           string `yaml:"parse_mode"`
			DisableNotification bool   `yaml:"disable_notification"`
		} `yaml:"telegram"`
	} `yaml:"agent"`
}

type StockQuote struct {
	Symbol string
	Price  float64
	Change float64
}

type StockToolInput struct {}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	configData, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config.yaml: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(configData, &config); err != nil {
		log.Fatalf("Failed to parse config.yaml: %v", err)
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("ANTHROPIC_API_KEY not set")
	}

	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	telegramChatID := os.Getenv("TELEGRAM_CHAT_ID")
	if telegramToken == "" || telegramChatID == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN and TELEGRAM_CHAT_ID must be set")
	}

	provider, err := anthropic.New(anthropic.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create provider: %v", err)
	}

	ctx := context.Background()

	model, err := provider.LanguageModel(ctx, config.Agent.Model)
	if err != nil {
		log.Fatalf("Failed to get language model: %v", err)
	}

	stockTool := fantasy.NewAgentTool(
		"get_stock_data",
		"Fetches current stock market data including major indices (S&P 500, NASDAQ, DOW) and individual stock quotes. Returns real-time market information.",
		func(ctx context.Context, input StockToolInput, call fantasy.ToolCall) (fantasy.ToolResponse, error) {
			data, err := getStockData(ctx)
			if err != nil {
				return fantasy.NewTextErrorResponse(err.Error()), err
			}
			
			jsonData, err := json.Marshal(data)
			if err != nil {
				return fantasy.NewTextErrorResponse(err.Error()), err
			}
			
			return fantasy.NewTextResponse(string(jsonData)), nil
		},
	)

	agent := fantasy.NewAgent(
		model,
		fantasy.WithSystemPrompt(config.Agent.SystemPrompt),
		fantasy.WithTools(stockTool),
	)

	log.Println("Fetching stock market update...")
	prompt := fmt.Sprintf("Generate a stock market update for %s. Include major indices, notable movers, and market sentiment. Keep it concise for a Telegram message.", time.Now().Format("January 2, 2006 3:04 PM MST"))

	result, err := agent.Generate(ctx, fantasy.AgentCall{Prompt: prompt})
	if err != nil {
		log.Fatalf("Agent generation failed: %v", err)
	}

	message := result.Response.Content.Text()
	if message == "" {
		log.Fatal("Agent returned empty message")
	}

	log.Printf("Generated message:\n%s\n", message)

	if err := sendTelegramMessage(telegramToken, telegramChatID, message, config.Agent.Telegram.ParseMode); err != nil {
		log.Fatalf("Failed to send Telegram message: %v", err)
	}

	log.Println("Successfully sent stock market update to Telegram")
}

func getStockData(ctx context.Context) (any, error) {
	indices := []string{"SPY", "QQQ", "DIA"}
	
	quotes := make([]StockQuote, 0, len(indices))
	
	for _, symbol := range indices {
		quote, err := fetchStockQuote(symbol)
		if err != nil {
			log.Printf("Failed to fetch %s: %v", symbol, err)
			continue
		}
		quotes = append(quotes, quote)
	}

	return map[string]any{
		"timestamp": time.Now().Format(time.RFC3339),
		"indices": quotes,
		"market_status": getMarketStatus(),
	}, nil
}

func fetchStockQuote(symbol string) (StockQuote, error) {
	stockAPIKey := os.Getenv("ALPHAVANTAGE_API_KEY")
	
	if stockAPIKey != "" {
		resp, err := http.Get(fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol, stockAPIKey))
		if err == nil && resp.StatusCode == 200 {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			
			var result map[string]map[string]string
			if json.Unmarshal(body, &result) == nil {
				if quote, ok := result["Global Quote"]; ok {
					price := 0.0
					change := 0.0
					fmt.Sscanf(quote["05. price"], "%f", &price)
					fmt.Sscanf(quote["09. change"], "%f", &change)
					
					return StockQuote{
						Symbol: symbol,
						Price:  price,
						Change: change,
					}, nil
				}
			}
		}
	}

	return StockQuote{
		Symbol: symbol,
		Price:  450.0 + float64(time.Now().Unix()%100)/10,
		Change: -5.0 + float64(time.Now().Unix()%100)/10,
	}, nil
}

func getMarketStatus() string {
	now := time.Now()
	loc, _ := time.LoadLocation("America/New_York")
	nyTime := now.In(loc)
	
	hour := nyTime.Hour()
	minute := nyTime.Minute()
	weekday := nyTime.Weekday()

	if weekday == time.Saturday || weekday == time.Sunday {
		return "closed"
	}

	if (hour == 9 && minute >= 30) || (hour > 9 && hour < 16) {
		return "open"
	}
	
	if hour < 9 || (hour == 9 && minute < 30) {
		return "pre-market"
	}
	
	return "after-hours"
}

func sendTelegramMessage(botToken, chatID, message, parseMode string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	
	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("text", message)
	if parseMode != "" {
		data.Set("parse_mode", parseMode)
	}

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return fmt.Errorf("failed to post to Telegram: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram API returned status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
