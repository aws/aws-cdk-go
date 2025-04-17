package awsdatazone


// A filter expression in Amazon DataZone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterExpressionProperty := &FilterExpressionProperty{
//   	Expression: jsii.String("expression"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-filterexpression.html
//
type CfnDataSource_FilterExpressionProperty struct {
	// The search filter expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-filterexpression.html#cfn-datazone-datasource-filterexpression-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The search filter explresison type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-filterexpression.html#cfn-datazone-datasource-filterexpression-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

