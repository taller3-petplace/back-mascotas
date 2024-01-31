package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	secretEnvVar                 = "secret"
	cryptographicAlgorithmEnvVar = "algorithm"
)

// AppContextCreator middleware use by each endpoint to create a context.AppContext
func AppContextCreator() gin.HandlerFunc {
	return func(c *gin.Context) {
		appRequestContext, err := NewAppContext(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(appRequestContext)
		c.Next()
	}
}

// AppContext context used by this app. It contains data, mainly from the user, that came in the request that can be use anywhere
type AppContext struct {
	TelegramRequest bool
	TelegramID      string
	UserID          string
	Email           string
}

// jwtData data that comes in the JWT
type jwtData struct {
	userID     string
	telegramID string
	email      string
}

type appContextKey struct{}

type appContextValue struct {
	Context AppContext
}

func NewAppContext(request *http.Request) (context.Context, error) {

	requestFromTelegram := request.Header.Get(Telegram) == "true"
	appContext := AppContext{
		TelegramRequest: requestFromTelegram,
	}

	if !requestFromTelegram {
		tokenString := request.Header.Get(JWT)
		tokenData, err := extractDataFromJWT(tokenString)
		if err != nil {
			return nil, fmt.Errorf("error extracting data from JWT: %v", err)
		}
		appContext.UserID = tokenData.userID
		appContext.Email = tokenData.email
		appContext.TelegramID = tokenData.telegramID
	}

	return context.WithValue(
		request.Context(),
		appContextKey{},
		appContextValue{
			appContext,
		},
	), nil
}

// GetAppContext from the given context extracts the AppContext that should have been added by the middleware
func GetAppContext(ctx context.Context) (AppContext, error) {
	if ctx == nil {
		return AppContext{}, errNilContext
	}

	contextValue := ctx.Value(appContextKey{})
	if contextValue == nil {
		return AppContext{}, errMissingAppContext
	}

	appContext := contextValue.(appContextValue)
	return appContext.Context, nil
}
