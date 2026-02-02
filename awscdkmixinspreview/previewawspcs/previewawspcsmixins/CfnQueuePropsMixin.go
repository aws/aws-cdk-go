package previewawspcsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspcs/previewawspcsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an AWS PCS queue resource.
//
// For more information, see [Creating a queue in AWS PCS](https://docs.aws.amazon.com/pcs/latest/userguide/working-with_queues_create.html) in the *AWS PCS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnQueuePropsMixin := awscdkmixinspreview.Mixins.NewCfnQueuePropsMixin(&CfnQueueMixinProps{
//   	ClusterId: jsii.String("clusterId"),
//   	ComputeNodeGroupConfigurations: []interface{}{
//   		&ComputeNodeGroupConfigurationProperty{
//   			ComputeNodeGroupId: jsii.String("computeNodeGroupId"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	SlurmConfiguration: &SlurmConfigurationProperty{
//   		SlurmCustomSettings: []interface{}{
//   			&SlurmCustomSettingProperty{
//   				ParameterName: jsii.String("parameterName"),
//   				ParameterValue: jsii.String("parameterValue"),
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-queue.html
//
type CfnQueuePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnQueueMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnQueuePropsMixin
type jsiiProxy_CfnQueuePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnQueuePropsMixin) Props() *CfnQueueMixinProps {
	var returns *CfnQueueMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueuePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::PCS::Queue`.
func NewCfnQueuePropsMixin(props *CfnQueueMixinProps, options *mixins.CfnPropertyMixinOptions) CfnQueuePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnQueuePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnQueuePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::PCS::Queue`.
func NewCfnQueuePropsMixin_Override(c CfnQueuePropsMixin, props *CfnQueueMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnQueuePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnQueuePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnQueuePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnQueuePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnQueuePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

