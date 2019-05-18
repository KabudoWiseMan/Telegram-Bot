package main

import (
	"log"
	"github.com/Syfaro/telegram-bot-api"
	"math/rand"
	"time"
	"html"
	"strconv"
	"io/ioutil"
	"net/http"
	"os"
)

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
    resp.Write([]byte("Hi there! I'm Sending Love to Alin bot!"))
}

func emoji(i int) string {
	if i != 0 {
		return html.UnescapeString("&#" + strconv.Itoa(i) + ";")
	}
	r := random(len(emojiA))
	return html.UnescapeString("&#" + strconv.Itoa(emojiA[r]) + ";")
}

func random(l int) int {
	rand.Seed(time.Now().Unix())
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return r.Intn(l)
}

func love() string {
	r := random(len(iLikeYour))
	return iLikeYour[r]
}

func poem() string {
	r := random(len(poems))
	return poems[r]
}

func my() string {
	r := random(len(mys))
	return mys[r]
}

func chooseFile(s string) (url string) {
	switch s {
		case "photo":
			files, err := ioutil.ReadDir("photos")
			if err != nil {
				log.Fatal(err)
			}
			r := random(len(files))
			if r == 0 {
				r++
			}
			url = "photos/" + files[r].Name()
		case "audio":
			files, err := ioutil.ReadDir("audios")
			if err != nil {
				log.Fatal(err)
			}
			r := random(len(files))
			if r == 0 {
				r++
			}
			url = "audios/" + files[r].Name()
		default:
			log.Panic()
	}
	return
}

func main() {
	bot, err := tgbotapi.NewBotAPI("663486910:AAFnS81mK2a_************************")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.ListenForWebhook("/" + bot.Token)

	http.HandleFunc("/", MainHandler)
    go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		UserName := update.Message.From.UserName
		ChatID := update.Message.Chat.ID
		if (ChatID != 5867**** && ChatID != 49467****) {
			msg := tgbotapi.NewMessage(ChatID, "Вы нежелательный гость, пожалуйста, покиньте этого бота, о вашем визите будет доложено")
			bot.Send(msg)
			msgToHost := tgbotapi.NewMessage(5867****, "К боту обратился пользователь: " + UserName + "\nChatID: " + strconv.FormatInt(ChatID, 10) + "\nCообщение: " + update.Message.Text)
			bot.Send(msgToHost)
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(ChatID, "")
			switch update.Message.Command() {
				case "start":
					msg.Text = "Привет" + emoji(128522) + " я бот, который может говорить тебе приятности от Севы." + emoji(10084) + " Просто отправь любое сообщение, чтобы получить заряд позитива." + emoji(128536) + " Мои дополнительные возможности можно посмотреть, набрав /help"
				case "help":
					msg.Text = "Набери:\n\n" + 
					"/poem — получить стихотворение\n" +
					"/emoji — получить много emoji\n" +
					"/photo — получить фото\n" +
					"/song — получить песню в моём исполнении\n" +
					"/your — получить мой стих или мою фразу"
				case "poem":
					msg.Text = poem()
				case "emoji":
					msg.Text = ""
					for i := 0; i <= random(1200); i++ {
						msg.Text += emoji(0)
					}
				case "photo":
					file := chooseFile("photo")
					photo := tgbotapi.NewPhotoUpload(ChatID, file)
					bot.Send(photo)
					continue
				case "song":
					file := chooseFile("audio")
					audio := tgbotapi.NewAudioUpload(ChatID, file)
					bot.Send(audio)
					continue
				case "your":
					msg.Text = my()
				default:
					msg.Text = "Я не знаю такой команды" + emoji(128517) + " Попроси Севу добавить её" + emoji(128522)
			}
			bot.Send(msg)
		} else {
			Text := update.Message.Text
			log.Printf("[%s] %d %s", UserName, ChatID, Text)
			reply := love()
			for i := 0; i <= random(6); i++ {
				reply += emoji(0)
			}
			msg := tgbotapi.NewMessage(ChatID, reply)
			bot.Send(msg)
		}
	}
}