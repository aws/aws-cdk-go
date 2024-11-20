package awsfis


// Properties for defining a `CfnExperimentTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudWatchLogsConfiguration interface{}
//   var s3Configuration interface{}
//
//   cfnExperimentTemplateProps := &CfnExperimentTemplateProps{
//   	Description: jsii.String("description"),
//   	RoleArn: jsii.String("roleArn"),
//   	StopConditions: []interface{}{
//   		&ExperimentTemplateStopConditionProperty{
//   			Source: jsii.String("source"),
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Targets: map[string]interface{}{
//   		"targetsKey": &ExperimentTemplateTargetProperty{
//   			"resourceType": jsii.String("resourceType"),
//   			"selectionMode": jsii.String("selectionMode"),
//
//   			// the properties below are optional
//   			"filters": []interface{}{
//   				&ExperimentTemplateTargetFilterProperty{
//   					"path": jsii.String("path"),
//   					"values": []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			"parameters": map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			"resourceArns": []*string{
//   				jsii.String("resourceArns"),
//   			},
//   			"resourceTags": map[string]*string{
//   				"resourceTagsKey": jsii.String("resourceTags"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Actions: map[string]interface{}{
//   		"actionsKey": &ExperimentTemplateActionProperty{
//   			"actionId": jsii.String("actionId"),
//
//   			// the properties below are optional
//   			"description": jsii.String("description"),
//   			"parameters": map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			"startAfter": []*string{
//   				jsii.String("startAfter"),
//   			},
//   			"targets": map[string]*string{
//   				"targetsKey": jsii.String("targets"),
//   			},
//   		},
//   	},
//   	ExperimentOptions: &ExperimentTemplateExperimentOptionsProperty{
//   		AccountTargeting: jsii.String("accountTargeting"),
//   		EmptyTargetResolutionMode: jsii.String("emptyTargetResolutionMode"),
//   	},
//   	ExperimentReportConfiguration: &ExperimentTemplateExperimentReportConfigurationProperty{
//   		DataSources: &DataSourcesProperty{
//   			CloudWatchDashboards: []interface{}{
//   				&CloudWatchDashboardProperty{
//   					DashboardIdentifier: jsii.String("dashboardIdentifier"),
//   				},
//   			},
//   		},
//   		Outputs: &OutputsProperty{
//   			ExperimentReportS3Configuration: &ExperimentReportS3ConfigurationProperty{
//   				BucketName: jsii.String("bucketName"),
//
//   				// the properties below are optional
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		PostExperimentDuration: jsii.String("postExperimentDuration"),
//   		PreExperimentDuration: jsii.String("preExperimentDuration"),
//   	},
//   	LogConfiguration: &ExperimentTemplateLogConfigurationProperty{
//   		LogSchemaVersion: jsii.Number(123),
//
//   		// the properties below are optional
//   		CloudWatchLogsConfiguration: cloudWatchLogsConfiguration,
//   		S3Configuration: s3Configuration,
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html
//
type CfnExperimentTemplateProps struct {
	// The description for the experiment template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of an IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The stop conditions for the experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-stopconditions
	//
	StopConditions interface{} `field:"required" json:"stopConditions" yaml:"stopConditions"`
	// The targets for the experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-targets
	//
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// The actions for the experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The experiment options for an experiment template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-experimentoptions
	//
	ExperimentOptions interface{} `field:"optional" json:"experimentOptions" yaml:"experimentOptions"`
	// Describes the report configuration for the experiment template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-experimentreportconfiguration
	//
	ExperimentReportConfiguration interface{} `field:"optional" json:"experimentReportConfiguration" yaml:"experimentReportConfiguration"`
	// The configuration for experiment logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The tags for the experiment template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

