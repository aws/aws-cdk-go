package awscodedeploy


// Information about notification triggers for the deployment group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerConfigProperty := &triggerConfigProperty{
//   	triggerEvents: []*string{
//   		jsii.String("triggerEvents"),
//   	},
//   	triggerName: jsii.String("triggerName"),
//   	triggerTargetArn: jsii.String("triggerTargetArn"),
//   }
//
type CfnDeploymentGroup_TriggerConfigProperty struct {
	// The event type or types that trigger notifications.
	TriggerEvents *[]*string `field:"optional" json:"triggerEvents" yaml:"triggerEvents"`
	// The name of the notification trigger.
	TriggerName *string `field:"optional" json:"triggerName" yaml:"triggerName"`
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic through which notifications about deployment or instance events are sent.
	TriggerTargetArn *string `field:"optional" json:"triggerTargetArn" yaml:"triggerTargetArn"`
}

