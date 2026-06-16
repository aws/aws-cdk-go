package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::BedrockAgentCore::ConfigurationBundle Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnConfigurationBundlePropsMixin := awscdkcfnpropertymixins.Aws_bedrockagentcore.NewCfnConfigurationBundlePropsMixin(&CfnConfigurationBundleMixinProps{
//   	BranchName: jsii.String("branchName"),
//   	BundleName: jsii.String("bundleName"),
//   	CommitMessage: jsii.String("commitMessage"),
//   	Components: map[string]interface{}{
//   		"componentsKey": &ComponentConfigurationProperty{
//   			"configuration": configuration,
//   		},
//   	},
//   	CreatedBy: &VersionCreatedBySourceProperty{
//   		Arn: jsii.String("arn"),
//   		Name: jsii.String("name"),
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html
//
type CfnConfigurationBundlePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnConfigurationBundleMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfigurationBundlePropsMixin
type jsiiProxy_CfnConfigurationBundlePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnConfigurationBundlePropsMixin) Props() *CfnConfigurationBundleMixinProps {
	var returns *CfnConfigurationBundleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationBundlePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::ConfigurationBundle`.
func NewCfnConfigurationBundlePropsMixin(props *CfnConfigurationBundleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnConfigurationBundlePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfigurationBundlePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfigurationBundlePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnConfigurationBundlePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::ConfigurationBundle`.
func NewCfnConfigurationBundlePropsMixin_Override(c CfnConfigurationBundlePropsMixin, props *CfnConfigurationBundleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnConfigurationBundlePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnConfigurationBundlePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfigurationBundlePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnConfigurationBundlePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationBundlePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnConfigurationBundlePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationBundlePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConfigurationBundlePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

