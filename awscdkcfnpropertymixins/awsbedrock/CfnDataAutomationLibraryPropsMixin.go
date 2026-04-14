package awsbedrock

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Bedrock::DataAutomationLibrary.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDataAutomationLibraryPropsMixin := awscdkcfnpropertymixins.Aws_bedrock.NewCfnDataAutomationLibraryPropsMixin(&CfnDataAutomationLibraryMixinProps{
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsEncryptionContext: map[string]*string{
//   			"kmsEncryptionContextKey": jsii.String("kmsEncryptionContext"),
//   		},
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	LibraryDescription: jsii.String("libraryDescription"),
//   	LibraryName: jsii.String("libraryName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationlibrary.html
//
type CfnDataAutomationLibraryPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDataAutomationLibraryMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataAutomationLibraryPropsMixin
type jsiiProxy_CfnDataAutomationLibraryPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDataAutomationLibraryPropsMixin) Props() *CfnDataAutomationLibraryMixinProps {
	var returns *CfnDataAutomationLibraryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataAutomationLibraryPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::DataAutomationLibrary`.
func NewCfnDataAutomationLibraryPropsMixin(props *CfnDataAutomationLibraryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDataAutomationLibraryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataAutomationLibraryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataAutomationLibraryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnDataAutomationLibraryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::DataAutomationLibrary`.
func NewCfnDataAutomationLibraryPropsMixin_Override(c CfnDataAutomationLibraryPropsMixin, props *CfnDataAutomationLibraryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnDataAutomationLibraryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDataAutomationLibraryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataAutomationLibraryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnDataAutomationLibraryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataAutomationLibraryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnDataAutomationLibraryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataAutomationLibraryPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataAutomationLibraryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

