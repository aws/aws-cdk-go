package awsemrcontainers


// The values of a StartJobRun API request used in job runs started using the job template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var configurationProperty_ ConfigurationProperty
//
//   jobTemplateDataProperty := &JobTemplateDataProperty{
//   	ConfigurationOverrides: &ParametricConfigurationOverridesProperty{
//   		ApplicationConfiguration: []interface{}{
//   			&ConfigurationProperty{
//   				Classification: jsii.String("classification"),
//   				Configurations: []interface{}{
//   					configurationProperty_,
//   				},
//   				Properties: map[string]*string{
//   					"propertiesKey": jsii.String("properties"),
//   				},
//   			},
//   		},
//   		MonitoringConfiguration: &ParametricMonitoringConfigurationProperty{
//   			CloudWatchMonitoringConfiguration: &ParametricCloudWatchMonitoringConfigurationProperty{
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   			},
//   			PersistentAppUi: jsii.String("persistentAppUi"),
//   			S3MonitoringConfiguration: &ParametricS3MonitoringConfigurationProperty{
//   				LogUri: jsii.String("logUri"),
//   			},
//   		},
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	JobDriver: &JobDriverProperty{
//   		SparkSqlJobDriver: &SparkSqlJobDriverProperty{
//   			EntryPoint: jsii.String("entryPoint"),
//   			SparkSqlParameters: jsii.String("sparkSqlParameters"),
//   		},
//   		SparkSubmitJobDriver: &SparkSubmitJobDriverProperty{
//   			EntryPoint: jsii.String("entryPoint"),
//   			EntryPointArguments: []*string{
//   				jsii.String("entryPointArguments"),
//   			},
//   			SparkSubmitParameters: jsii.String("sparkSubmitParameters"),
//   		},
//   	},
//   	JobTags: map[string]*string{
//   		"jobTagsKey": jsii.String("jobTags"),
//   	},
//   	ParameterConfiguration: map[string]interface{}{
//   		"parameterConfigurationKey": &TemplateParameterConfigurationProperty{
//   			"defaultValue": jsii.String("defaultValue"),
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	ReleaseLabel: jsii.String("releaseLabel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-jobtemplatedata.html
//
type CfnJobTemplatePropsMixin_JobTemplateDataProperty struct {
	// A configuration specification to be used to override existing configurations.
	//
	// This data type allows job template parameters to be specified within.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-jobtemplatedata.html#cfn-emrcontainers-jobtemplate-jobtemplatedata-configurationoverrides
	//
	ConfigurationOverrides interface{} `field:"optional" json:"configurationOverrides" yaml:"configurationOverrides"`
	// The execution role ARN of the job run, or a template parameter reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-jobtemplatedata.html#cfn-emrcontainers-jobtemplate-jobtemplatedata-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Specify the driver that the job runs on.
	//
	// Exactly one of the two available job drivers is required, either SparkSqlJobDriver or SparkSubmitJobDriver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-jobtemplatedata.html#cfn-emrcontainers-jobtemplate-jobtemplatedata-jobdriver
	//
	JobDriver interface{} `field:"optional" json:"jobDriver" yaml:"jobDriver"`
	// The tags assigned to jobs started using the job template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-jobtemplatedata.html#cfn-emrcontainers-jobtemplate-jobtemplatedata-jobtags
	//
	JobTags interface{} `field:"optional" json:"jobTags" yaml:"jobTags"`
	// The configuration of parameters existing in the job template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-jobtemplatedata.html#cfn-emrcontainers-jobtemplate-jobtemplatedata-parameterconfiguration
	//
	ParameterConfiguration interface{} `field:"optional" json:"parameterConfiguration" yaml:"parameterConfiguration"`
	// The release version of Amazon EMR, or a template parameter reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-jobtemplatedata.html#cfn-emrcontainers-jobtemplate-jobtemplatedata-releaselabel
	//
	ReleaseLabel *string `field:"optional" json:"releaseLabel" yaml:"releaseLabel"`
}

