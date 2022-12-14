package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConnectorProfileCredentialsProperty := &customConnectorProfileCredentialsProperty{
//   	authenticationType: jsii.String("authenticationType"),
//
//   	// the properties below are optional
//   	apiKey: &apiKeyCredentialsProperty{
//   		apiKey: jsii.String("apiKey"),
//
//   		// the properties below are optional
//   		apiSecretKey: jsii.String("apiSecretKey"),
//   	},
//   	basic: &basicAuthCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	custom: &customAuthCredentialsProperty{
//   		customAuthenticationType: jsii.String("customAuthenticationType"),
//
//   		// the properties below are optional
//   		credentialsMap: map[string]*string{
//   			"credentialsMapKey": jsii.String("credentialsMap"),
//   		},
//   	},
//   	oauth2: &oAuth2CredentialsProperty{
//   		accessToken: jsii.String("accessToken"),
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//   		oAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   }
//
type CfnConnectorProfile_CustomConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.CustomConnectorProfileCredentialsProperty.AuthenticationType`.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// `CfnConnectorProfile.CustomConnectorProfileCredentialsProperty.ApiKey`.
	ApiKey interface{} `field:"optional" json:"apiKey" yaml:"apiKey"`
	// `CfnConnectorProfile.CustomConnectorProfileCredentialsProperty.Basic`.
	Basic interface{} `field:"optional" json:"basic" yaml:"basic"`
	// `CfnConnectorProfile.CustomConnectorProfileCredentialsProperty.Custom`.
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// `CfnConnectorProfile.CustomConnectorProfileCredentialsProperty.Oauth2`.
	Oauth2 interface{} `field:"optional" json:"oauth2" yaml:"oauth2"`
}

