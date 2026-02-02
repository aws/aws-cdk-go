package previewawssecurityhubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssecurityhub/previewawssecurityhubmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SecurityHub::ConfigurationPolicy` resource creates a central configuration policy with the defined settings.
//
// Only the AWS Security Hub CSPM delegated administrator can create this resource in the home Region. For more information, see [Central configuration in Security Hub CSPM](https://docs.aws.amazon.com/securityhub/latest/userguide/central-configuration-intro.html) in the *AWS Security Hub CSPM User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnConfigurationPolicyPropsMixin(&CfnConfigurationPolicyMixinProps{
//   	ConfigurationPolicy: &PolicyProperty{
//   		SecurityHub: &SecurityHubPolicyProperty{
//   			EnabledStandardIdentifiers: []*string{
//   				jsii.String("enabledStandardIdentifiers"),
//   			},
//   			SecurityControlsConfiguration: &SecurityControlsConfigurationProperty{
//   				DisabledSecurityControlIdentifiers: []*string{
//   					jsii.String("disabledSecurityControlIdentifiers"),
//   				},
//   				EnabledSecurityControlIdentifiers: []*string{
//   					jsii.String("enabledSecurityControlIdentifiers"),
//   				},
//   				SecurityControlCustomParameters: []interface{}{
//   					&SecurityControlCustomParameterProperty{
//   						Parameters: map[string]interface{}{
//   							"parametersKey": &ParameterConfigurationProperty{
//   								"value": &ParameterValueProperty{
//   									"boolean": jsii.Boolean(false),
//   									"double": jsii.Number(123),
//   									"enum": jsii.String("enum"),
//   									"enumList": []*string{
//   										jsii.String("enumList"),
//   									},
//   									"integer": jsii.Number(123),
//   									"integerList": []interface{}{
//   										jsii.Number(123),
//   									},
//   									"string": jsii.String("string"),
//   									"stringList": []*string{
//   										jsii.String("stringList"),
//   									},
//   								},
//   								"valueType": jsii.String("valueType"),
//   							},
//   						},
//   						SecurityControlId: jsii.String("securityControlId"),
//   					},
//   				},
//   			},
//   			ServiceEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html
//
type CfnConfigurationPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConfigurationPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfigurationPolicyPropsMixin
type jsiiProxy_CfnConfigurationPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConfigurationPolicyPropsMixin) Props() *CfnConfigurationPolicyMixinProps {
	var returns *CfnConfigurationPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityHub::ConfigurationPolicy`.
func NewCfnConfigurationPolicyPropsMixin(props *CfnConfigurationPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConfigurationPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigurationPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigurationPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnConfigurationPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityHub::ConfigurationPolicy`.
func NewCfnConfigurationPolicyPropsMixin_Override(c CfnConfigurationPolicyPropsMixin, props *CfnConfigurationPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnConfigurationPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConfigurationPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfigurationPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnConfigurationPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnConfigurationPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConfigurationPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

