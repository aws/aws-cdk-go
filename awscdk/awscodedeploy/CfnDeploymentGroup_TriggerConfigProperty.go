package awscodedeploy


// Information about notification triggers for the deployment group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerConfigProperty := &TriggerConfigProperty{
//   	TriggerEvents: []*string{
//   		jsii.String("triggerEvents"),
//   	},
//   	TriggerName: jsii.String("triggerName"),
//   	TriggerTargetArn: jsii.String("triggerTargetArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html
//
type CfnDeploymentGroup_TriggerConfigProperty struct {
	// The event type or types that trigger notifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggerevents
	//
	TriggerEvents *[]*string `field:"optional" json:"triggerEvents" yaml:"triggerEvents"`
	// The name of the notification trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggername
	//
	TriggerName *string `field:"optional" json:"triggerName" yaml:"triggerName"`
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic through which notifications about deployment or instance events are sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggertargetarn
	//
	TriggerTargetArn *string `field:"optional" json:"triggerTargetArn" yaml:"triggerTargetArn"`
}

