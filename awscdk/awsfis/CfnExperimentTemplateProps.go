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
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
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
//   	LogConfiguration: &ExperimentTemplateLogConfigurationProperty{
//   		LogSchemaVersion: jsii.Number(123),
//
//   		// the properties below are optional
//   		CloudWatchLogsConfiguration: cloudWatchLogsConfiguration,
//   		S3Configuration: s3Configuration,
//   	},
//   }
//
type CfnExperimentTemplateProps struct {
	// A description for the experiment template.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The stop conditions.
	StopConditions interface{} `field:"required" json:"stopConditions" yaml:"stopConditions"`
	// The tags to apply to the experiment template.
	Tags *map[string]*string `field:"required" json:"tags" yaml:"tags"`
	// The targets for the experiment.
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// The actions for the experiment.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The configuration for experiment logging.
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
}

