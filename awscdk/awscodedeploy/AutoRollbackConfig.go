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
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("CodeDeployDeploymentGroup"), &ServerDeploymentGroupProps{
//   	Application: Application,
//   	DeploymentGroupName: jsii.String("MyDeploymentGroup"),
//   	AutoScalingGroups: []iAutoScalingGroup{
//   		asg,
//   	},
//   	// adds User Data that installs the CodeDeploy agent on your auto-scaling groups hosts
//   	// default: true
//   	InstallAgent: jsii.Boolean(true),
//   	// adds EC2 instances matching tags
//   	Ec2InstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
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
//   	OnPremiseInstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
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
//   	Alarms: []iAlarm{
//   		alarm,
//   	},
//   	// whether to ignore failure to fetch the status of alarms from CloudWatch
//   	// default: false
//   	IgnorePollAlarmsFailure: jsii.Boolean(false),
//   	// auto-rollback configuration
//   	AutoRollback: &AutoRollbackConfig{
//   		FailedDeployment: jsii.Boolean(true),
//   		 // default: true
//   		StoppedDeployment: jsii.Boolean(true),
//   		 // default: false
//   		DeploymentInAlarm: jsii.Boolean(true),
//   	},
//   })
//
// Experimental.
type AutoRollbackConfig struct {
	// Whether to automatically roll back a deployment during which one of the configured CloudWatch alarms for this Deployment Group went off.
	// Experimental.
	DeploymentInAlarm *bool `field:"optional" json:"deploymentInAlarm" yaml:"deploymentInAlarm"`
	// Whether to automatically roll back a deployment that fails.
	// Experimental.
	FailedDeployment *bool `field:"optional" json:"failedDeployment" yaml:"failedDeployment"`
	// Whether to automatically roll back a deployment that was manually stopped.
	// Experimental.
	StoppedDeployment *bool `field:"optional" json:"stoppedDeployment" yaml:"stoppedDeployment"`
}

