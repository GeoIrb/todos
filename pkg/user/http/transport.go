package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/geoirb/todos/pkg/user"
	"github.com/valyala/fasthttp"
)

// LoginTransport ...
type LoginTransport interface {
	DecodeRequest(req *fasthttp.Request) (login user.Login, err error)
	EncodeResponse(res *fasthttp.Response, auth user.Auth, err error)
}

type loginTransport struct{}

func newLoginTransport() LoginTransport {
	return &loginTransport{}
}

func (t loginTransport) DecodeRequest(req *fasthttp.Request) (login user.Login, err error) {
	var request loginRequest
	err = json.Unmarshal(req.Body(), &request)
	login = user.Login(request)
	return
}

func (t loginTransport) EncodeResponse(res *fasthttp.Response, auth user.Auth, err error) {
	response := loginResponse(auth)
	responseBuilder(res, response, err)
}

type GetUserListTransport interface {
	DecodeRequest(req *fasthttp.Request) (token string, filter user.Filter, err error)
	EncodeResponse(res *fasthttp.Response, users []user.UserInfo, err error)
}

type getUserListTransport struct{}

func newGetUserTransport() GetUserListTransport {
	return &getUserListTransport{}
}

func (t *getUserListTransport) DecodeRequest(req *fasthttp.Request) (token string, filter user.Filter, err error) {
	token = string(req.Header.Peek("Authorization"))
	args := req.URI().QueryArgs()
	if arg := args.Peek("id"); arg != nil {
		id, _ := strconv.Atoi(string(arg))
		filter.ID = &id
	}
	if arg := args.Peek("email"); arg != nil {
		email := string(arg)
		filter.Email = &email
	}
	return
}

func (t *getUserListTransport) EncodeResponse(res *fasthttp.Response, users []user.UserInfo, err error) {
	var response getUserListResponse
	response.Users = make([]userInfo, 0, len(users))
	for _, user := range users {
		response.Users = append(response.Users, userInfo(user))
	}
	responseBuilder(res, response, err)
}

type NewUserTransport interface {
	DecodeRequest(req *fasthttp.Request) (info user.Registartion, err error)
	EncodeResponse(res *fasthttp.Response, err error)
}

type newUserTransport struct{}

func newNewUserTransport() NewUserTransport {
	return &newUserTransport{}
}

func (t *newUserTransport) DecodeRequest(req *fasthttp.Request) (info user.Registartion, err error) {
	var request registrationRequest
	err = json.Unmarshal(req.Body(), &request)
	info = user.Registartion(request)
	return
}

func (t *newUserTransport) EncodeResponse(res *fasthttp.Response, err error) {
	responseBuilder(res, nil, err)
}

type ActivateUserTransport interface {
	DecodeRequest(req *fasthttp.Request) (info user.Create, err error)
	EncodeResponse(res *fasthttp.Response, err error)
}

type activateUserTransport struct{}

func newActivateUserTransport() ActivateUserTransport {
	return &activateUserTransport{}
}

func (t *activateUserTransport) DecodeRequest(req *fasthttp.Request) (info user.Create, err error) {
	var request createRequest
	err = json.Unmarshal(req.Body(), &request)
	info = user.Create(request)
	return
}

func (t *activateUserTransport) EncodeResponse(res *fasthttp.Response, err error) {
	responseBuilder(res, nil, err)
}

var responseCode map[error]int = map[error]int{
	user.ErrUserIsExist:        http.StatusBadRequest,
	user.ErrUserNotFound:       http.StatusUnauthorized,
	user.ErrFailedAuthenticate: http.StatusUnauthorized,
	user.ErrTokenExpired:       http.StatusUnauthorized,
}

func responseBuilder(res *fasthttp.Response, response interface{}, err error) {
	code := http.StatusOK
	if err != nil {
		var isExist bool
		if code, isExist = responseCode[err]; !isExist {
			code = http.StatusInternalServerError
		}
		res.SetStatusCode(code)
		res.SetBody([]byte(err.Error()))
		return
	}

	body, errMarshal := json.Marshal(response)
	if errMarshal != nil {
		code = http.StatusInternalServerError
		body = []byte(errMarshal.Error())
	}
	res.SetStatusCode(code)
	res.SetBody(body)
}
