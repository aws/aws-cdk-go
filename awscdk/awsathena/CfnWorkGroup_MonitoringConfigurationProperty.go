package awsathena


// Contains the configuration settings for managed log persistence, delivering logs to Amazon S3 buckets, Amazon CloudWatch log groups etc.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringConfigurationProperty := &MonitoringConfigurationProperty{
//   	CloudWatchLoggingConfiguration: &CloudWatchLoggingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroup: jsii.String("logGroup"),
//   		LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   		LogTypes: map[string][]*string{
//   			"logTypesKey": []*string{
//   				jsii.String("logTypes"),
//   			},
//   		},
//   	},
//   	ManagedLoggingConfiguration: &ManagedLoggingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	S3LoggingConfiguration: &S3LoggingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		KmsKey: jsii.String("kmsKey"),
//   		LogLocation: jsii.String("logLocation"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-monitoringconfiguration.html
//
type CfnWorkGroup_MonitoringConfigurationProperty struct {
	// Configuration settings for delivering logs to Amazon CloudWatch log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-monitoringconfiguration.html#cfn-athena-workgroup-monitoringconfiguration-cloudwatchloggingconfiguration
	//
	CloudWatchLoggingConfiguration interface{} `field:"optional" json:"cloudWatchLoggingConfiguration" yaml:"cloudWatchLoggingConfiguration"`
	// Configuration settings for managed log persistence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-monitoringconfiguration.html#cfn-athena-workgroup-monitoringconfiguration-managedloggingconfiguration
	//
	ManagedLoggingConfiguration interface{} `field:"optional" json:"managedLoggingConfiguration" yaml:"managedLoggingConfiguration"`
	// Configuration settings for delivering logs to Amazon S3 buckets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-monitoringconfiguration.html#cfn-athena-workgroup-monitoringconfiguration-s3loggingconfiguration
	//
	S3LoggingConfiguration interface{} `field:"optional" json:"s3LoggingConfiguration" yaml:"s3LoggingConfiguration"`
}

