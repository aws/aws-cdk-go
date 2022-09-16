package awssesactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Construction properties for a Lambda action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var topic topic
//
//   lambdaProps := &lambdaProps{
//   	function: function_,
//
//   	// the properties below are optional
//   	invocationType: awscdk.Aws_ses_actions.lambdaInvocationType_EVENT,
//   	topic: topic,
//   }
//
// Experimental.
type LambdaProps struct {
	// The Lambda function to invoke.
	// Experimental.
	Function awslambda.IFunction `field:"required" json:"function" yaml:"function"`
	// The invocation type of the Lambda function.
	// Experimental.
	InvocationType LambdaInvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The SNS topic to notify when the Lambda action is taken.
	// Experimental.
	Topic awssns.ITopic `field:"optional" json:"topic" yaml:"topic"`
}

