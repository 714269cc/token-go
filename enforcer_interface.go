package token_go

import (
	"github.com/weloe/token-go/config"
	"github.com/weloe/token-go/ctx"
	"github.com/weloe/token-go/log"
	"github.com/weloe/token-go/model"
	"github.com/weloe/token-go/persist"
)

var _ IEnforcer = &Enforcer{}

type IEnforcer interface {
	// Login login api
	Login(id string, ctx ctx.Context) (string, error)
	LoginById(id string) (string, error)
	LoginByModel(id string, loginModel *model.Login, ctx ctx.Context) (string, error)
	Logout(ctx ctx.Context) error
	LogoutById(id string) error
	LogoutByToken(token string) error
	IsLogin(ctx ctx.Context) (bool, error)
	IsLoginByToken(token string) (bool, error)
	IsLoginById(id string) (bool, error)
	GetLoginId(ctx ctx.Context) (string, error)
	GetIdByToken(token string) string
	GetLoginCount(id string) int

	Replaced(id string, device string) error
	// Banned banned api
	Banned(id string, service string, level int, time int64) error
	UnBanned(id string, services ...string) error
	IsBanned(id string, service string) bool
	GetBannedLevel(id string, service string) (int64, error)
	GetBannedTime(id string, service string) int64

	Kickout(id string, device string) error

	GetRequestToken(ctx ctx.Context) string
	AddTokenGenerateFun(tokenStyle string, f model.GenerateFunc) error

	CheckLogin(ctx ctx.Context) error

	SetAuth(manager interface{})
	CheckRole(ctx ctx.Context, role string) error
	CheckPermission(ctx ctx.Context, permission string) error

	SetType(t string)
	GetType() string
	GetAdapter() persist.Adapter
	SetAdapter(adapter persist.Adapter)
	SetWatcher(watcher persist.Watcher)
	GetWatcher() persist.Watcher
	SetLogger(logger log.Logger)
	GetLogger() log.Logger
	EnableLog()
	IsLogEnable() bool
	GetSession(id string) *model.Session
	UpdateSession(id string, session *model.Session) error
	SetSession(id string, session *model.Session, timeout int64) error
	GetTokenConfig() config.TokenConfig
}
