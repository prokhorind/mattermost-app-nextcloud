package function

import (
	"github.com/gin-gonic/gin"
	"github.com/prokhorind/nextcloud/function/calendar"
	"github.com/prokhorind/nextcloud/function/file"
	"github.com/prokhorind/nextcloud/function/help"
	"github.com/prokhorind/nextcloud/function/install"
	"github.com/prokhorind/nextcloud/function/oauth"
)

func InitHandlers(r *gin.Engine) {
	r.POST("/bindings", install.Bindings)
	r.POST("/configure", oauth.Configure)
	r.POST("/connect", oauth.HandleConnect)
	r.POST("/disconnect", oauth.Disconnect)
	r.POST("/oauth2/complete", oauth.Oauth2Complete)
	r.POST("/oauth2/connect", oauth.Oauth2Connect)
	r.POST("/file/search/form", file.FileShareForm)
	r.POST("/file-share", file.FileShare)
	r.POST("/create-calendar-event", calendar.HandleCreateEvent)
	r.POST("/create-calendar-event-form", calendar.HandleCreateEventForm)
	r.POST("/get-calendar-events-today", calendar.HandleGetEventsToday)
	r.POST("/get-calendar-events-tomorrow", calendar.HandleGetEventsTomorrow)
	r.POST("/get-calendar-events-select-date-form", calendar.GetUserSelectedEventsDate)
	r.POST("/get-calendar-events-select-date/:calendar", calendar.HandleGetEventsAtSelectedDay)
	r.POST("/delete-event/:calendarId/events/:eventId", calendar.HandleDeleteCalendarEvent)
	r.POST("/do-nothing", calendar.DoNothing)
	r.POST("/redirect/meeting", calendar.RedirectToAMeeting)
	r.POST("/help", help.HandleHelpCommand)

	r.POST("/get-parsed-date", calendar.HandleGetParsedCalendarDate)
	r.POST("/file-upload-form", file.FileUploadForm)
	r.POST("/file-upload", file.FileUpload)

	r.POST("/ping", install.Ping)
	r.POST("/calendars", calendar.HandleGetUserCalendars)
	r.POST("/users/:userId/calendars/:calendarId/events/:eventId/status/:status", calendar.HandleChangeEventStatus)
}
