package awslex


// Provides an expression that evaluates to true or false.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionProperty := &ConditionProperty{
//   	ExpressionString: jsii.String("expressionString"),
//   }
//
type CfnBot_ConditionProperty struct {
	// The expression string that is evaluated.
	ExpressionString *string `field:"required" json:"expressionString" yaml:"expressionString"`
}

