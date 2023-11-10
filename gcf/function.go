package gcf

import (
	"github.com/GoogleCloudPlatform/function-framework-go/functions"
	"github.com/whatsauth/webhook"
)

func init() {
	functions.HTTP("Webhook", webhook.Post)
}
