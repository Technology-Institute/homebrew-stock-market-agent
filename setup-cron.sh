#!/bin/bash
# Example crontab entries for the stock market agent
# Edit your crontab with: crontab -e

# Get the full path to the agent
AGENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "Stock Market Agent - Cron Setup"
echo "================================"
echo ""
echo "Agent directory: $AGENT_DIR"
echo ""
echo "Example crontab entries (adjust times based on your timezone):"
echo ""
echo "# Run at 9:30 AM and 4:00 PM EST (14:30 and 21:00 UTC) on weekdays"
echo "30 14,21 * * 1-5 cd $AGENT_DIR && ./stock-market-agent >> $AGENT_DIR/agent.log 2>&1"
echo ""
echo "# Run every hour during market hours (9 AM - 5 PM EST = 14:00 - 22:00 UTC)"
echo "0 14-22 * * 1-5 cd $AGENT_DIR && ./stock-market-agent >> $AGENT_DIR/agent.log 2>&1"
echo ""
echo "# Run daily at 4:30 PM EST (21:30 UTC) with market close summary"
echo "30 21 * * 1-5 cd $AGENT_DIR && ./stock-market-agent >> $AGENT_DIR/agent.log 2>&1"
echo ""
echo "To install, run: crontab -e"
echo "Then paste one of the lines above (adjust timezone as needed)"
