package awsappflow


// The profile properties required by the custom connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConnectorProfilePropertiesProperty := &CustomConnectorProfilePropertiesProperty{
//   	OAuth2Properties: &OAuth2PropertiesProperty{
//   		OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   		TokenUrl: jsii.String("tokenUrl"),
//   		TokenUrlCustomProperties: map[string]*string{
//   			"tokenUrlCustomPropertiesKey": jsii.String("tokenUrlCustomProperties"),
//   		},
//   	},
//   	ProfileProperties: map[string]*string{
//   		"profilePropertiesKey": jsii.String("profileProperties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofileproperties.html
//
type CfnConnectorProfile_CustomConnectorProfilePropertiesProperty struct {
	// The OAuth 2.0 properties required for OAuth 2.0 authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofileproperties.html#cfn-appflow-connectorprofile-customconnectorprofileproperties-oauth2properties
	//
	OAuth2Properties interface{} `field:"optional" json:"oAuth2Properties" yaml:"oAuth2Properties"`
	// A map of properties that are required to create a profile for the custom connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofileproperties.html#cfn-appflow-connectorprofile-customconnectorprofileproperties-profileproperties
	//
	ProfileProperties interface{} `field:"optional" json:"profileProperties" yaml:"profileProperties"`
}

