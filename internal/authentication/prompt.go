package authentication

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/dpastoor/wbi/internal/config"
)

func PromptAuthentication() string {
	name := ""
	prompt := &survey.Select{
		Message: "Choose an authentication method:",
		Options: []string{"SAML", "OIDC", "Active Directory/LDAP", "PAM", "Other"},
	}
	err := survey.AskOne(prompt, &name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func ConvertAuthType(authChoice string) config.AuthType {
	switch authChoice {
	case "SAML":
		return config.SAML
	case "OIDC":
		return config.OIDC
	case "Active Directory/LDAP":
		return config.LDAP
	case "PAM":
		return config.PAM
	case "Other":
		return config.Other
	}
	return config.Other
}

func HandleAuthChoice(WBConfig *config.WBConfig, targetOS string) {
	switch WBConfig.AuthType {
	case config.SAML:
		HandleSAMLConfig(&WBConfig.SAMLConfig)
		fmt.Println("Setting up SAML based authentication is a 2 step process. The configurations just entered will be setup on Workbench to complete step 1. \n\nTo complete step 2, you must configure your identify provider with Workbench following steps outlined here: https://docs.posit.co/ide/server-pro/authenticating_users/saml_sso.html#step-2.-configure-your-identity-provider-with-workbench")
	case config.OIDC:
		fmt.Println("Setting up OpenID Connect based authentication is a 2 step process. First configure your OpenID provider with the steps outlined here to complete step 1: https://docs.posit.co/ide/server-pro/authenticating_users/openid_connect_authentication.html#configuring-your-openid-provider \n\n As you register Workbench in the IdP, save the client-id and client-secret. Follow the next step of prompts to configure Workbench to complete step 2.")
		HandleOIDCConfig(&WBConfig.OIDCConfig)
	case config.LDAP:
		switch targetOS {
		case "ubuntu22", "ubuntu20", "ubuntu18":
			fmt.Println("Posit Workbench connects to LDAP via PAM. Please follow this article for more details on configuration: \nhttps://support.posit.co/hc/en-us/articles/360024137174-Integrating-Ubuntu-with-Active-Directory-for-RStudio-Workbench-RStudio-Server-Pro")
		case "rhel9", "rhel8", "rhel7":
			fmt.Println("Posit Workbench connects to LDAP via PAM. Please follow this article for more details on configuration: \nhttps://support.posit.co/hc/en-us/articles/360016587973-Integrating-RStudio-Workbench-RStudio-Server-Pro-with-Active-Directory-using-CentOS-RHEL")
		default:
			log.Fatal("Unsupported operating system")
		}
	case config.PAM:
		fmt.Println("PAM requires no additional configuration, however there are some username considerations and home directory provisioning steps that can be taken. To learn more please visit: https://docs.posit.co/ide/server-pro/authenticating_users/pam_authentication.html")
	case config.Other:
		fmt.Println("To learn about configuring your desired method of authentication please visit: https://docs.posit.co/ide/server-pro/authenticating_users/authenticating_users.html")
	}
}