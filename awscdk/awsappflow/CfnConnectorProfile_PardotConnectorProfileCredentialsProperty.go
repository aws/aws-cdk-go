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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofilecredentials.html
//
type CfnConnectorProfile_PardotConnectorProfileCredentialsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofilecredentials.html#cfn-appflow-connectorprofile-pardotconnectorprofilecredentials-accesstoken
	//
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofilecredentials.html#cfn-appflow-connectorprofile-pardotconnectorprofilecredentials-clientcredentialsarn
	//
	ClientCredentialsArn *string `field:"optional" json:"clientCredentialsArn" yaml:"clientCredentialsArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofilecredentials.html#cfn-appflow-connectorprofile-pardotconnectorprofilecredentials-connectoroauthrequest
	//
	ConnectorOAuthRequest interface{} `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofilecredentials.html#cfn-appflow-connectorprofile-pardotconnectorprofilecredentials-refreshtoken
	//
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

