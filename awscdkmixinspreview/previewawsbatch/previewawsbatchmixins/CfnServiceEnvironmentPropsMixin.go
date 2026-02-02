package previewawsbatchmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbatch/previewawsbatchmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a service environment for running service jobs.
//
// Service environments define capacity limits for specific service types such as SageMaker Training jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceEnvironmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnServiceEnvironmentPropsMixin(&CfnServiceEnvironmentMixinProps{
//   	CapacityLimits: []interface{}{
//   		&CapacityLimitProperty{
//   			CapacityUnit: jsii.String("capacityUnit"),
//   			MaxCapacity: jsii.Number(123),
//   		},
//   	},
//   	ServiceEnvironmentName: jsii.String("serviceEnvironmentName"),
//   	ServiceEnvironmentType: jsii.String("serviceEnvironmentType"),
//   	State: jsii.String("state"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-serviceenvironment.html
//
type CfnServiceEnvironmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnServiceEnvironmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServiceEnvironmentPropsMixin
type jsiiProxy_CfnServiceEnvironmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnServiceEnvironmentPropsMixin) Props() *CfnServiceEnvironmentMixinProps {
	var returns *CfnServiceEnvironmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceEnvironmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Batch::ServiceEnvironment`.
func NewCfnServiceEnvironmentPropsMixin(props *CfnServiceEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnServiceEnvironmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServiceEnvironmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServiceEnvironmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnServiceEnvironmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Batch::ServiceEnvironment`.
func NewCfnServiceEnvironmentPropsMixin_Override(c CfnServiceEnvironmentPropsMixin, props *CfnServiceEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnServiceEnvironmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnServiceEnvironmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceEnvironmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnServiceEnvironmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServiceEnvironmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnServiceEnvironmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServiceEnvironmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnServiceEnvironmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

