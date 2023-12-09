package larki

import (
	"context"
	"io"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type Client struct {
	*lark.Client
	*Config
	*BotInfo
	EventDispatcher  *dispatcher.EventDispatcher
	MessageEvent     <-chan *MessageEvent
	BotAddedEvent    <-chan *BotAddedEvent
	ChatCreatedEvent <-chan *ChatCreatedEvent
	MessageClient
	ImageClient
}

type Config struct {
	AppID       string
	AppSecret   string
	VerifyToken string
	EncryptKey  string
}

type MessageClient interface {
	GetMessage(ctx context.Context, messageId string) (*larkim.Message, error)
	ReplyMessage(ctx context.Context, message, messageId, messageType string) error
	ReplyText(ctx context.Context, messageId, title string, text ...string) error
	ReplyImage(ctx context.Context, messageId, imageKey string) error
	ReplyCard(ctx context.Context, messageId, card string) error
	ReplyCardTemplate(ctx context.Context, messageId, templateId string, vars map[string]interface{}) error
	SendMessage(ctx context.Context, receiverIdType, message, receiveId, messageType string) (string, error)
	SendMessageToGroup(ctx context.Context, groupId, message, messageType string) (string, error)
	SendTextToGroup(ctx context.Context, groupId, title string, text ...string) (string, error)
	SendImageToGroup(ctx context.Context, groupId, imageKey string) (string, error)
	SendCardToGroup(ctx context.Context, groupId, card string) (string, error)
	SendCardTemplateToGroup(ctx context.Context, groupId, templateId string, vars map[string]interface{}) (string, error)
	SendMessageToUser(ctx context.Context, openId, message, messageType string) (string, error)
	SendTextToUser(ctx context.Context, openId, title string, text ...string) (string, error)
	SendImageToUser(ctx context.Context, openId, imageKey string) (string, error)
	SendCardToUser(ctx context.Context, openId, card string) (string, error)
	SendCardTemplateToUser(ctx context.Context, openId, templateId string, vars map[string]interface{}) (string, error)
}

type ImageClient interface {
	GetImage(ctx context.Context, messageId, imageKey string) (io.Reader, error)
	UploadImage(ctx context.Context, reader io.Reader) (string, error)
}

type DocumentClient interface{}

type BotInfo struct {
	ActivateStatus int    `json:"activate_status"`
	AppName        string `json:"app_name"`
	AvatarUrl      string `json:"avatar_url"`
	OpenID         string `json:"open_id"`
}

type MessageEvent struct {
	*larkim.P2MessageReceiveV1Data
}

type BotAddedEvent struct {
	*larkim.P2ChatMemberBotAddedV1Data
}

type ChatCreatedEvent struct {
	*larkim.P1P2PChatCreatedV1Data
}

type botInfoResp struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Bot  BotInfo `json:"bot"`
}

type textContent struct {
	Text string `json:"text"`
}

type imageContent struct {
	ImageKey string `json:"image_key"`
}

type templateCardContent struct {
	Type string                  `json:"type"`
	Data templateCardContentData `json:"data"`
}

type templateCardContentData struct {
	TemplateId        string                 `json:"template_id"`
	TemplateVariables map[string]interface{} `json:"template_variables"`
}
