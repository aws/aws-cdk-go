package awsconnect


// Contains information about a queue for a quick connect.
//
// The contact flow must be of type Transfer to Queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueQuickConnectConfigProperty := &queueQuickConnectConfigProperty{
//   	contactFlowArn: jsii.String("contactFlowArn"),
//   	queueArn: jsii.String("queueArn"),
//   }
//
type CfnQuickConnect_QueueQuickConnectConfigProperty struct {
	// The Amazon Resource Name (ARN) of the contact flow.
	ContactFlowArn *string `field:"required" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The Amazon Resource Name (ARN) of the queue.
	QueueArn *string `field:"required" json:"queueArn" yaml:"queueArn"`
}

