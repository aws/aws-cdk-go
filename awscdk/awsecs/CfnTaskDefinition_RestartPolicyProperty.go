package awsecs


// You can enable a restart policy for each container defined in your task definition, to overcome transient failures faster and maintain task availability.
//
// When you enable a restart policy for a container, Amazon ECS can restart the container if it exits, without needing to replace the task. For more information, see [Restart individual containers in Amazon ECS tasks with container restart policies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-restart-policy.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restartPolicyProperty := &RestartPolicyProperty{
//   	Enabled: jsii.Boolean(false),
//   	IgnoredExitCodes: []interface{}{
//   		jsii.Number(123),
//   	},
//   	RestartAttemptPeriod: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-restartpolicy.html
//
type CfnTaskDefinition_RestartPolicyProperty struct {
	// Specifies whether a restart policy is enabled for the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-restartpolicy.html#cfn-ecs-taskdefinition-restartpolicy-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of exit codes that Amazon ECS will ignore and not attempt a restart on.
	//
	// You can specify a maximum of 50 container exit codes. By default, Amazon ECS does not ignore any exit codes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-restartpolicy.html#cfn-ecs-taskdefinition-restartpolicy-ignoredexitcodes
	//
	IgnoredExitCodes interface{} `field:"optional" json:"ignoredExitCodes" yaml:"ignoredExitCodes"`
	// A period of time (in seconds) that the container must run for before a restart can be attempted.
	//
	// A container can be restarted only once every `restartAttemptPeriod` seconds. If a container isn't able to run for this time period and exits early, it will not be restarted. You can set a minimum `restartAttemptPeriod` of 60 seconds and a maximum `restartAttemptPeriod` of 1800 seconds. By default, a container must run for 300 seconds before it can be restarted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-restartpolicy.html#cfn-ecs-taskdefinition-restartpolicy-restartattemptperiod
	//
	RestartAttemptPeriod *float64 `field:"optional" json:"restartAttemptPeriod" yaml:"restartAttemptPeriod"`
}

