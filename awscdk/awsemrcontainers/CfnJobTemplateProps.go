package awsemrcontainers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnJobTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ ConfigurationProperty
//
//   cfnJobTemplateProps := &CfnJobTemplateProps{
//   	JobTemplateData: &JobTemplateDataProperty{
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//   		JobDriver: &JobDriverProperty{
//   			SparkSqlJobDriver: &SparkSqlJobDriverProperty{
//   				EntryPoint: jsii.String("entryPoint"),
//   				SparkSqlParameters: jsii.String("sparkSqlParameters"),
//   			},
//   			SparkSubmitJobDriver: &SparkSubmitJobDriverProperty{
//   				EntryPoint: jsii.String("entryPoint"),
//
//   				// the properties below are optional
//   				EntryPointArguments: []*string{
//   					jsii.String("entryPointArguments"),
//   				},
//   				SparkSubmitParameters: jsii.String("sparkSubmitParameters"),
//   			},
//   		},
//   		ReleaseLabel: jsii.String("releaseLabel"),
//
//   		// the properties below are optional
//   		ConfigurationOverrides: &ParametricConfigurationOverridesProperty{
//   			ApplicationConfiguration: []interface{}{
//   				&ConfigurationProperty{
//   					Classification: jsii.String("classification"),
//
//   					// the properties below are optional
//   					Configurations: []interface{}{
//   						configurationProperty_,
//   					},
//   					Properties: map[string]*string{
//   						"propertiesKey": jsii.String("properties"),
//   					},
//   				},
//   			},
//   			MonitoringConfiguration: &ParametricMonitoringConfigurationProperty{
//   				CloudWatchMonitoringConfiguration: &ParametricCloudWatchMonitoringConfigurationProperty{
//   					LogGroupName: jsii.String("logGroupName"),
//   					LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   				},
//   				PersistentAppUi: jsii.String("persistentAppUi"),
//   				S3MonitoringConfiguration: &ParametricS3MonitoringConfigurationProperty{
//   					LogUri: jsii.String("logUri"),
//   				},
//   			},
//   		},
//   		JobTags: map[string]*string{
//   			"jobTagsKey": jsii.String("jobTags"),
//   		},
//   		ParameterConfiguration: map[string]interface{}{
//   			"parameterConfigurationKey": &TemplateParameterConfigurationProperty{
//   				"defaultValue": jsii.String("defaultValue"),
//   				"type": jsii.String("type"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-jobtemplate.html
//
type CfnJobTemplateProps struct {
	// The values of a StartJobRun API request used in job runs started using the job template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-jobtemplate.html#cfn-emrcontainers-jobtemplate-jobtemplatedata
	//
	JobTemplateData interface{} `field:"required" json:"jobTemplateData" yaml:"jobTemplateData"`
	// The specified name of the job template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-jobtemplate.html#cfn-emrcontainers-jobtemplate-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The tags that are associated with the job template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-jobtemplate.html#cfn-emrcontainers-jobtemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

