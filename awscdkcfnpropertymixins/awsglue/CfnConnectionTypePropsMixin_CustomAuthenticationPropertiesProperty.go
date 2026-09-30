package awsglue


// Custom authentication configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   customAuthenticationPropertiesProperty := &CustomAuthenticationPropertiesProperty{
//   	AuthenticationParameters: []interface{}{
//   		&SecretConnectorPropertyProperty{
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-customauthenticationproperties.html
//
type CfnConnectionTypePropsMixin_CustomAuthenticationPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-customauthenticationproperties.html#cfn-glue-connectiontype-customauthenticationproperties-authenticationparameters
	//
	AuthenticationParameters interface{} `field:"optional" json:"authenticationParameters" yaml:"authenticationParameters"`
}

