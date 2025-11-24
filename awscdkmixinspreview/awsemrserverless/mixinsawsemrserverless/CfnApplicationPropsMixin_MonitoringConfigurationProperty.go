package mixinsawsemrserverless


// The configuration setting for monitoring logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   monitoringConfigurationProperty := &MonitoringConfigurationProperty{
//   	CloudWatchLoggingConfiguration: &CloudWatchLoggingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   		LogTypeMap: []interface{}{
//   			&LogTypeMapKeyValuePairProperty{
//   				Key: jsii.String("key"),
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	ManagedPersistenceMonitoringConfiguration: &ManagedPersistenceMonitoringConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	},
//   	PrometheusMonitoringConfiguration: &PrometheusMonitoringConfigurationProperty{
//   		RemoteWriteUrl: jsii.String("remoteWriteUrl"),
//   	},
//   	S3MonitoringConfiguration: &S3MonitoringConfigurationProperty{
//   		EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   		LogUri: jsii.String("logUri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-monitoringconfiguration.html
//
type CfnApplicationPropsMixin_MonitoringConfigurationProperty struct {
	// The Amazon CloudWatch configuration for monitoring logs.
	//
	// You can configure your jobs to send log information to CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-monitoringconfiguration.html#cfn-emrserverless-application-monitoringconfiguration-cloudwatchloggingconfiguration
	//
	CloudWatchLoggingConfiguration interface{} `field:"optional" json:"cloudWatchLoggingConfiguration" yaml:"cloudWatchLoggingConfiguration"`
	// The managed log persistence configuration for a job run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-monitoringconfiguration.html#cfn-emrserverless-application-monitoringconfiguration-managedpersistencemonitoringconfiguration
	//
	ManagedPersistenceMonitoringConfiguration interface{} `field:"optional" json:"managedPersistenceMonitoringConfiguration" yaml:"managedPersistenceMonitoringConfiguration"`
	// The monitoring configuration object you can configure to send metrics to Amazon Managed Service for Prometheus for a job run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-monitoringconfiguration.html#cfn-emrserverless-application-monitoringconfiguration-prometheusmonitoringconfiguration
	//
	PrometheusMonitoringConfiguration interface{} `field:"optional" json:"prometheusMonitoringConfiguration" yaml:"prometheusMonitoringConfiguration"`
	// The Amazon S3 configuration for monitoring log publishing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-monitoringconfiguration.html#cfn-emrserverless-application-monitoringconfiguration-s3monitoringconfiguration
	//
	S3MonitoringConfiguration interface{} `field:"optional" json:"s3MonitoringConfiguration" yaml:"s3MonitoringConfiguration"`
}

