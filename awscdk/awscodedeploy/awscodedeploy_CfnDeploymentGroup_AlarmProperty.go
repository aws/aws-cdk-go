package awscodedeploy


// The `Alarm` property type specifies a CloudWatch alarm to use for an AWS CodeDeploy deployment group.
//
// The `Alarm` property of the [CodeDeploy DeploymentGroup AlarmConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html) property contains a list of `Alarm` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmProperty := &alarmProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnDeploymentGroup_AlarmProperty struct {
	// The name of the alarm.
	//
	// Maximum length is 255 characters. Each alarm name can be used only once in a list of alarms.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

