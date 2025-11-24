package mixinsawsiotsitewise


// Contains an asset transform property.
//
// A transform is a one-to-one mapping of a property's data points from one form to another. For example, you can use a transform to convert a Celsius data stream to Fahrenheit by applying the transformation expression to each data point of the Celsius stream. A transform can only have a data type of `DOUBLE` and consume properties with data types of `INTEGER` or `DOUBLE` .
//
// For more information, see [Transforms](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#transforms) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformProperty := &TransformProperty{
//   	Expression: jsii.String("expression"),
//   	Variables: []interface{}{
//   		&ExpressionVariableProperty{
//   			Name: jsii.String("name"),
//   			Value: &VariableValueProperty{
//   				HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   				HierarchyId: jsii.String("hierarchyId"),
//   				HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   				PropertyExternalId: jsii.String("propertyExternalId"),
//   				PropertyId: jsii.String("propertyId"),
//   				PropertyLogicalId: jsii.String("propertyLogicalId"),
//   				PropertyPath: []interface{}{
//   					&PropertyPathDefinitionProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-transform.html
//
type CfnAssetModelPropsMixin_TransformProperty struct {
	// The mathematical expression that defines the transformation function.
	//
	// You can specify up to 10 variables per expression. You can specify up to 10 functions per expression.
	//
	// For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-transform.html#cfn-iotsitewise-assetmodel-transform-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The list of variables used in the expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-transform.html#cfn-iotsitewise-assetmodel-transform-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

