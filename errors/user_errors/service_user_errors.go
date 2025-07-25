package user_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-point-of-sale-shared/domain/response"
)

var (
	ErrUserNotFoundRes       = response.NewErrorResponse("User not found", http.StatusNotFound)
	ErrUserEmailAlready      = response.NewErrorResponse("User email already exists", http.StatusBadRequest)
	ErrUserPassword          = response.NewErrorResponse("Failed invalid password", http.StatusBadRequest)
	ErrFailedPasswordNoMatch = response.NewErrorResponse("Failed password not match", http.StatusBadRequest)
	ErrFailedFindAll         = response.NewErrorResponse("Failed to fetch users", http.StatusInternalServerError)
	ErrFailedFindActive      = response.NewErrorResponse("Failed to fetch active users", http.StatusInternalServerError)
	ErrFailedFindTrashed     = response.NewErrorResponse("Failed to fetch trashed users", http.StatusInternalServerError)
	ErrInternalServerError   = response.NewErrorResponse("Internal server error", http.StatusInternalServerError)

	ErrFailedSendEmail = response.NewErrorResponse("Failed to send email", http.StatusInternalServerError)

	ErrFailedCreateUser = response.NewErrorResponse("Failed to create user", http.StatusInternalServerError)
	ErrFailedUpdateUser = response.NewErrorResponse("Failed to update user", http.StatusInternalServerError)

	ErrFailedTrashedUser     = response.NewErrorResponse("Failed to move user to trash", http.StatusInternalServerError)
	ErrFailedRestoreUser     = response.NewErrorResponse("Failed to restore user", http.StatusInternalServerError)
	ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete user permanently", http.StatusInternalServerError)

	ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all users", http.StatusInternalServerError)
	ErrFailedDeleteAll  = response.NewErrorResponse("Failed to delete all users permanently", http.StatusInternalServerError)
)
