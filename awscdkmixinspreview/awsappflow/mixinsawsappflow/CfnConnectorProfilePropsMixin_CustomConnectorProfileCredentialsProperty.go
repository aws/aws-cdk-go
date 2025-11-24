package mixinsawsappflow


// The connector-specific profile credentials that are required when using the custom connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customConnectorProfileCredentialsProperty := &CustomConnectorProfileCredentialsProperty{
//   	ApiKey: &ApiKeyCredentialsProperty{
//   		ApiKey: jsii.String("apiKey"),
//   		ApiSecretKey: jsii.String("apiSecretKey"),
//   	},
//   	AuthenticationType: jsii.String("authenticationType"),
//   	Basic: &BasicAuthCredentialsProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	Custom: &CustomAuthCredentialsProperty{
//   		CredentialsMap: map[string]*string{
//   			"credentialsMapKey": jsii.String("credentialsMap"),
//   		},
//   		CustomAuthenticationType: jsii.String("customAuthenticationType"),
//   	},
//   	Oauth2: &OAuth2CredentialsProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		OAuthRequest: &ConnectorOAuthRequestProperty{
//   			AuthCode: jsii.String("authCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofilecredentials.html
//
type CfnConnectorProfilePropsMixin_CustomConnectorProfileCredentialsProperty struct {
	// The API keys required for the authentication of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofilecredentials.html#cfn-appflow-connectorprofile-customconnectorprofilecredentials-apikey
	//
	ApiKey interface{} `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The authentication type that the custom connector uses for authenticating while creating a connector profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofilecredentials.html#cfn-appflow-connectorprofile-customconnectorprofilecredentials-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The basic credentials that are required for the authentication of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofilecredentials.html#cfn-appflow-connectorprofile-customconnectorprofilecredentials-basic
	//
	Basic interface{} `field:"optional" json:"basic" yaml:"basic"`
	// If the connector uses the custom authentication mechanism, this holds the required credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofilecredentials.html#cfn-appflow-connectorprofile-customconnectorprofilecredentials-custom
	//
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// The OAuth 2.0 credentials required for the authentication of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofilecredentials.html#cfn-appflow-connectorprofile-customconnectorprofilecredentials-oauth2
	//
	Oauth2 interface{} `field:"optional" json:"oauth2" yaml:"oauth2"`
}

