package awsglue


// Defines a property configuration for connection types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorPropertyProperty := &ConnectorPropertyProperty{
//   	Name: jsii.String("name"),
//   	PropertyType: jsii.String("propertyType"),
//   	Required: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	AllowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	DefaultValue: jsii.String("defaultValue"),
//   	KeyOverride: jsii.String("keyOverride"),
//   	PropertyLocation: jsii.String("propertyLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html
//
type CfnConnectionType_ConnectorPropertyProperty struct {
	// The name of the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-propertytype
	//
	PropertyType *string `field:"required" json:"propertyType" yaml:"propertyType"`
	// Indicates whether the property is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-required
	//
	Required interface{} `field:"required" json:"required" yaml:"required"`
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
	// Specifies where this property should be included in REST requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectorproperty.html#cfn-glue-connectiontype-connectorproperty-propertylocation
	//
	PropertyLocation *string `field:"optional" json:"propertyLocation" yaml:"propertyLocation"`
}

