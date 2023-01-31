package help

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/prokhorind/nextcloud/function/locales"
	"net/http"
	"strings"
	"unicode"
)

func HandleHelpCommand(c *gin.Context) {
	creq := apps.CallRequest{}
	json.NewDecoder(c.Request.Body).Decode(&creq)
	builder := strings.Builder{}
	builder.WriteString(getSingleHelpMessage(c, creq, "title"))
	builder.WriteString("\n")
	builder.WriteString(createHelpForSingleCommand(c, creq, "configure"))
	builder.WriteString("\n")
	builder.WriteString(createHelpForSingleCommand(c, creq, "connect"))
	builder.WriteString("\n")
	builder.WriteString(createHelpForSingleCommand(c, creq, "share"))
	builder.WriteString("\n")
	builder.WriteString(createHelpForSingleCommand(c, creq, "calendars"))
	builder.WriteString("\n")
	builder.WriteString("\n")
	builder.WriteString(getSingleHelpMessage(c, creq, "tips"))
	c.JSON(http.StatusOK, apps.NewTextResponse(builder.String()))
}

func getSingleHelpMessage(c *gin.Context, request apps.CallRequest, message string) string {
	locale := request.Context.ActingUser.Locale
	messageSource := locales.MessageSource{c, locale}
	return messageSource.GetMessage("help." + message)
}

func createHelpForSingleCommand(c *gin.Context, request apps.CallRequest, command string) string {
	locale := request.Context.ActingUser.Locale
	messageSource := locales.MessageSource{c, locale}
	description := messageSource.GetMessage(fmt.Sprintf("help.%s", command))
	return fmt.Sprintf("%s - %s", changeFirstLetterToUpperCase(command), description)
}

func changeFirstLetterToUpperCase(str string) string {
	a := []rune(str)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
