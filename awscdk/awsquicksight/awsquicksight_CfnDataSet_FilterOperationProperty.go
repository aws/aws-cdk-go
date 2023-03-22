package awsquicksight


// A transform operation that filters rows based on a condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterOperationProperty := &filterOperationProperty{
//   	conditionExpression: jsii.String("conditionExpression"),
//   }
//
type CfnDataSet_FilterOperationProperty struct {
	// An expression that must evaluate to a Boolean value.
	//
	// Rows for which the expression evaluates to true are kept in the dataset.
	ConditionExpression *string `field:"required" json:"conditionExpression" yaml:"conditionExpression"`
}

