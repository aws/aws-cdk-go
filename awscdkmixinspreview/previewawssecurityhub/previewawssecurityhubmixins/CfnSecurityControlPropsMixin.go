package previewawssecurityhubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssecurityhub/previewawssecurityhubmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SecurityHub::SecurityControl` resource specifies custom parameter values for an AWS Security Hub CSPM control.
//
// For a list of controls that support custom parameters, see [Security Hub CSPM controls reference](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-controls-reference.html) . You can also use this resource to specify the use of default parameter values for a control. For more information about custom parameters, see [Custom control parameters](https://docs.aws.amazon.com/securityhub/latest/userguide/custom-control-parameters.html) in the *AWS Security Hub CSPM User Guide* .
//
// Tags aren't supported for this resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityControlPropsMixin := awscdkmixinspreview.Mixins.NewCfnSecurityControlPropsMixin(&CfnSecurityControlMixinProps{
//   	LastUpdateReason: jsii.String("lastUpdateReason"),
//   	Parameters: map[string]interface{}{
//   		"parametersKey": &ParameterConfigurationProperty{
//   			"value": &ParameterValueProperty{
//   				"boolean": jsii.Boolean(false),
//   				"double": jsii.Number(123),
//   				"enum": jsii.String("enum"),
//   				"enumList": []*string{
//   					jsii.String("enumList"),
//   				},
//   				"integer": jsii.Number(123),
//   				"integerList": []interface{}{
//   					jsii.Number(123),
//   				},
//   				"string": jsii.String("string"),
//   				"stringList": []*string{
//   					jsii.String("stringList"),
//   				},
//   			},
//   			"valueType": jsii.String("valueType"),
//   		},
//   	},
//   	SecurityControlArn: jsii.String("securityControlArn"),
//   	SecurityControlId: jsii.String("securityControlId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-securitycontrol.html
//
type CfnSecurityControlPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSecurityControlMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSecurityControlPropsMixin
type jsiiProxy_CfnSecurityControlPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSecurityControlPropsMixin) Props() *CfnSecurityControlMixinProps {
	var returns *CfnSecurityControlMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityControlPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityHub::SecurityControl`.
func NewCfnSecurityControlPropsMixin(props *CfnSecurityControlMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSecurityControlPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSecurityControlPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSecurityControlPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnSecurityControlPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityHub::SecurityControl`.
func NewCfnSecurityControlPropsMixin_Override(c CfnSecurityControlPropsMixin, props *CfnSecurityControlMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnSecurityControlPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSecurityControlPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityControlPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnSecurityControlPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityControlPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnSecurityControlPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityControlPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSecurityControlPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

