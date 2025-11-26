package previewawsssmmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsssm/previewawsssmmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SSM::Parameter` resource creates an SSM parameter in AWS Systems Manager Parameter Store.
//
// > To create an SSM parameter, you must have the AWS Identity and Access Management ( IAM ) permissions `ssm:PutParameter` and `ssm:AddTagsToResource` . On stack creation, AWS CloudFormation adds the following three tags to the parameter: `aws:cloudformation:stack-name` , `aws:cloudformation:logical-id` , and `aws:cloudformation:stack-id` , in addition to any custom tags you specify.
// >
// > To add, update, or remove tags during stack update, you must have IAM permissions for both `ssm:AddTagsToResource` and `ssm:RemoveTagsFromResource` . For more information, see [Managing access using policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/security-iam.html#security_iam_access-manage) in the *AWS Systems Manager User Guide* .
//
// For information about valid values for parameters, see [About requirements and constraints for parameter names](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-su-create.html#sysman-parameter-name-constraints) in the *AWS Systems Manager User Guide* and [PutParameter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PutParameter.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnParameterPropsMixin := awscdkmixinspreview.Mixins.NewCfnParameterPropsMixin(&CfnParameterMixinProps{
//   	AllowedPattern: jsii.String("allowedPattern"),
//   	DataType: jsii.String("dataType"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Policies: jsii.String("policies"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Tier: jsii.String("tier"),
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html
//
type CfnParameterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnParameterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnParameterPropsMixin
type jsiiProxy_CfnParameterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnParameterPropsMixin) Props() *CfnParameterMixinProps {
	var returns *CfnParameterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParameterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSM::Parameter`.
func NewCfnParameterPropsMixin(props *CfnParameterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnParameterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnParameterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnParameterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnParameterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSM::Parameter`.
func NewCfnParameterPropsMixin_Override(c CfnParameterPropsMixin, props *CfnParameterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnParameterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnParameterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnParameterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnParameterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnParameterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnParameterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnParameterPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnParameterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

