package awsscheduler


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqsParametersProperty := &sqsParametersProperty{
//   	messageGroupId: jsii.String("messageGroupId"),
//   }
//
type CfnSchedule_SqsParametersProperty struct {
	// `CfnSchedule.SqsParametersProperty.MessageGroupId`.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

