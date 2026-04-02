package awsemr


// Contains CloudWatch log configuration metadata and settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   monitoringConfigurationProperty := &MonitoringConfigurationProperty{
//   	CloudWatchLogConfiguration: &CloudWatchLogConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   		LogTypes: map[string][]*string{
//   			"logTypesKey": []*string{
//   				jsii.String("logTypes"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-monitoringconfiguration.html
//
type CfnClusterPropsMixin_MonitoringConfigurationProperty struct {
	// Holds CloudWatch log configuration settings and metadata that specify settings like log files to monitor and where to send them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-monitoringconfiguration.html#cfn-emr-cluster-monitoringconfiguration-cloudwatchlogconfiguration
	//
	CloudWatchLogConfiguration interface{} `field:"optional" json:"cloudWatchLogConfiguration" yaml:"cloudWatchLogConfiguration"`
}

