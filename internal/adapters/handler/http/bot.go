package adapters

import (
	"fmt"
	_ "holdem/docs"
	"huskyholdem/bot"
	service "huskyholdem/service"
	"huskyholdem/utils"
	"math/rand"
	"strings"

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

// CreateNewBot godoc
// @Summary		Create new bot
// @Description	Create a new bot with name, imgUrl, and email
// @Authentication	Bearer AuthToken
// @Tags			Bot
// @Accept		json
// @Produce		json
// @Param		name	body	string	false	"Bot Name"
// @Param		imgUrl	body	string	false	"Bot Image URL"
// @Param		email	body	string	false	"Bot Email"
// @Success		200			{object}	utils.responseMessage	"Bot ID"
// @Failure		400			{object}	utils.errorResponseMessage	"Bad request"
// @Router		/bot/new [post]
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
		botInfo.Name = "Bot-" + randomString(16)
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

// UpdateBotMetadata godoc
// @Summary		Update bot metadata
// @Description	Update bot metadata with name and imgUrl
// @Authentication	Bearer AuthToken
// @Tags			Bot
// @Accept		json
// @Produce		json
// @Param		botId	path	string	true	"Bot ID"
// @Param		name	body	string	false	"Bot Name"
// @Param		imgUrl	body	string	false	"Bot Image URL"
// @Success		200			{object}	utils.responseMessage	"Bot Metadata"
// @Failure		400			{object}	utils.errorResponseMessage	"Bad request"
// @Router		/bot/{botId} [put]
func (h *BotHandler) UpdateBotMetadata(c *gin.Context) {
	var botInfo bot.BotMetaData

	if err := c.ShouldBindJSON(&botInfo); err != nil {
		utils.HandleError(c, utils.ErrBadRequest)
		return
	}

	bot_id := c.Param("botId")

	bot, err_find_bot := h.BotService.GetBotByID(bot_id)

	if err_find_bot != nil {
		utils.HandleError(c, err_find_bot)
		return
	}

	if botInfo.Name == "" {
		botInfo.Name = bot.Name
	}

	if botInfo.ImgUrl == "" {
		botInfo.ImgUrl = bot.ImgUrl
	}

	err := h.BotService.UpdateBotMetadata(bot_id, &botInfo)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleSuccess(c, botInfo)
}

// GetBotByID godoc
// @Summary		Get bot by ID
// @Description	Get bot by ID
// @Authentication	Bearer AuthToken
// @Tags			Bot
// @Accept		json
// @Produce		json
// @Param		botId	path	string	true	"Bot ID"
// @Success		200			{object}	bot.Bot	"Bot"
// @Failure		400			{object}	utils.errorResponseMessage	"Bad request"
// @Router		/bot/{botId} [get]
func (h *BotHandler) GetBotByID(c *gin.Context) {
	bot_id := c.Param("botId")

	bot, err := h.BotService.GetBotByID(bot_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleSuccess(c, bot)
}

// GenerateBotAuthToken godoc
// @Summary		Generate bot auth token
// @Description	Generate bot auth token
// @Authentication	Bearer AuthToken
// @Tags			Bot
// @Accept		json
// @Produce		json
// @Param		botId	path	string	true	"Bot ID"
// @Success		200			{object}	utils.responseMessage	"Bot Auth Token"
// @Failure		400			{object}	utils.errorResponseMessage	"Bad request"
// @Router		/bot/{botId}/generate [post]
func (h *BotHandler) GenerateBotAuthToken(c *gin.Context) {
	bot_id := c.Param("botId")

	token, err := h.BotService.GenerateBotAuthToken(bot_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleSuccess(c, token)
}

// GetBotByToken godoc
// @Summary		Get bot by token
// @Description	Get bot by token
// @Authentication	Bearer BotToken
// @Tags			Bot
// @Accept		json
// @Produce		json
// @Success		200			{object}	bot.Bot	"Bot"
// @Failure		400			{object}	utils.errorResponseMessage	"Bad request"
// @Router		/bot/token/self [get]
func (h *BotHandler) GetBotByToken(c *gin.Context) {

	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	botClaims, _ := service.ParseBotToken(token)

	bot, err := h.BotService.GetBotByID(botClaims.BotId)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleSuccess(c, bot)
}

// UpdateBotMetadataByToken godoc
// @Summary		Update bot metadata by token
// @Description	Update bot metadata by token
// @Authentication	Bearer BotToken
// @Tags			Bot
// @Accept		json
// @Produce		json
// @Param		name	body	string	false	"Bot Name"
// @Param		imgUrl	body	string	false	"Bot Image URL"
// @Success		200			{object}	bot.BotMetaData	"Bot Metadata"
// @Failure		400			{object}	utils.errorResponseMessage	"Bad request"
// @Router		/bot/token/update [put]
func (h *BotHandler) UpdateBotMetadataByToken(c *gin.Context) {
	var botInfo bot.BotMetaData

	if err := c.ShouldBindJSON(&botInfo); err != nil {
		utils.HandleError(c, utils.ErrBadRequest)
		return
	}

	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	botClaims, _ := service.ParseBotToken(token)

	err := h.BotService.UpdateBotMetadata(botClaims.BotId, &botInfo)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleSuccess(c, botInfo)
}

// AddKeyValueToCache godoc
// @Summary		Add key value to cache
// @Description	Add key value to cache
// @Authentication	Bearer BotToken
// @Tags			Bot
// @Accept		json
// @Produce		json
// @Param		key	body	string	true	"Key"
// @Param		value	body	string	true	"Value"
// @Success		200			{object}	utils.responseMessage	"Key Value"
// @Failure		400			{object}	utils.errorResponseMessage	"Bad request"
// @Router		/bot/token/key [post]
func (h *BotHandler) AddKeyValueToCache(c *gin.Context) {
	type messagePayload struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	var keyValue messagePayload

	if err := c.ShouldBindJSON(&keyValue); err != nil {
		utils.HandleError(c, utils.ErrBadRequest)
		return
	}

	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	botClaims, _ := service.ParseBotToken(token)

	err := h.BotService.AddKeyValuesToCache(botClaims.BotId, keyValue.Key, keyValue.Value)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleSuccess(c, keyValue)
}

// GetBotKeyFromCache godoc
// @Summary		Get bot key from cache
// @Description	Get bot key from cache
// @Authentication	Bearer BotToken
// @Tags			Bot
// @Accept		json
// @Produce		json
// @Param		key	path	string	true	"Key"
// @Success		200			{object}	utils.responseMessage	"Key Value"
// @Failure		400			{object}	utils.errorResponseMessage	"Bad request"
// @Router		/bot/token/key/{key} [get]
func (h *BotHandler) GetBotKeyFromCache(c *gin.Context) {
	key := c.Param("key")

	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	botClaims, _ := service.ParseBotToken(token)

	value, err := h.BotService.GetKeyFromCache(botClaims.BotId, key)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleSuccess(c, gin.H{"key": key, "value": value})
}

// RemoveBotKeyFromCache godoc
// @Summary		Remove bot key from cache
// @Description	Remove bot key from cache
// @Authentication	Bearer BotToken
// @Tags			Bot
// @Accept		json
// @Produce		json
// @Param		key	path	string	true	"Key"
// @Success		200			{object}	utils.responseMessage	"Key"
// @Failure		400			{object}	utils.errorResponseMessage	"Bad request"
// @Router		/bot/token/key/{key} [delete]
func (h *BotHandler) RemoveBotKeyFromCache(c *gin.Context) {
	key := c.Param("key")

	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	botClaims, _ := service.ParseBotToken(token)

	err := h.BotService.RemoveKeyValueFromCache(botClaims.BotId, key)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleSuccess(c, gin.H{"key": key})
}
