# Stock Market Telegram Agent

A Go-based AI agent that sends stock market updates to a Telegram channel using the [Fantasy](https://github.com/charmbracelet/fantasy) framework.

## Features

- Fetches real-time stock market data for major indices (S&P 500, NASDAQ, DOW)
- Uses AI to generate concise, informative market updates
- Sends formatted messages to Telegram
- Configurable via YAML config file
- Designed to run as a cron job

## Setup

1. Copy the sample environment file:
   ```bash
   cp .sample-env .env
   ```

2. Edit `.env` and add your API keys:
   - `ANTHROPIC_API_KEY`: Your Anthropic API key from [console.anthropic.com](https://console.anthropic.com/)
   - `TELEGRAM_BOT_TOKEN`: Your Telegram bot token from [@BotFather](https://t.me/botfather)
   - `TELEGRAM_CHAT_ID`: Your Telegram chat/channel ID
   - `ALPHAVANTAGE_API_KEY`: Alpha Vantage API key for real stock data

   **Note:** Make sure your Anthropic API key has access to Claude models and sufficient credits.

3. Configure the agent behavior in `config.yaml`:
   - System prompt
   - Model selection
   - Schedule (for cron)
   - Telegram formatting options

4. Build the agent:
   ```bash
   go build -o stock-market-agent
   ```

## Usage

### Run manually:
```bash
./stock-market-agent
```

### Run as a cron job:
Add to your crontab (edit with `crontab -e`):
```bash
# Run at 9:30 AM and 4:00 PM EST on weekdays (market open/close)
30 14,21 * * 1-5 cd /home/bjoern/Github/go-agent && ./stock-market-agent >> agent.log 2>&1
```

Note: Adjust the times based on your server's timezone. The example above assumes UTC.

## Configuration

### config.yaml

- `agent.system_prompt`: Instructions for the AI on how to format market updates
- `agent.model`: The AI model to use (default: anthropic/claude-3.5-sonnet)
- `agent.provider`: The provider (default: openrouter)
- `agent.schedule.cron`: Cron expression for scheduling
- `agent.telegram.parse_mode`: Telegram formatting (Markdown, HTML, or empty)

### Stock Data API

By default, the agent uses simulated data. For real market data, get a free API key from:
- [Alpha Vantage](https://www.alphavantage.co/support/#api-key)
- [Finnhub](https://finnhub.io/)
- [Yahoo Finance](https://www.yahoofinanceapi.com/)

Add the key to your `.env` file as `STOCK_API_KEY`.

## Getting Telegram Credentials

1. **Bot Token**: Talk to [@BotFather](https://t.me/botfather) on Telegram:
   - Send `/newbot`
   - Follow the prompts
   - Copy the token

2. **Chat ID**:
   - For personal messages: Use [@userinfobot](https://t.me/userinfobot)
   - For channels: Add bot as admin, then use [@getidsbot](https://t.me/getidsbot)

## License

MIT
