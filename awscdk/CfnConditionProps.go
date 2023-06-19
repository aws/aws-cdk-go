package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnConditionExpression iCfnConditionExpression
//
//   cfnConditionProps := &CfnConditionProps{
//   	Expression: cfnConditionExpression,
//   }
//
// Experimental.
type CfnConditionProps struct {
	// The expression that the condition will evaluate.
	// Experimental.
	Expression ICfnConditionExpression `field:"optional" json:"expression" yaml:"expression"`
}

