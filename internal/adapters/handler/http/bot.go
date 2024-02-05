package adapters

import (
	"fmt"
	_ "holdem/docs"
	service "huskyholdem/service"
	"huskyholdem/utils"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TODO: Move this to random generator in utils
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// TODO: Move this to random generator in utils
func randomString(length int) string {
	test := uuid.New().String()
	sum := 0
	for i := 0; i < len(test); i++ {
		sum += int(test[i])
	}
	rand.Seed(int64(sum))
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type BotHandler struct {
	BotService  *service.BotService
	UserService *service.UserService
}

func NewBotHandler(botService *service.BotService, userService *service.UserService) *BotHandler {
	return &BotHandler{
		BotService:  botService,
		UserService: userService,
	}
}

func (h *BotHandler) CreateNewBot(c *gin.Context) {

	type messagePayload struct {
		Name   string `json:"name"`
		ImgUrl string `json:"imgUrl"`
		Email  string `json:"email"`
	}

	var botInfo messagePayload

	if err := c.ShouldBindJSON(&botInfo); err != nil {
		fmt.Println("error: ", err)
		fmt.Println("error: ", botInfo)
		utils.HandleError(c, utils.ErrBadRequest)
		return
	}

	if botInfo.Name == "" {
		botInfo.Name = "Bot-" + randomString(6)
	}

	if botInfo.ImgUrl == "" {
		botInfo.ImgUrl = "https://images.wagwalkingweb.com/media/daily_wag/blog_articles/hero/1685787498.877709/fun-facts-about-siberian-huskies-1.png"
	}

	id, err := h.BotService.CreateNewBot(botInfo.Name, botInfo.ImgUrl, botInfo.Email)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleSuccess(c, gin.H{"bot_id": id})
}

func UpdateBotMetadata(c *gin.Context) {
	// var botInfo bot.BotMetaData
}
