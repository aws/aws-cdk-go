package awsiotsitewise


// Contains expression variable information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expressionVariableProperty := &ExpressionVariableProperty{
//   	Name: jsii.String("name"),
//   	Value: &VariableValueProperty{
//   		PropertyLogicalId: jsii.String("propertyLogicalId"),
//
//   		// the properties below are optional
//   		HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-expressionvariable.html
//
type CfnAssetModel_ExpressionVariableProperty struct {
	// The friendly name of the variable to be used in the expression.
	//
	// The maximum length is 64 characters with the pattern `^[a-z][a-z0-9_]*$` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-expressionvariable.html#cfn-iotsitewise-assetmodel-expressionvariable-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The variable that identifies an asset property from which to use values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-expressionvariable.html#cfn-iotsitewise-assetmodel-expressionvariable-value
	//
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

