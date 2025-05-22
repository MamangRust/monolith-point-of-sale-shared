package merchantdocument_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

var (
	ErrMerchantDocumentNotFoundRes        = response.NewErrorResponse("Merchant Document not found", http.StatusNotFound)
	ErrFailedFindAllMerchantDocuments     = response.NewErrorResponse("Failed to fetch Merchant Documents", http.StatusInternalServerError)
	ErrFailedFindActiveMerchantDocuments  = response.NewErrorResponse("Failed to fetch active Merchant Documents", http.StatusInternalServerError)
	ErrFailedFindTrashedMerchantDocuments = response.NewErrorResponse("Failed to fetch trashed Merchant Documents", http.StatusInternalServerError)
	ErrFailedFindMerchantDocumentById     = response.NewErrorResponse("Failed to find Merchant Document by ID", http.StatusInternalServerError)

	ErrFailedCreateMerchantDocument = response.NewErrorResponse("Failed to create Merchant Document", http.StatusInternalServerError)
	ErrFailedUpdateMerchantDocument = response.NewErrorResponse("Failed to update Merchant Document", http.StatusInternalServerError)

	ErrFailedTrashMerchantDocument   = response.NewErrorResponse("Failed to trash Merchant Document", http.StatusInternalServerError)
	ErrFailedRestoreMerchantDocument = response.NewErrorResponse("Failed to restore Merchant Document", http.StatusInternalServerError)
	ErrFailedDeleteMerchantDocument  = response.NewErrorResponse("Failed to delete Merchant Document permanently", http.StatusInternalServerError)

	ErrFailedRestoreAllMerchantDocuments = response.NewErrorResponse("Failed to restore all Merchant Documents", http.StatusInternalServerError)
	ErrFailedDeleteAllMerchantDocuments  = response.NewErrorResponse("Failed to delete all Merchant Documents permanently", http.StatusInternalServerError)
)
