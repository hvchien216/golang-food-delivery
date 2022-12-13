package skio

import (
	"context"
	"errors"
	"fmt"
	"food_delivery/component/appctx"
	"food_delivery/component/tokenprovider/jwt"
	"food_delivery/modules/user/userstorage"
	"food_delivery/modules/user/usertransport/skuser"
	socketio "github.com/googollee/go-socket.io"
)

func Setup(appCtx appctx.AppContext, engine *rtEngine) {
	server := engine.server

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected: ", s.ID(), "Ip: ", s.RemoteAddr(), s.ID())
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
			s.Close()
			return
		}
		user.Mask(false)

		appSck := NewAppSocket(s, user)
		engine.saveAppSocket(user.Id, appSck)

		s.Emit("authenticated", user)

		server.OnEvent("/", "UserUpdateLocation", skuser.OnUserUpdateLocation(appCtx, user))

	})
}
