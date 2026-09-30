package awsglue


// Defines a secret property configuration.
//
// SECRET-type properties cannot have DefaultValue or AllowedValues.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   secretConnectorPropertyProperty := &SecretConnectorPropertyProperty{
//   	KeyOverride: jsii.String("keyOverride"),
//   	Name: jsii.String("name"),
//   	PropertyLocation: jsii.String("propertyLocation"),
//   	PropertyType: jsii.String("propertyType"),
//   	Required: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html
//
type CfnConnectionTypePropsMixin_SecretConnectorPropertyProperty struct {
	// A key name to use when sending this property in API requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html#cfn-glue-connectiontype-secretconnectorproperty-keyoverride
	//
	KeyOverride *string `field:"optional" json:"keyOverride" yaml:"keyOverride"`
	// The name of the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html#cfn-glue-connectiontype-secretconnectorproperty-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies where this property should be included in REST requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html#cfn-glue-connectiontype-secretconnectorproperty-propertylocation
	//
	PropertyLocation *string `field:"optional" json:"propertyLocation" yaml:"propertyLocation"`
	// The data type of this property.
	//
	// Must be SECRET for secret properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html#cfn-glue-connectiontype-secretconnectorproperty-propertytype
	//
	PropertyType *string `field:"optional" json:"propertyType" yaml:"propertyType"`
	// Indicates whether the property is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html#cfn-glue-connectiontype-secretconnectorproperty-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

