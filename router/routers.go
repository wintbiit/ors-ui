/*
 * RMServerAssist
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},

	{
		"GameBlueGet",
		http.MethodGet,
		"/game/blue",
		GameBlueGet,
	},

	{
		"GameProgressGet",
		http.MethodGet,
		"/game/progress",
		GameProgressGet,
	},

	{
		"GameRedGet",
		http.MethodGet,
		"/game/red",
		GameRedGet,
	},

	{
		"GameRoundCurrentGet",
		http.MethodGet,
		"/game/round/current",
		GameRoundCurrentGet,
	},

	{
		"GameRoundTotalGet",
		http.MethodGet,
		"/game/round/total",
		GameRoundTotalGet,
	},

	{
		"GameTimeGet",
		http.MethodGet,
		"/game/time",
		GameTimeGet,
	},

	{
		"GameTypeGet",
		http.MethodGet,
		"/game/type",
		GameTypeGet,
	},

	{
		"JudgeCoinBluePost",
		http.MethodPost,
		"/judge/coin/blue",
		JudgeCoinBluePost,
	},

	{
		"JudgeCoinRedPost",
		http.MethodPost,
		"/judge/coin/red",
		JudgeCoinRedPost,
	},

	{
		"JudgeCommandPost",
		http.MethodPost,
		"/judge/command",
		JudgeCommandPost,
	},

	{
		"JudgeCountdownPost",
		http.MethodPost,
		"/judge/countdown",
		JudgeCountdownPost,
	},

	{
		"JudgeKickallPost",
		http.MethodPost,
		"/judge/kickall",
		JudgeKickallPost,
	},

	{
		"JudgeKillallPost",
		http.MethodPost,
		"/judge/killall",
		JudgeKillallPost,
	},

	{
		"JudgePauseCampLengthPost",
		http.MethodPost,
		"/judge/pause/:camp/:length",
		JudgePauseCampLengthPost,
	},

	{
		"JudgePreparePost",
		http.MethodPost,
		"/judge/prepare",
		JudgePreparePost,
	},

	{
		"JudgeResetPost",
		http.MethodPost,
		"/judge/reset",
		JudgeResetPost,
	},

	{
		"JudgeResetallPost",
		http.MethodPost,
		"/judge/resetall",
		JudgeResetallPost,
	},

	{
		"JudgeSelfcheckPost",
		http.MethodPost,
		"/judge/selfcheck",
		JudgeSelfcheckPost,
	},

	{
		"JudgeSettleBluePost",
		http.MethodPost,
		"/judge/settle/blue",
		JudgeSettleBluePost,
	},

	{
		"JudgeSettleErrorPost",
		http.MethodPost,
		"/judge/settle/error",
		JudgeSettleErrorPost,
	},

	{
		"JudgeSettleRedPost",
		http.MethodPost,
		"/judge/settle/red",
		JudgeSettleRedPost,
	},

	{
		"RobotIdBalancePost",
		http.MethodPost,
		"/robot/:id/balance",
		RobotIdBalancePost,
	},

	{
		"RobotIdExpPost",
		http.MethodPost,
		"/robot/:id/exp",
		RobotIdExpPost,
	},

	{
		"RobotIdFindPost",
		http.MethodPost,
		"/robot/:id/find",
		RobotIdFindPost,
	},

	{
		"RobotIdGet",
		http.MethodGet,
		"/robot/:id",
		RobotIdGet,
	},

	{
		"RobotIdKickPost",
		http.MethodPost,
		"/robot/:id/kick",
		RobotIdKickPost,
	},

	{
		"RobotIdKillPost",
		http.MethodPost,
		"/robot/:id/kill",
		RobotIdKillPost,
	},

	{
		"RobotIdLevelupPost",
		http.MethodPost,
		"/robot/:id/levelup",
		RobotIdLevelupPost,
	},

	{
		"RobotIdRedPost",
		http.MethodPost,
		"/robot/:id/red",
		RobotIdRedPost,
	},

	{
		"RobotIdResetPost",
		http.MethodPost,
		"/robot/:id/reset",
		RobotIdResetPost,
	},

	{
		"RobotIdYellowPost",
		http.MethodPost,
		"/robot/:id/yellow",
		RobotIdYellowPost,
	},

	{
		"RobotListGet",
		http.MethodGet,
		"/robot/list",
		RobotListGet,
	},

	{
		"PingGet",
		http.MethodGet,
		"/ping",
		PingGet,
	},

	{
		"VersionGet",
		http.MethodGet,
		"/version",
		VersionGet,
	},
}

func respond(c *gin.Context, status int, data interface{}, message string) {
	timeStamp := time.Now().Unix()

	c.JSON(status, Response{
		Message:   message,
		Timestamp: timeStamp,
		Data:      data,
	})
}

func respondOk(c *gin.Context, data interface{}) {
	respond(c, http.StatusOK, data, "ok")
}

func respondInternalServerError(c *gin.Context) {
	respond(c, http.StatusInternalServerError, nil, "internal server error")
}

func respondBadRequest(c *gin.Context) {
	respond(c, http.StatusBadRequest, nil, "bad request")
}
