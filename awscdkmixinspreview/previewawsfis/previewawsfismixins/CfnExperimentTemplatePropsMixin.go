package previewawsfismixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsfis/previewawsfismixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an experiment template.
//
// An experiment template includes the following components:
//
// - *Targets* : A target can be a specific resource in your AWS environment, or one or more resources that match criteria that you specify, for example, resources that have specific tags.
// - *Actions* : The actions to carry out on the target. You can specify multiple actions, the duration of each action, and when to start each action during an experiment.
// - *Stop conditions* : If a stop condition is triggered while an experiment is running, the experiment is automatically stopped. You can define a stop condition as a CloudWatch alarm.
//
// For more information, see [Experiment templates](https://docs.aws.amazon.com/fis/latest/userguide/experiment-templates.html) in the *AWS Fault Injection Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var cloudWatchLogsConfiguration interface{}
//   var s3Configuration interface{}
//
//   cfnExperimentTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnExperimentTemplatePropsMixin(&CfnExperimentTemplateMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-experimenttemplate.html
//
type CfnExperimentTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnExperimentTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnExperimentTemplatePropsMixin
type jsiiProxy_CfnExperimentTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnExperimentTemplatePropsMixin) Props() *CfnExperimentTemplateMixinProps {
	var returns *CfnExperimentTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::FIS::ExperimentTemplate`.
func NewCfnExperimentTemplatePropsMixin(props *CfnExperimentTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnExperimentTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnExperimentTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnExperimentTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fis.mixins.CfnExperimentTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::FIS::ExperimentTemplate`.
func NewCfnExperimentTemplatePropsMixin_Override(c CfnExperimentTemplatePropsMixin, props *CfnExperimentTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fis.mixins.CfnExperimentTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnExperimentTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExperimentTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_fis.mixins.CfnExperimentTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExperimentTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_fis.mixins.CfnExperimentTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnExperimentTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnExperimentTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

