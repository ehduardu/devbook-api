package config

import (
	"api/src/utils"
)

type firebaseCredentials struct {
	Type                    string `json:"type,omitempty"`
	ProjectId               string `json:"project_id,omitempty"`
	PrivateKeyId            string `json:"private_key_id,omitempty"`
	PrivateKey              string `json:"private_key,omitempty"`
	ClientEmail             string `json:"client_email,omitempty"`
	ClientId                string `json:"client_id,omitempty"`
	AuthURI                 string `json:"auth_uri,omitempty"`
	TokenURI                string `json:"token_uri,omitempty"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url,omitempty"`
	ClientX509CertURL       string `json:"client_x509_cert_url,omitempty"`
	DatabaseURL             string `json:"database_url,omitempty"`
}

type internalSecret struct {
	SecretKey string `json:"secret_key,omitempty"`
}

var (
	INTERNAL_SECRET      internalSecret
	FIREBASE_CREDENTIALS firebaseCredentials
)

func LoadVariables() {
	utils.ParseFile("./.credentials/internal-secret.json", &INTERNAL_SECRET)
	utils.ParseFile("./.credentials/firebase-adminsdk.json", &FIREBASE_CREDENTIALS)
}
