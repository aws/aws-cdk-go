package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pardotConnectorProfileCredentialsProperty := &PardotConnectorProfileCredentialsProperty{
//   	AccessToken: jsii.String("accessToken"),
//   	ClientCredentialsArn: jsii.String("clientCredentialsArn"),
//   	ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   		AuthCode: jsii.String("authCode"),
//   		RedirectUri: jsii.String("redirectUri"),
//   	},
//   	RefreshToken: jsii.String("refreshToken"),
//   }
//
type CfnConnectorProfile_PardotConnectorProfileCredentialsProperty struct {
	// `CfnConnectorProfile.PardotConnectorProfileCredentialsProperty.AccessToken`.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// `CfnConnectorProfile.PardotConnectorProfileCredentialsProperty.ClientCredentialsArn`.
	ClientCredentialsArn *string `field:"optional" json:"clientCredentialsArn" yaml:"clientCredentialsArn"`
	// `CfnConnectorProfile.PardotConnectorProfileCredentialsProperty.ConnectorOAuthRequest`.
	ConnectorOAuthRequest interface{} `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
	// `CfnConnectorProfile.PardotConnectorProfileCredentialsProperty.RefreshToken`.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

