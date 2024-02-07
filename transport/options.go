package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	helper "github.com/antoha2/task"
)

const authorizationHeader = "Authorization"

// func (webImpl *webImpl) UserIdentify(next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

func (webImpl *webImpl) UserIdentify(ctx context.Context, r *http.Request) context.Context {

	header := r.Header.Get(authorizationHeader)
	if header == "" {
		newErr := "аутентификация - пустой заголовок"
		log.Println(newErr)

		return nil
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErr := "аутентификация - неправильный заголовок"
		log.Println(newErr)

		return nil
	}

	userId, err := webImpl.authService.ParseToken(ctx, headerParts[1])
	if err != nil {
		newErr := fmt.Sprintf("(task) аутентификация - ошибка ParseToken()- %s", err)
		log.Println(newErr)

		return nil
	}
	if userId == 0 {
		newErr := "аутентификация - нет прав доступа"
		log.Println(newErr)

		return nil
	}

	userRoles, err := webImpl.authService.GetRoles(ctx, userId)
	if err != nil {
		newErr := "аутентификация - ошибка GetRoles()"
		log.Println(newErr)
		return nil
	}

	if len(userRoles) == 0 {
		newErr := " не назначена роль"
		log.Println(newErr)

		return nil
	}

	ctx = context.WithValue(ctx, helper.USER_ID, userId)
	ctx = context.WithValue(ctx, helper.USER_ROLE, userRoles)

	//next(w, r.WithContext(ctx))
	return ctx

}
