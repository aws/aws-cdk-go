package awsglue


// Defines a property configuration for connection types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectorPropertyProperty := &ConnectorPropertyProperty{
//   	AllowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	DefaultValue: jsii.String("defaultValue"),
//   	KeyOverride: jsii.String("keyOverride"),
//   	Name: jsii.String("name"),
//   	PropertyLocation: jsii.String("propertyLocation"),
//   	PropertyType: jsii.String("propertyType"),
//   	Required: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html
//
type CfnConnectionTypePropsMixin_ConnectorPropertyProperty struct {
	// A list of allowed values for the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-allowedvalues
	//
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// The default value for the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// A key name to use when sending this property in API requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-keyoverride
	//
	KeyOverride *string `field:"optional" json:"keyOverride" yaml:"keyOverride"`
	// The name of the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies where this property should be included in REST requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-propertylocation
	//
	PropertyLocation *string `field:"optional" json:"propertyLocation" yaml:"propertyLocation"`
	// The data type of this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-propertytype
	//
	PropertyType *string `field:"optional" json:"propertyType" yaml:"propertyType"`
	// Indicates whether the property is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

