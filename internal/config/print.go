package config

import "fmt"

// Prints the WBConfig configuration struct information to the console
func (WBConfig *WBConfig) ConfigStructToText() {
	WBConfig.PythonConfig.PythonStructToText()
	WBConfig.SSLConfig.SSLStructToText()
	WBConfig.AuthConfig.AuthStructToText()
	fmt.Println("\n=== Please restart Workbench after making these changes")
}

// Prints the PythonConfig configuration struct information to the console
func (PythonConfig *PythonConfig) PythonStructToText() {
	fmt.Println("\n=== Add to config file: /etc/rstudio/jupyter.conf:")
	fmt.Println("jupyter-exe=" + PythonConfig.JupyterPath)
}

// Prints the SSLConfig configuration struct information to the console
func (SSLConfig *SSLConfig) SSLStructToText() {
	fmt.Println("\n=== Add to config file: /etc/rstudio/rserver.conf:")
	fmt.Println("ssl-enabled=1")
	fmt.Println("ssl-certificate=" + SSLConfig.CertPath)
	fmt.Println("ssl-certificate-key=" + SSLConfig.KeyPath)
}

// Prints the AuthConfig configuration struct information to the console
func (AuthConfig *AuthConfig) AuthStructToText() {
	switch AuthConfig.AuthType {
	case SAML:
		AuthConfig.SAMLConfig.AuthSAMLStructToText()
	case OIDC:
		AuthConfig.OIDCConfig.AuthOIDCStructToText()
	default:

	}
}

// Prints the SAMLConfig configuration struct information to the console
func (SAMLConfig *SAMLConfig) AuthSAMLStructToText() {
	fmt.Println("\n=== Add to config file: /etc/rstudio/rserver.conf:")
	fmt.Println("auth-saml=1")
	fmt.Println("auth-saml-sp-attribute-username=" + SAMLConfig.AuthSamlSpAttributeUsername)
	fmt.Println("auth-saml-metadata-url=" + SAMLConfig.AuthSamlMetadataURL)
}

// Prints the OIDCConfig configuration struct information to the console
func (OIDCConfig *OIDCConfig) AuthOIDCStructToText() {
	fmt.Println("\n=== Add to config file: /etc/rstudio/rserver.conf:")
	fmt.Println("auth-openid=1")
	fmt.Println("auth-openid-issuer=" + OIDCConfig.AuthOpenIDIssuer)
	fmt.Println("auth-openid-username-claim=" + OIDCConfig.AuthOpenIDUsernameClaim)

	fmt.Println("\n=== Add to config file: /etc/rstudio/openid-client-secret:")
	fmt.Println("client-id=" + OIDCConfig.ClientID)
	fmt.Println("client-secret=" + OIDCConfig.ClientSecret)
}