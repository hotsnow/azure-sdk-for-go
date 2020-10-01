// // Copyright (c) Microsoft Corporation. All rights reserved.
// // Licensed under the MIT License.

package azidentity

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/mock"
)

func TestInteractiveBrowserCredential_GetTokenSuccessMock(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.AppendResponse(mock.WithBody([]byte(accessTokenRespSuccess)))
	cred, err := NewInteractiveBrowserCredential(tenantID, clientID, secret, &TokenCredentialOptions{AuthorityHost: srv.URL()})
	if err != nil {
		t.Fatalf("Unable to create credential. Received: %v", err)
	}
	authCodeReceiver = func(tenantID string, clientID string, scopes []string) (*interactiveConfig, error) {
		return &interactiveConfig{
			authCode:    "12345",
			redirectURI: srv.URL(),
		}, nil
	}
	tk, err := cred.GetToken(context.Background(), azcore.TokenRequestOptions{Scopes: []string{"https://storage.azure.com/.default"}})
	if err != nil {
		t.Fatalf("Expected an empty error but received: %v", err)
	}
	if tk.Token != "new_token" {
		t.Fatal("Received unexpected token")
	}
}

func TestInteractiveBrowserCredential_GetTokenInvalidCredentials(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse(mock.WithBody([]byte(accessTokenRespError)), mock.WithStatusCode(http.StatusUnauthorized))
	cred, err := NewInteractiveBrowserCredential(tenantID, clientID, wrongSecret, &TokenCredentialOptions{AuthorityHost: srv.URL()})
	if err != nil {
		t.Fatalf("Unable to create credential. Received: %v", err)
	}
	authCodeReceiver = func(tenantID string, clientID string, scopes []string) (*interactiveConfig, error) {
		return &interactiveConfig{
			authCode:    "12345",
			redirectURI: srv.URL(),
		}, nil
	}
	_, err = cred.GetToken(context.Background(), azcore.TokenRequestOptions{Scopes: []string{scope}})
	if err == nil {
		t.Fatalf("Expected an error but did not receive one.")
	}
	var authFailed *AuthenticationFailedError
	if !errors.As(err, &authFailed) {
		t.Fatalf("Expected: AuthenticationFailedError, Received: %T", err)
	} else {
		var respError *AADAuthenticationFailedError
		if !errors.As(authFailed.Unwrap(), &respError) {
			t.Fatalf("Expected: AADAuthenticationFailedError, Received: %T", err)
		} else {
			if len(respError.Message) == 0 {
				t.Fatalf("Did not receive an error message")
			}
			if len(respError.Description) == 0 {
				t.Fatalf("Did not receive an error description")
			}
			if len(respError.Timestamp) == 0 {
				t.Fatalf("Did not receive a timestamp")
			}
			if len(respError.TraceID) == 0 {
				t.Fatalf("Did not receive a TraceID")
			}
			if len(respError.CorrelationID) == 0 {
				t.Fatalf("Did not receive a CorrelationID")
			}
			if len(respError.URI) == 0 {
				t.Fatalf("Did not receive an error URI")
			}
			if respError.Response == nil {
				t.Fatalf("Did not receive an error response")
			}
		}
	}
}