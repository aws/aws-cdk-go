package previewawsgluemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglue/previewawsgluemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::Job` resource specifies an AWS Glue job in the data catalog.
//
// For more information, see [Adding Jobs in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) and [Job Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-job.html#aws-glue-api-jobs-job-Job) in the *AWS Glue Developer Guide.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var defaultArguments interface{}
//   var nonOverridableArguments interface{}
//   var tags interface{}
//
//   cfnJobPropsMixin := awscdkmixinspreview.Mixins.NewCfnJobPropsMixin(&CfnJobMixinProps{
//   	AllocatedCapacity: jsii.Number(123),
//   	Command: &JobCommandProperty{
//   		Name: jsii.String("name"),
//   		PythonVersion: jsii.String("pythonVersion"),
//   		Runtime: jsii.String("runtime"),
//   		ScriptLocation: jsii.String("scriptLocation"),
//   	},
//   	Connections: &ConnectionsListProperty{
//   		Connections: []*string{
//   			jsii.String("connections"),
//   		},
//   	},
//   	DefaultArguments: defaultArguments,
//   	Description: jsii.String("description"),
//   	ExecutionClass: jsii.String("executionClass"),
//   	ExecutionProperty: &ExecutionPropertyProperty{
//   		MaxConcurrentRuns: jsii.Number(123),
//   	},
//   	GlueVersion: jsii.String("glueVersion"),
//   	JobMode: jsii.String("jobMode"),
//   	JobRunQueuingEnabled: jsii.Boolean(false),
//   	LogUri: jsii.String("logUri"),
//   	MaintenanceWindow: jsii.String("maintenanceWindow"),
//   	MaxCapacity: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	NonOverridableArguments: nonOverridableArguments,
//   	NotificationProperty: &NotificationPropertyProperty{
//   		NotifyDelayAfter: jsii.Number(123),
//   	},
//   	NumberOfWorkers: jsii.Number(123),
//   	Role: jsii.String("role"),
//   	SecurityConfiguration: jsii.String("securityConfiguration"),
//   	Tags: tags,
//   	Timeout: jsii.Number(123),
//   	WorkerType: jsii.String("workerType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html
//
type CfnJobPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnJobMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnJobPropsMixin
type jsiiProxy_CfnJobPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnJobPropsMixin) Props() *CfnJobMixinProps {
	var returns *CfnJobMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::Job`.
func NewCfnJobPropsMixin(props *CfnJobMixinProps, options *mixins.CfnPropertyMixinOptions) CfnJobPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnJobPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJobPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnJobPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::Job`.
func NewCfnJobPropsMixin_Override(c CfnJobPropsMixin, props *CfnJobMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnJobPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnJobPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnJobPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnJobPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJobPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnJobPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

