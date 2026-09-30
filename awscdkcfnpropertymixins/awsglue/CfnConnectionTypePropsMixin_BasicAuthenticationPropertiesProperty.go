package awsglue


// Basic authentication configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   basicAuthenticationPropertiesProperty := &BasicAuthenticationPropertiesProperty{
//   	Password: &SecretConnectorPropertyProperty{
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   	Username: &SecretConnectorPropertyProperty{
//   		KeyOverride: jsii.String("keyOverride"),
//   		Name: jsii.String("name"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-basicauthenticationproperties.html
//
type CfnConnectionTypePropsMixin_BasicAuthenticationPropertiesProperty struct {
	// Defines a secret property configuration.
	//
	// SECRET-type properties cannot have DefaultValue or AllowedValues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-basicauthenticationproperties.html#cfn-glue-connectiontype-basicauthenticationproperties-password
	//
	Password interface{} `field:"optional" json:"password" yaml:"password"`
	// Defines a secret property configuration.
	//
	// SECRET-type properties cannot have DefaultValue or AllowedValues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-basicauthenticationproperties.html#cfn-glue-connectiontype-basicauthenticationproperties-username
	//
	Username interface{} `field:"optional" json:"username" yaml:"username"`
}

