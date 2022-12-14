package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConnectorProfilePropertiesProperty := &customConnectorProfilePropertiesProperty{
//   	oAuth2Properties: &oAuth2PropertiesProperty{
//   		oAuth2GrantType: jsii.String("oAuth2GrantType"),
//   		tokenUrl: jsii.String("tokenUrl"),
//   		tokenUrlCustomProperties: map[string]*string{
//   			"tokenUrlCustomPropertiesKey": jsii.String("tokenUrlCustomProperties"),
//   		},
//   	},
//   	profileProperties: map[string]*string{
//   		"profilePropertiesKey": jsii.String("profileProperties"),
//   	},
//   }
//
type CfnConnectorProfile_CustomConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.CustomConnectorProfilePropertiesProperty.OAuth2Properties`.
	OAuth2Properties interface{} `field:"optional" json:"oAuth2Properties" yaml:"oAuth2Properties"`
	// `CfnConnectorProfile.CustomConnectorProfilePropertiesProperty.ProfileProperties`.
	ProfileProperties interface{} `field:"optional" json:"profileProperties" yaml:"profileProperties"`
}

