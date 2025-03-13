package awssesactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
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
//   lambdaProps := &LambdaProps{
//   	Function: function_,
//
//   	// the properties below are optional
//   	InvocationType: awscdk.Aws_ses_actions.LambdaInvocationType_EVENT,
//   	Topic: topic,
//   }
//
type LambdaProps struct {
	// The Lambda function to invoke.
	Function awslambda.IFunction `field:"required" json:"function" yaml:"function"`
	// The invocation type of the Lambda function.
	// Default: Event.
	//
	InvocationType LambdaInvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The SNS topic to notify when the Lambda action is taken.
	// Default: no notification.
	//
	Topic awssns.ITopic `field:"optional" json:"topic" yaml:"topic"`
}

