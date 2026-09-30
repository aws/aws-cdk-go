package awsglue


// Defines a secret property configuration.
//
// SECRET-type properties cannot have DefaultValue or AllowedValues.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretConnectorPropertyProperty := &SecretConnectorPropertyProperty{
//   	Name: jsii.String("name"),
//   	PropertyType: jsii.String("propertyType"),
//   	Required: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	KeyOverride: jsii.String("keyOverride"),
//   	PropertyLocation: jsii.String("propertyLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html
//
type CfnConnectionType_SecretConnectorPropertyProperty struct {
	// The name of the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html#cfn-glue-connectiontype-secretconnectorproperty-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of this property.
	//
	// Must be SECRET for secret properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html#cfn-glue-connectiontype-secretconnectorproperty-propertytype
	//
	PropertyType *string `field:"required" json:"propertyType" yaml:"propertyType"`
	// Indicates whether the property is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html#cfn-glue-connectiontype-secretconnectorproperty-required
	//
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// A key name to use when sending this property in API requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html#cfn-glue-connectiontype-secretconnectorproperty-keyoverride
	//
	KeyOverride *string `field:"optional" json:"keyOverride" yaml:"keyOverride"`
	// Specifies where this property should be included in REST requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-secretconnectorproperty.html#cfn-glue-connectiontype-secretconnectorproperty-propertylocation
	//
	PropertyLocation *string `field:"optional" json:"propertyLocation" yaml:"propertyLocation"`
}

