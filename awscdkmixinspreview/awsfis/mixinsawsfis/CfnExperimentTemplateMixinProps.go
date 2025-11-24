package mixinsawsfis


// Properties for CfnExperimentTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var cloudWatchLogsConfiguration interface{}
//   var s3Configuration interface{}
//
//   cfnExperimentTemplateMixinProps := &CfnExperimentTemplateMixinProps{
//   	Actions: map[string]interface{}{
//   		"actionsKey": &ExperimentTemplateActionProperty{
//   			"actionId": jsii.String("actionId"),
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
//   	Description: jsii.String("description"),
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
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   		PostExperimentDuration: jsii.String("postExperimentDuration"),
//   		PreExperimentDuration: jsii.String("preExperimentDuration"),
//   	},
//   	LogConfiguration: &ExperimentTemplateLogConfigurationProperty{
//   		CloudWatchLogsConfiguration: cloudWatchLogsConfiguration,
//   		LogSchemaVersion: jsii.Number(123),
//   		S3Configuration: s3Configuration,
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	StopConditions: []interface{}{
//   		&ExperimentTemplateStopConditionProperty{
//   			Source: jsii.String("source"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Targets: map[string]interface{}{
//   		"targetsKey": &ExperimentTemplateTargetProperty{
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
//   			"resourceType": jsii.String("resourceType"),
//   			"selectionMode": jsii.String("selectionMode"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html
//
type CfnExperimentTemplateMixinProps struct {
	// The actions for the experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The description for the experiment template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
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
	// The Amazon Resource Name (ARN) of an IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The stop conditions for the experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-stopconditions
	//
	StopConditions interface{} `field:"optional" json:"stopConditions" yaml:"stopConditions"`
	// The tags for the experiment template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The targets for the experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html#cfn-fis-experimenttemplate-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

