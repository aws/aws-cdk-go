package awsglue


// Configuration that defines the base URL and additional request parameters needed during connection creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionPropertiesConfigurationProperty := &ConnectionPropertiesConfigurationProperty{
//   	AdditionalRequestParameters: []interface{}{
//   		&ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   	},
//   	Url: &ConnectorPropertyProperty{
//   		Name: jsii.String("name"),
//   		PropertyType: jsii.String("propertyType"),
//   		Required: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		DefaultValue: jsii.String("defaultValue"),
//   		KeyOverride: jsii.String("keyOverride"),
//   		PropertyLocation: jsii.String("propertyLocation"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectionpropertiesconfiguration.html
//
type CfnConnectionType_ConnectionPropertiesConfigurationProperty struct {
	// Key-value pairs of additional request parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectionpropertiesconfiguration.html#cfn-glue-connectiontype-connectionpropertiesconfiguration-additionalrequestparameters
	//
	AdditionalRequestParameters interface{} `field:"optional" json:"additionalRequestParameters" yaml:"additionalRequestParameters"`
	// Defines a property configuration for connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-connectionpropertiesconfiguration.html#cfn-glue-connectiontype-connectionpropertiesconfiguration-url
	//
	Url interface{} `field:"optional" json:"url" yaml:"url"`
}

