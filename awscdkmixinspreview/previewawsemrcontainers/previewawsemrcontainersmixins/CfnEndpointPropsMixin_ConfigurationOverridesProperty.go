package previewawsemrcontainersmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var eMREKSConfigurationProperty_ EMREKSConfigurationProperty
//
//   configurationOverridesProperty := &ConfigurationOverridesProperty{
//   	ApplicationConfiguration: []interface{}{
//   		&EMREKSConfigurationProperty{
//   			Classification: jsii.String("classification"),
//   			Configurations: []interface{}{
//   				eMREKSConfigurationProperty_,
//   			},
//   			Properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	MonitoringConfiguration: &MonitoringConfigurationProperty{
//   		CloudWatchMonitoringConfiguration: &CloudWatchMonitoringConfigurationProperty{
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   		},
//   		ContainerLogRotationConfiguration: &ContainerLogRotationConfigurationProperty{
//   			MaxFilesToKeep: jsii.Number(123),
//   			RotationSize: jsii.String("rotationSize"),
//   		},
//   		PersistentAppUi: jsii.String("persistentAppUi"),
//   		S3MonitoringConfiguration: &S3MonitoringConfigurationProperty{
//   			LogUri: jsii.String("logUri"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-configurationoverrides.html
//
type CfnEndpointPropsMixin_ConfigurationOverridesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-configurationoverrides.html#cfn-emrcontainers-endpoint-configurationoverrides-applicationconfiguration
	//
	ApplicationConfiguration interface{} `field:"optional" json:"applicationConfiguration" yaml:"applicationConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-configurationoverrides.html#cfn-emrcontainers-endpoint-configurationoverrides-monitoringconfiguration
	//
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
}

