package awsscheduler


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deadLetterConfigProperty := &deadLetterConfigProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnSchedule_DeadLetterConfigProperty struct {
	// `CfnSchedule.DeadLetterConfigProperty.Arn`.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

