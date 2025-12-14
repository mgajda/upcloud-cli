package clierrors

import (
	"errors"

	"github.com/UpCloudLtd/upcloud-go-api/v8/upcloud"
)

var _ ClientError = InvalidCredentialsError{}

type InvalidCredentialsError struct{}

func (err InvalidCredentialsError) ErrorCode() int {
	return InvalidCredentials
}

func (err InvalidCredentialsError) Error() string {
	return "invalid credentials or insufficient permissions. For token operations, ensure you're using account owner credentials or a token with 'account' permission scope. See https://upcloudltd.github.io/upcloud-cli/#configure-credentials"
}

func CheckAuthenticationFailed(err error) bool {
	prob := &upcloud.Problem{}

	if errors.As(err, &prob) {
		errCode := prob.ErrorCode()
		// Check for various authentication error codes from the API
		// UNAUTHORIZED is used by the tokens API endpoint
		if errCode == upcloud.ErrCodeAuthenticationFailed ||
		   errCode == "INVALID_CREDENTIALS" ||
		   errCode == "UNAUTHORIZED" {
			return true
		}
	}

	return false
}
