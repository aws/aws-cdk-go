package awsemrserverless


// The configuration setting for monitoring.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringConfigurationProperty := &MonitoringConfigurationProperty{
//   	ManagedPersistenceMonitoringConfiguration: &ManagedPersistenceMonitoringConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	},
//   	S3MonitoringConfiguration: &S3MonitoringConfigurationProperty{
//   		EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   		LogUri: jsii.String("logUri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-monitoringconfiguration.html
//
type CfnApplication_MonitoringConfigurationProperty struct {
	// The managed log persistence configuration for a job run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-monitoringconfiguration.html#cfn-emrserverless-application-monitoringconfiguration-managedpersistencemonitoringconfiguration
	//
	ManagedPersistenceMonitoringConfiguration interface{} `field:"optional" json:"managedPersistenceMonitoringConfiguration" yaml:"managedPersistenceMonitoringConfiguration"`
	// The Amazon S3 configuration for monitoring log publishing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-monitoringconfiguration.html#cfn-emrserverless-application-monitoringconfiguration-s3monitoringconfiguration
	//
	S3MonitoringConfiguration interface{} `field:"optional" json:"s3MonitoringConfiguration" yaml:"s3MonitoringConfiguration"`
}

