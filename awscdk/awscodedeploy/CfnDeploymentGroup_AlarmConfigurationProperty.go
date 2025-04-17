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
//   alarmConfigurationProperty := &AlarmConfigurationProperty{
//   	Alarms: []interface{}{
//   		&AlarmProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Enabled: jsii.Boolean(false),
//   	IgnorePollAlarmFailure: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html
//
type CfnDeploymentGroup_AlarmConfigurationProperty struct {
	// A list of alarms configured for the deployment or deployment group.
	//
	// A maximum of 10 alarms can be added.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-alarms
	//
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// Indicates whether the alarm configuration is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates whether a deployment should continue if information about the current state of alarms cannot be retrieved from Amazon CloudWatch .
	//
	// The default value is `false` .
	//
	// - `true` : The deployment proceeds even if alarm status information can't be retrieved from CloudWatch .
	// - `false` : The deployment stops if alarm status information can't be retrieved from CloudWatch .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-ignorepollalarmfailure
	//
	IgnorePollAlarmFailure interface{} `field:"optional" json:"ignorePollAlarmFailure" yaml:"ignorePollAlarmFailure"`
}

