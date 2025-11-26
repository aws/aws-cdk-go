package previewawsiotmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiot/previewawsiotmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a job template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var abortConfig interface{}
//   var jobExecutionsRolloutConfig interface{}
//   var presignedUrlConfig interface{}
//   var timeoutConfig interface{}
//
//   cfnJobTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnJobTemplatePropsMixin(&CfnJobTemplateMixinProps{
//   	AbortConfig: abortConfig,
//   	Description: jsii.String("description"),
//   	DestinationPackageVersions: []*string{
//   		jsii.String("destinationPackageVersions"),
//   	},
//   	Document: jsii.String("document"),
//   	DocumentSource: jsii.String("documentSource"),
//   	JobArn: jsii.String("jobArn"),
//   	JobExecutionsRetryConfig: &JobExecutionsRetryConfigProperty{
//   		RetryCriteriaList: []interface{}{
//   			&RetryCriteriaProperty{
//   				FailureType: jsii.String("failureType"),
//   				NumberOfRetries: jsii.Number(123),
//   			},
//   		},
//   	},
//   	JobExecutionsRolloutConfig: jobExecutionsRolloutConfig,
//   	JobTemplateId: jsii.String("jobTemplateId"),
//   	MaintenanceWindows: []interface{}{
//   		&MaintenanceWindowProperty{
//   			DurationInMinutes: jsii.Number(123),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	PresignedUrlConfig: presignedUrlConfig,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutConfig: timeoutConfig,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html
//
type CfnJobTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnJobTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnJobTemplatePropsMixin
type jsiiProxy_CfnJobTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnJobTemplatePropsMixin) Props() *CfnJobTemplateMixinProps {
	var returns *CfnJobTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::JobTemplate`.
func NewCfnJobTemplatePropsMixin(props *CfnJobTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnJobTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnJobTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJobTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnJobTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::JobTemplate`.
func NewCfnJobTemplatePropsMixin_Override(c CfnJobTemplatePropsMixin, props *CfnJobTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnJobTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnJobTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnJobTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnJobTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJobTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

