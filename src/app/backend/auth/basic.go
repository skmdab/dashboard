// Copyright 2017 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// basic.go

package auth

import (
	"errors"

	authApi "github.com/kubernetes/dashboard/src/app/backend/auth/api"
	"k8s.io/client-go/tools/clientcmd/api"
)

// Implements Authenticator interface
type basicAuthenticator struct {
	// You would typically store user credentials securely, such as in a database
	// For simplicity, this example hardcodes a single user
	validUsername string
	validPassword string
}

// GetAuthInfo implements Authenticator interface. See Authenticator for more information.
func (self *basicAuthenticator) GetAuthInfo() (api.AuthInfo, error) {
	// Hardcoded user credentials for demonstration purposes
	if self.validUsername == "root" && self.validPassword == "rootpassword" {
		return api.AuthInfo{
			Username: self.validUsername,
			Password: self.validPassword,
		}, nil
	}

	return api.AuthInfo{}, errors.New("invalid credentials")
}

// NewBasicAuthenticator returns Authenticator based on LoginSpec.
func NewBasicAuthenticator(spec *authApi.LoginSpec) authApi.Authenticator {
	// Set valid user credentials based on the provided spec
	return &basicAuthenticator{
		validUsername: "root",
		validPassword: "rootpassword",
	}
}

