package awsappflow


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
	// `CfnConnectorProfile.OAuth2PropertiesProperty.OAuth2GrantType`.
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
	// `CfnConnectorProfile.OAuth2PropertiesProperty.TokenUrl`.
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// `CfnConnectorProfile.OAuth2PropertiesProperty.TokenUrlCustomProperties`.
	TokenUrlCustomProperties interface{} `field:"optional" json:"tokenUrlCustomProperties" yaml:"tokenUrlCustomProperties"`
}

