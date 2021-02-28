package msgraph

import (
	"context"
	"errors"
)

//GetMe returns email address for the authenticated user
func (a *API) GetMe(ctx context.Context) (string, error) {
	req := a.Client.Me().Request()
	user, err := req.Get(ctx)
	if err != nil {
		return "", err
	}
	if user.Mail == nil {
		return "", errors.New("could not retrieve your email address")
	}
	return *user.Mail, nil
}
