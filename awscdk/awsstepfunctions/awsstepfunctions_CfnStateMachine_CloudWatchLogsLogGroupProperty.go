package awsstepfunctions


// Defines a CloudWatch log group.
//
// > For more information see [Standard Versus Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-standard-vs-express.html) in the AWS Step Functions Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsLogGroupProperty := &cloudWatchLogsLogGroupProperty{
//   	logGroupArn: jsii.String("logGroupArn"),
//   }
//
type CfnStateMachine_CloudWatchLogsLogGroupProperty struct {
	// The ARN of the the CloudWatch log group to which you want your logs emitted to.
	//
	// The ARN must end with `:*`.
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
}

