package awsiotsitewise


// Contains a property type, which can be one of `Attribute` , `Measurement` , `Metric` , or `Transform` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propertyTypeProperty := &propertyTypeProperty{
//   	typeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	attribute: &attributeProperty{
//   		defaultValue: jsii.String("defaultValue"),
//   	},
//   	metric: &metricProperty{
//   		expression: jsii.String("expression"),
//   		variables: []interface{}{
//   			&expressionVariableProperty{
//   				name: jsii.String("name"),
//   				value: &variableValueProperty{
//   					propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   					// the properties below are optional
//   					hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   				},
//   			},
//   		},
//   		window: &metricWindowProperty{
//   			tumbling: &tumblingWindowProperty{
//   				interval: jsii.String("interval"),
//
//   				// the properties below are optional
//   				offset: jsii.String("offset"),
//   			},
//   		},
//   	},
//   	transform: &transformProperty{
//   		expression: jsii.String("expression"),
//   		variables: []interface{}{
//   			&expressionVariableProperty{
//   				name: jsii.String("name"),
//   				value: &variableValueProperty{
//   					propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   					// the properties below are optional
//   					hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnAssetModel_PropertyTypeProperty struct {
	// The type of property type, which can be one of `Attribute` , `Measurement` , `Metric` , or `Transform` .
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// Specifies an asset attribute property.
	//
	// An attribute generally contains static information, such as the serial number of an [industrial IoT](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Internet_of_things#Industrial_applications) wind turbine.
	//
	// This is required if the `TypeName` is `Attribute` and has a `DefaultValue` .
	Attribute interface{} `field:"optional" json:"attribute" yaml:"attribute"`
	// Specifies an asset metric property.
	//
	// A metric contains a mathematical expression that uses aggregate functions to process all input data points over a time interval and output a single data point, such as to calculate the average hourly temperature.
	//
	// This is required if the `TypeName` is `Metric` .
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// Specifies an asset transform property.
	//
	// A transform contains a mathematical expression that maps a property's data points from one form to another, such as a unit conversion from Celsius to Fahrenheit.
	//
	// This is required if the `TypeName` is `Transform` .
	Transform interface{} `field:"optional" json:"transform" yaml:"transform"`
}

