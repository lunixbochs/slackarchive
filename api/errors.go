package api

import (
	"net/http"

	errors "github.com/lunixbochs/slackarchive/api/errors"
)

var (
	ErrUserExists                    error = errors.New("user_exists", "User already exists", 404)
	ErrUserNotFound                  error = errors.New("user_notfound", "User not found", 404)
	ErrPasswordInvalid               error = errors.New("password_invalid", "Invalid password", 403)
	ErrPaymentCancelled              error = errors.New("payment_cancelled", "Payment cancelled", 404)
	ErrLoginInvalid                  error = errors.New("login_invalid", "Invalid password", 404)
	ErrTokenExpired                  error = errors.New("token_expired", "Token expired", 404)
	ErrTokenNotFound                 error = errors.New("token_notfound", "Token not found", 404)
	ErrApplicationNotFound           error = errors.New("application_not_found", "Application not found", 404)
	ErrPaymentChecksumFailed         error = errors.New("payment_checksumfailed", "Payment checksum failed", 404)
	ErrNotAuthorized                 error = errors.New("authentication_failed", "Authentication failed", http.StatusUnauthorized)
	ErrNotFound                            = errors.New("not-found", "Not authorized", 404)
	ErrValidationFailed                    = errors.New("validation-failed", "Validation errors", 417)
	ErrTimeout                             = errors.New("Timeout", "timeout", 500)
	ErrUnknownMethod                       = errors.New("Method not supported", "method-not-supported", 500)
	ErrDomainAlreadyExists                 = errors.New("domain-already-exists", "Domain already exists", 409)
	ErrApplicationAlreadyExists            = errors.New("application-already-exists", "Application already exists", 409)
	ErrDomainNotFound                      = errors.New("domain-not-found", "Domain not found", 404)
	ErrUserEmailAlreadyVerified            = errors.New("email-already-verified", "Email has been verified already", 417)
	ErrTeamNotFound                        = errors.New("team-not-found", "Team not found", 404)
	ErrTeamNotAnOwner                      = errors.New("team-not-an-owner", "Team not an owner", 404)
	ErrMemberNotFound                      = errors.New("member-not-found", "Member not found", 404)
	ErrTeamMemberCannotRemoveSelf          = errors.New("member-cannot-remove-self", "Member cannot remove self", 404)
	ErrCertificateNotFound                 = errors.New("certificate-not-found", "Certificate not found", 417)
	ErrCertificateCouldNotDecode           = errors.New("certificate-could-not-decode", "Could not decode certificate", 417)
	ErrCertificateCouldNotParse            = errors.New("certificate-could-not-parse", "Could not parse certificate", 417)
	ErrDatabaseAlreadyExists               = errors.New("already-exists", "Already exists", 409)
	ErrDatabaseOther                       = errors.New("other", "Other", 500)
	ErrCertificateVerificationFailed       = errors.New("certificate-verification-failed", "Certificate verification failed", 417)
)
