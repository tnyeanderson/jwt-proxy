package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

const (
	EnvClientID  = "JWT_PROXY_CLIENT_ID"
	EnvListener  = "JWT_PROXY_LISTENER"
	EnvIssuerURI = "JWT_PROXY_ISSUER_URI"
	EnvKeySetURI = "JWT_PROXY_KEYSET_URI"
)

type Config struct {
	ClientID  string
	Listener  string
	IssuerURI string
	KeySetURI string
}

func getConfig() (*Config, error) {
	c := &Config{
		Listener: ":8989",
	}

	if v := os.Getenv(EnvListener); v != "" {
		c.Listener = v
	}

	if v := os.Getenv(EnvClientID); v != "" {
		c.ClientID = v
	} else {
		return nil, fmt.Errorf("missing required env var: %s", EnvClientID)
	}

	if v := os.Getenv(EnvIssuerURI); v != "" {
		c.IssuerURI = v
	} else {
		return nil, fmt.Errorf("missing required env var: %s", EnvIssuerURI)
	}

	if v := os.Getenv(EnvKeySetURI); v != "" {
		c.KeySetURI = v
	} else {
		return nil, fmt.Errorf("missing required env var: %s", EnvKeySetURI)
	}

	return c, nil
}

func main() {
	//slog.SetLogLoggerLevel(slog.LevelDebug)

	config, err := getConfig()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	verifier := oidc.NewVerifier(
		config.IssuerURI,
		oidc.NewRemoteKeySet(context.Background(), config.KeySetURI),
		&oidc.Config{ClientID: config.ClientID},
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rawToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if rawToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(nil)
			return
		}

		token, err := verifier.Verify(context.Background(), rawToken)
		if err != nil {
			slog.Debug(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(nil)
			return
		}

		slog.Debug("token validated", "token", token)

		w.Write(nil)
	})

	if err := http.ListenAndServe(":8989", nil); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
