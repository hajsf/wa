package global

import (
	"context"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var Ctx context.Context
var Cancel context.CancelFunc
var Bundle *i18n.Bundle

var Log waLog.Logger
