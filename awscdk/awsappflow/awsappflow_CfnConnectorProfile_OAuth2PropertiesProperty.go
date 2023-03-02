package awsappflow


// The OAuth 2.0 properties required for OAuth 2.0 authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuth2PropertiesProperty := &oAuth2PropertiesProperty{
//   	oAuth2GrantType: jsii.String("oAuth2GrantType"),
//   	tokenUrl: jsii.String("tokenUrl"),
//   	tokenUrlCustomProperties: map[string]*string{
//   		"tokenUrlCustomPropertiesKey": jsii.String("tokenUrlCustomProperties"),
//   	},
//   }
//
type CfnConnectorProfile_OAuth2PropertiesProperty struct {
	// The OAuth 2.0 grant type used by connector for OAuth 2.0 authentication.
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
	// The token URL required for OAuth 2.0 authentication.
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// Associates your token URL with a map of properties that you define.
	//
	// Use this parameter to provide any additional details that the connector requires to authenticate your request.
	TokenUrlCustomProperties interface{} `field:"optional" json:"tokenUrlCustomProperties" yaml:"tokenUrlCustomProperties"`
}

