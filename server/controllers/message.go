package controllers

import (
	"context"
	"github.com/mylxsw/aidea-server/pkg/repo"
	"github.com/mylxsw/aidea-server/server/auth"
	"github.com/mylxsw/aidea-server/server/controllers/common"
	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/glacier/web"
	"net/http"
	"strconv"
)

type MessageController struct {
	repo *repo.Repository `autowire:"@"`
	// conf *config.Config   `autowire:"@"`
}

func NewMessageController(resolver infra.Resolver) web.Controller {
	ctl := MessageController{}
	resolver.MustAutoWire(&ctl)
	return &ctl
}

func (ctl *MessageController) Register(router web.Router) {
	router.Group("/messages", func(router web.Router) {
		router.Post("/{message_id}/rating", ctl.UpdateMessageRating)
	})
}

// UpdateMessageRating 更新消息评分
func (ctl *MessageController) UpdateMessageRating(ctx context.Context, webCtx web.Context, user *auth.User) web.Response {
	messageID, err := strconv.Atoi(webCtx.PathVar("message_id"))
	if err != nil {
		return webCtx.JSONError(common.ErrInvalidRequest, http.StatusBadRequest)
	}

	rating := webCtx.Int64Input("rating", 5)

	if err := ctl.repo.Message.UpdateRating(ctx, user.ID, int64(messageID), rating); err != nil {
		return webCtx.JSONError(common.ErrInternalError, http.StatusInternalServerError)
	}

	return webCtx.JSON(web.M{})
}
