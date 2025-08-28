package awsquicksight


// A transform operation that filters rows based on a condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterOperationProperty := &FilterOperationProperty{
//   	ConditionExpression: jsii.String("conditionExpression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filteroperation.html
//
type CfnDataSet_FilterOperationProperty struct {
	// An expression that must evaluate to a Boolean value.
	//
	// Rows for which the expression evaluates to true are kept in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-filteroperation.html#cfn-quicksight-dataset-filteroperation-conditionexpression
	//
	ConditionExpression *string `field:"optional" json:"conditionExpression" yaml:"conditionExpression"`
}

