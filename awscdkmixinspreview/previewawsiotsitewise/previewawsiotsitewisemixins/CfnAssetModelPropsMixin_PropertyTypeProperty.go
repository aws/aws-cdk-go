package previewawsiotsitewisemixins


// Contains a property type, which can be one of `attribute` , `measurement` , `metric` , or `transform` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   propertyTypeProperty := &PropertyTypeProperty{
//   	Attribute: &AttributeProperty{
//   		DefaultValue: jsii.String("defaultValue"),
//   	},
//   	Metric: &MetricProperty{
//   		Expression: jsii.String("expression"),
//   		Variables: []interface{}{
//   			&ExpressionVariableProperty{
//   				Name: jsii.String("name"),
//   				Value: &VariableValueProperty{
//   					HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   					HierarchyId: jsii.String("hierarchyId"),
//   					HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   					PropertyExternalId: jsii.String("propertyExternalId"),
//   					PropertyId: jsii.String("propertyId"),
//   					PropertyLogicalId: jsii.String("propertyLogicalId"),
//   					PropertyPath: []interface{}{
//   						&PropertyPathDefinitionProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Window: &MetricWindowProperty{
//   			Tumbling: &TumblingWindowProperty{
//   				Interval: jsii.String("interval"),
//   				Offset: jsii.String("offset"),
//   			},
//   		},
//   	},
//   	Transform: &TransformProperty{
//   		Expression: jsii.String("expression"),
//   		Variables: []interface{}{
//   			&ExpressionVariableProperty{
//   				Name: jsii.String("name"),
//   				Value: &VariableValueProperty{
//   					HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   					HierarchyId: jsii.String("hierarchyId"),
//   					HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   					PropertyExternalId: jsii.String("propertyExternalId"),
//   					PropertyId: jsii.String("propertyId"),
//   					PropertyLogicalId: jsii.String("propertyLogicalId"),
//   					PropertyPath: []interface{}{
//   						&PropertyPathDefinitionProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	TypeName: jsii.String("typeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-propertytype.html
//
type CfnAssetModelPropsMixin_PropertyTypeProperty struct {
	// Specifies an asset attribute property.
	//
	// An attribute generally contains static information, such as the serial number of an [IIoT](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Internet_of_things#Industrial_applications) wind turbine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-propertytype.html#cfn-iotsitewise-assetmodel-propertytype-attribute
	//
	Attribute interface{} `field:"optional" json:"attribute" yaml:"attribute"`
	// Specifies an asset metric property.
	//
	// A metric contains a mathematical expression that uses aggregate functions to process all input data points over a time interval and output a single data point, such as to calculate the average hourly temperature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-propertytype.html#cfn-iotsitewise-assetmodel-propertytype-metric
	//
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// Specifies an asset transform property.
	//
	// A transform contains a mathematical expression that maps a property's data points from one form to another, such as a unit conversion from Celsius to Fahrenheit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-propertytype.html#cfn-iotsitewise-assetmodel-propertytype-transform
	//
	Transform interface{} `field:"optional" json:"transform" yaml:"transform"`
	// The type of property type, which can be one of `Attribute` , `Measurement` , `Metric` , or `Transform` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-propertytype.html#cfn-iotsitewise-assetmodel-propertytype-typename
	//
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
}

