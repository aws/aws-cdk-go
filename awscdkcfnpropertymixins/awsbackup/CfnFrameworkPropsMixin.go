package awsbackup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a framework with one or more controls.
//
// A framework is a collection of controls that you can use to evaluate your backup practices. By using pre-built customizable controls to define your policies, you can evaluate whether your backup practices comply with your policies and which resources are not yet in compliance.
//
// For a sample CloudFormation template, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/bam-cfn-integration.html#bam-cfn-frameworks-template) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var controlScope interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnFrameworkPropsMixin := awscdkcfnpropertymixins.Aws_backup.NewCfnFrameworkPropsMixin(&CfnFrameworkMixinProps{
//   	FrameworkControls: []interface{}{
//   		&FrameworkControlProperty{
//   			ControlInputParameters: []interface{}{
//   				&ControlInputParameterProperty{
//   					ParameterName: jsii.String("parameterName"),
//   					ParameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   			ControlName: jsii.String("controlName"),
//   			ControlScope: controlScope,
//   		},
//   	},
//   	FrameworkDescription: jsii.String("frameworkDescription"),
//   	FrameworkName: jsii.String("frameworkName"),
//   	FrameworkTags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-framework.html
//
type CfnFrameworkPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFrameworkMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFrameworkPropsMixin
type jsiiProxy_CfnFrameworkPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFrameworkPropsMixin) Props() *CfnFrameworkMixinProps {
	var returns *CfnFrameworkMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFrameworkPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Backup::Framework`.
func NewCfnFrameworkPropsMixin(props *CfnFrameworkMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFrameworkPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFrameworkPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFrameworkPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnFrameworkPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Backup::Framework`.
func NewCfnFrameworkPropsMixin_Override(c CfnFrameworkPropsMixin, props *CfnFrameworkMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnFrameworkPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFrameworkPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFrameworkPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnFrameworkPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFrameworkPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_backup.CfnFrameworkPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFrameworkPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFrameworkPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

