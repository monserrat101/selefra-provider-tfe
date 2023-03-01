package resources

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hashicorp/go-tfe"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
)

type Config struct {
	WorkspaceId  string `yaml:"workspace_id" json:"workspace_id" mapstructure:"workspace_id"`
	TeamId       string `yaml:"team_id" json:"team_id" mapstructure:"team_id"`
	Organization string `yaml:"organization" json:"organization" mapstructure:"organization"`
	AgentPoolId  string `yaml:"agent_pool_id" json:"agent_pool_id" mapstructure:"agent_pool_id"`
	PolicySetId  string `yaml:"policy_set_id" json:"policy_set_id" mapstructure:"policy_set_id"`
	Token        string `yaml:"token" json:"token" mapstructure:"token"`
	ProjectId    string `yaml:"project_id" json:"project_id" mapstructure:"project_id"`
}

type Client struct {
	TerraformBridge *bridge.TerraformBridge

	*Config

	tfeClient *tfe.Client
}

func newClient(conf *Config) (*Client, error) {

	token, err := getToken()
	if err != nil {
		return nil, err
	}

	if token == "" && conf.Token != "" {
		token = conf.Token
	}

	config := &tfe.Config{
		Token: token,
	}

	tfeClient, err := tfe.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create tfe client: %v", err)
	}

	return &Client{
		tfeClient: tfeClient,
		Config:    conf,
	}, nil
}

func getToken() (string, error) {
	var token string
	token = os.Getenv("TFE_TOKEN")
	if token == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("get token failed: %v", err)
		}

		rcfContent, err := os.ReadFile(filepath.Join(homedir, ".terraformrc"))
		if err != nil {
			return "", fmt.Errorf("get token failed: %v", err)
		}
		exp, err := regexp.Compile(`token\s?=\s?"?\w+\.\w+.\w+"?`)
		if err != nil {
			return "", fmt.Errorf("get token failed: %v", err)
		}

		tokenExp := exp.Find(rcfContent)
		rawToken := strings.Split(string(tokenExp), "=")
		if len(rawToken) < 1 {
			return "", fmt.Errorf("failed to get tfe token, please set your tfe token")
		}
		token = strings.TrimSpace(strings.Replace(rawToken[1], "\"", "", -1))
	}

	if token == "" {
		return "", fmt.Errorf("failed to get tfe token, please set your tfe token")
	}

	return token, nil
}
