package awscdk


// Use the UpdatePolicy attribute to specify how AWS CloudFormation handles updates to the AWS::AutoScaling::AutoScalingGroup resource.
//
// AWS CloudFormation invokes one of three update policies depending on the type of change you make or whether a
// scheduled action is associated with the Auto Scaling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUpdatePolicy := &CfnUpdatePolicy{
//   	AutoScalingReplacingUpdate: &CfnAutoScalingReplacingUpdate{
//   		WillReplace: jsii.Boolean(false),
//   	},
//   	AutoScalingRollingUpdate: &CfnAutoScalingRollingUpdate{
//   		MaxBatchSize: jsii.Number(123),
//   		MinInstancesInService: jsii.Number(123),
//   		MinSuccessfulInstancesPercent: jsii.Number(123),
//   		PauseTime: jsii.String("pauseTime"),
//   		SuspendProcesses: []*string{
//   			jsii.String("suspendProcesses"),
//   		},
//   		WaitOnResourceSignals: jsii.Boolean(false),
//   	},
//   	AutoScalingScheduledAction: &CfnAutoScalingScheduledAction{
//   		IgnoreUnmodifiedGroupSizeProperties: jsii.Boolean(false),
//   	},
//   	CodeDeployLambdaAliasUpdate: &CfnCodeDeployLambdaAliasUpdate{
//   		ApplicationName: jsii.String("applicationName"),
//   		DeploymentGroupName: jsii.String("deploymentGroupName"),
//
//   		// the properties below are optional
//   		AfterAllowTrafficHook: jsii.String("afterAllowTrafficHook"),
//   		BeforeAllowTrafficHook: jsii.String("beforeAllowTrafficHook"),
//   	},
//   	EnableVersionUpgrade: jsii.Boolean(false),
//   	UseOnlineResharding: jsii.Boolean(false),
//   }
//
type CfnUpdatePolicy struct {
	// Specifies whether an Auto Scaling group and the instances it contains are replaced during an update.
	//
	// During replacement,
	// AWS CloudFormation retains the old group until it finishes creating the new one. If the update fails, AWS CloudFormation
	// can roll back to the old Auto Scaling group and delete the new Auto Scaling group.
	AutoScalingReplacingUpdate *CfnAutoScalingReplacingUpdate `field:"optional" json:"autoScalingReplacingUpdate" yaml:"autoScalingReplacingUpdate"`
	// To specify how AWS CloudFormation handles rolling updates for an Auto Scaling group, use the AutoScalingRollingUpdate policy.
	//
	// Rolling updates enable you to specify whether AWS CloudFormation updates instances that are in an Auto Scaling
	// group in batches or all at once.
	AutoScalingRollingUpdate *CfnAutoScalingRollingUpdate `field:"optional" json:"autoScalingRollingUpdate" yaml:"autoScalingRollingUpdate"`
	// To specify how AWS CloudFormation handles updates for the MinSize, MaxSize, and DesiredCapacity properties when the AWS::AutoScaling::AutoScalingGroup resource has an associated scheduled action, use the AutoScalingScheduledAction policy.
	AutoScalingScheduledAction *CfnAutoScalingScheduledAction `field:"optional" json:"autoScalingScheduledAction" yaml:"autoScalingScheduledAction"`
	// To perform an AWS CodeDeploy deployment when the version changes on an AWS::Lambda::Alias resource, use the CodeDeployLambdaAliasUpdate update policy.
	CodeDeployLambdaAliasUpdate *CfnCodeDeployLambdaAliasUpdate `field:"optional" json:"codeDeployLambdaAliasUpdate" yaml:"codeDeployLambdaAliasUpdate"`
	// To upgrade an Amazon ES domain to a new version of Elasticsearch rather than replacing the entire AWS::Elasticsearch::Domain resource, use the EnableVersionUpgrade update policy.
	EnableVersionUpgrade *bool `field:"optional" json:"enableVersionUpgrade" yaml:"enableVersionUpgrade"`
	// To modify a replication group's shards by adding or removing shards, rather than replacing the entire AWS::ElastiCache::ReplicationGroup resource, use the UseOnlineResharding update policy.
	UseOnlineResharding *bool `field:"optional" json:"useOnlineResharding" yaml:"useOnlineResharding"`
}

