package awsec2


// Specifies whether detailed monitoring is enabled for an instance.
//
// For more information about detailed monitoring, see [Enable or turn off detailed monitoring for your instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch-new.html) in the *Amazon EC2 User Guide* .
//
// `Monitoring` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringProperty := &monitoringProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnLaunchTemplate_MonitoringProperty struct {
	// Specify `true` to enable detailed monitoring.
	//
	// Otherwise, basic monitoring is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

