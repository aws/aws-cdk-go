package awsemr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logTypes interface{}
//
//   monitoringConfigurationProperty := &MonitoringConfigurationProperty{
//   	CloudWatchLogConfiguration: &CloudWatchLogConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   		LogTypes: logTypes,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-monitoringconfiguration.html
//
type CfnCluster_MonitoringConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-monitoringconfiguration.html#cfn-emr-cluster-monitoringconfiguration-cloudwatchlogconfiguration
	//
	CloudWatchLogConfiguration interface{} `field:"optional" json:"cloudWatchLogConfiguration" yaml:"cloudWatchLogConfiguration"`
}

