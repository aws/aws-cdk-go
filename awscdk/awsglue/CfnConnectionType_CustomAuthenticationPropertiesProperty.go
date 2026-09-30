package awsglue


// Custom authentication configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customAuthenticationPropertiesProperty := &CustomAuthenticationPropertiesProperty{
//   	AuthenticationParameters: []interface{}{
//   		&SecretConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-customauthenticationproperties.html
//
type CfnConnectionType_CustomAuthenticationPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-customauthenticationproperties.html#cfn-glue-connectiontype-customauthenticationproperties-authenticationparameters
	//
	AuthenticationParameters interface{} `field:"required" json:"authenticationParameters" yaml:"authenticationParameters"`
}

