package awscodedeploy


// The `AlarmConfiguration` property type configures CloudWatch alarms for an AWS CodeDeploy deployment group.
//
// `AlarmConfiguration` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmConfigurationProperty := &alarmConfigurationProperty{
//   	alarms: []interface{}{
//   		&alarmProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   	ignorePollAlarmFailure: jsii.Boolean(false),
//   }
//
type CfnDeploymentGroup_AlarmConfigurationProperty struct {
	// A list of alarms configured for the deployment group.
	//
	// A maximum of 10 alarms can be added to a deployment group.
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// Indicates whether the alarm configuration is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates whether a deployment should continue if information about the current state of alarms cannot be retrieved from Amazon CloudWatch .
	//
	// The default value is `false` .
	//
	// - `true` : The deployment proceeds even if alarm status information can't be retrieved from CloudWatch .
	// - `false` : The deployment stops if alarm status information can't be retrieved from CloudWatch .
	IgnorePollAlarmFailure interface{} `field:"optional" json:"ignorePollAlarmFailure" yaml:"ignorePollAlarmFailure"`
}

