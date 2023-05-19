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
	Login(id string, ctx ctx.Context) (string, error)
	LoginByModel(id string, loginModel *model.Login, ctx ctx.Context) (string, error)
	Logout(ctx ctx.Context) error
	LogoutByToken(token string) error
	IsLogin(ctx ctx.Context) (bool, error)
	IsLoginById(id string) (bool, error)
	GetLoginId(ctx ctx.Context) (string, error)
	GetIdByToken(token string) string
	GetLoginCount(id string) int

	Replaced(id string, device string) error
	// Banned TODO
	Banned(id string, service string) error
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
	SetSession(id string, session *model.Session, timeout int64) error
	GetTokenConfig() config.TokenConfig
}
