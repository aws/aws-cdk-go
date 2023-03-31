package awsiotsitewise


// Contains expression variable information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expressionVariableProperty := &expressionVariableProperty{
//   	name: jsii.String("name"),
//   	value: &variableValueProperty{
//   		propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   		// the properties below are optional
//   		hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   	},
//   }
//
type CfnAssetModel_ExpressionVariableProperty struct {
	// The friendly name of the variable to be used in the expression.
	//
	// The maximum length is 64 characters with the pattern `^[a-z][a-z0-9_]*$` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The variable that identifies an asset property from which to use values.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

