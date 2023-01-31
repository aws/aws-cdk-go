package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionProperty := &conditionProperty{
//   	expressionString: jsii.String("expressionString"),
//   }
//
type CfnBot_ConditionProperty struct {
	// `CfnBot.ConditionProperty.ExpressionString`.
	ExpressionString *string `field:"required" json:"expressionString" yaml:"expressionString"`
}

