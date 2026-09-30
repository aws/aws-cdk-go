package awsemrcontainers


// A configuration specification to be used to override existing configurations.
//
// This data type allows job template parameters to be specified within.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var configurationProperty_ ConfigurationProperty
//
//   parametricConfigurationOverridesProperty := &ParametricConfigurationOverridesProperty{
//   	ApplicationConfiguration: []interface{}{
//   		&ConfigurationProperty{
//   			Classification: jsii.String("classification"),
//   			Configurations: []interface{}{
//   				configurationProperty_,
//   			},
//   			Properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	MonitoringConfiguration: &ParametricMonitoringConfigurationProperty{
//   		CloudWatchMonitoringConfiguration: &ParametricCloudWatchMonitoringConfigurationProperty{
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   		},
//   		PersistentAppUi: jsii.String("persistentAppUi"),
//   		S3MonitoringConfiguration: &ParametricS3MonitoringConfigurationProperty{
//   			LogUri: jsii.String("logUri"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametricconfigurationoverrides.html
//
type CfnJobTemplatePropsMixin_ParametricConfigurationOverridesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametricconfigurationoverrides.html#cfn-emrcontainers-jobtemplate-parametricconfigurationoverrides-applicationconfiguration
	//
	ApplicationConfiguration interface{} `field:"optional" json:"applicationConfiguration" yaml:"applicationConfiguration"`
	// Configuration setting for monitoring.
	//
	// This data type allows job template parameters to be specified within.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametricconfigurationoverrides.html#cfn-emrcontainers-jobtemplate-parametricconfigurationoverrides-monitoringconfiguration
	//
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
}

