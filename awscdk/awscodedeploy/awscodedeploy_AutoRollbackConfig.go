package awscodedeploy


// The configuration for automatically rolling back deployments in a given Deployment Group.
//
// Example:
//   import autoscaling "github.com/aws/aws-cdk-go/awscdk"
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//   var application serverApplication
//   var asg autoScalingGroup
//   var alarm alarm
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("CodeDeployDeploymentGroup"), &serverDeploymentGroupProps{
//   	application: application,
//   	deploymentGroupName: jsii.String("MyDeploymentGroup"),
//   	autoScalingGroups: []iAutoScalingGroup{
//   		asg,
//   	},
//   	// adds User Data that installs the CodeDeploy agent on your auto-scaling groups hosts
//   	// default: true
//   	installAgent: jsii.Boolean(true),
//   	// adds EC2 instances matching tags
//   	ec2InstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
//   		// any instance with tags satisfying
//   		// key1=v1 or key1=v2 or key2 (any value) or value v3 (any key)
//   		// will match this group
//   		"key1": []*string{
//   			jsii.String("v1"),
//   			jsii.String("v2"),
//   		},
//   		"key2": []*string{
//   		},
//   		"": []*string{
//   			jsii.String("v3"),
//   		},
//   	}),
//   	// adds on-premise instances matching tags
//   	onPremiseInstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
//   		"key1": []*string{
//   			jsii.String("v1"),
//   			jsii.String("v2"),
//   		},
//   	}, map[string][]*string{
//   		"key2": []*string{
//   			jsii.String("v3"),
//   		},
//   	}),
//   	// CloudWatch alarms
//   	alarms: []iAlarm{
//   		alarm,
//   	},
//   	// whether to ignore failure to fetch the status of alarms from CloudWatch
//   	// default: false
//   	ignorePollAlarmsFailure: jsii.Boolean(false),
//   	// auto-rollback configuration
//   	autoRollback: &autoRollbackConfig{
//   		failedDeployment: jsii.Boolean(true),
//   		 // default: true
//   		stoppedDeployment: jsii.Boolean(true),
//   		 // default: false
//   		deploymentInAlarm: jsii.Boolean(true),
//   	},
//   })
//
type AutoRollbackConfig struct {
	// Whether to automatically roll back a deployment during which one of the configured CloudWatch alarms for this Deployment Group went off.
	DeploymentInAlarm *bool `field:"optional" json:"deploymentInAlarm" yaml:"deploymentInAlarm"`
	// Whether to automatically roll back a deployment that fails.
	FailedDeployment *bool `field:"optional" json:"failedDeployment" yaml:"failedDeployment"`
	// Whether to automatically roll back a deployment that was manually stopped.
	StoppedDeployment *bool `field:"optional" json:"stoppedDeployment" yaml:"stoppedDeployment"`
}

