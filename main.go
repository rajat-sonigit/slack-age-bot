package main
import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func printCommandEvents(analyticChannel <-chan *slacker.CommandEvent) {
	for event := range analyticChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6847844354790-6878255606512-qaNCdjr8k6ptMVL7OrzIHwHh")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06RU5EGR24-6855405371410-3b9aa989be7702a58644939fe7098dd8e571f362b0eec0ec6080c99bd5394d3a")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2024 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
		Interactive: func(slacker.InteractiveBotContext, *socketmode.Request, *slack.InteractionCallback) {
		},
		HideHelp: false,
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
