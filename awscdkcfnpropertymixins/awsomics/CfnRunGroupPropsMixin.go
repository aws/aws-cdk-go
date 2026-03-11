package awsomics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a run group to limit the compute resources for the runs that are added to the group.
//
// Returns an ARN, ID, and tags for the run group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRunGroupPropsMixin := awscdkcfnpropertymixins.Aws_omics.NewCfnRunGroupPropsMixin(&CfnRunGroupMixinProps{
//   	MaxCpus: jsii.Number(123),
//   	MaxDuration: jsii.Number(123),
//   	MaxGpus: jsii.Number(123),
//   	MaxRuns: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-rungroup.html
//
type CfnRunGroupPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRunGroupMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRunGroupPropsMixin
type jsiiProxy_CfnRunGroupPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRunGroupPropsMixin) Props() *CfnRunGroupMixinProps {
	var returns *CfnRunGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRunGroupPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Omics::RunGroup`.
func NewCfnRunGroupPropsMixin(props *CfnRunGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRunGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRunGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRunGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_omics.CfnRunGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Omics::RunGroup`.
func NewCfnRunGroupPropsMixin_Override(c CfnRunGroupPropsMixin, props *CfnRunGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_omics.CfnRunGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRunGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRunGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_omics.CfnRunGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRunGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_omics.CfnRunGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRunGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRunGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

