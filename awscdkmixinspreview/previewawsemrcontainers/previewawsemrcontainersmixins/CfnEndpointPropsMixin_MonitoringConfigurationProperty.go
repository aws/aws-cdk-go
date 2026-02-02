package previewawsemrcontainersmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   monitoringConfigurationProperty := &MonitoringConfigurationProperty{
//   	CloudWatchMonitoringConfiguration: &CloudWatchMonitoringConfigurationProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   	},
//   	ContainerLogRotationConfiguration: &ContainerLogRotationConfigurationProperty{
//   		MaxFilesToKeep: jsii.Number(123),
//   		RotationSize: jsii.String("rotationSize"),
//   	},
//   	PersistentAppUi: jsii.String("persistentAppUi"),
//   	S3MonitoringConfiguration: &S3MonitoringConfigurationProperty{
//   		LogUri: jsii.String("logUri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-monitoringconfiguration.html
//
type CfnEndpointPropsMixin_MonitoringConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-monitoringconfiguration.html#cfn-emrcontainers-endpoint-monitoringconfiguration-cloudwatchmonitoringconfiguration
	//
	CloudWatchMonitoringConfiguration interface{} `field:"optional" json:"cloudWatchMonitoringConfiguration" yaml:"cloudWatchMonitoringConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-monitoringconfiguration.html#cfn-emrcontainers-endpoint-monitoringconfiguration-containerlogrotationconfiguration
	//
	ContainerLogRotationConfiguration interface{} `field:"optional" json:"containerLogRotationConfiguration" yaml:"containerLogRotationConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-monitoringconfiguration.html#cfn-emrcontainers-endpoint-monitoringconfiguration-persistentappui
	//
	PersistentAppUi *string `field:"optional" json:"persistentAppUi" yaml:"persistentAppUi"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-monitoringconfiguration.html#cfn-emrcontainers-endpoint-monitoringconfiguration-s3monitoringconfiguration
	//
	S3MonitoringConfiguration interface{} `field:"optional" json:"s3MonitoringConfiguration" yaml:"s3MonitoringConfiguration"`
}

