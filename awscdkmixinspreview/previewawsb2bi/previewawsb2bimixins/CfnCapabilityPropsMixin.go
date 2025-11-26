package previewawsb2bimixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsb2bi/previewawsb2bimixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Instantiates a capability based on the specified parameters.
//
// A trading capability contains the information required to transform incoming EDI documents into JSON or XML outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityPropsMixin := awscdkmixinspreview.Mixins.NewCfnCapabilityPropsMixin(&CfnCapabilityMixinProps{
//   	Configuration: &CapabilityConfigurationProperty{
//   		Edi: &EdiConfigurationProperty{
//   			CapabilityDirection: jsii.String("capabilityDirection"),
//   			InputLocation: &S3LocationProperty{
//   				BucketName: jsii.String("bucketName"),
//   				Key: jsii.String("key"),
//   			},
//   			OutputLocation: &S3LocationProperty{
//   				BucketName: jsii.String("bucketName"),
//   				Key: jsii.String("key"),
//   			},
//   			TransformerId: jsii.String("transformerId"),
//   			Type: &EdiTypeProperty{
//   				X12Details: &X12DetailsProperty{
//   					TransactionSet: jsii.String("transactionSet"),
//   					Version: jsii.String("version"),
//   				},
//   			},
//   		},
//   	},
//   	InstructionsDocuments: []interface{}{
//   		&S3LocationProperty{
//   			BucketName: jsii.String("bucketName"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html
//
type CfnCapabilityPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCapabilityMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCapabilityPropsMixin
type jsiiProxy_CfnCapabilityPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCapabilityPropsMixin) Props() *CfnCapabilityMixinProps {
	var returns *CfnCapabilityMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCapabilityPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::B2BI::Capability`.
func NewCfnCapabilityPropsMixin(props *CfnCapabilityMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCapabilityPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCapabilityPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCapabilityPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::B2BI::Capability`.
func NewCfnCapabilityPropsMixin_Override(c CfnCapabilityPropsMixin, props *CfnCapabilityMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCapabilityPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCapabilityPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCapabilityPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnCapabilityPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCapabilityPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCapabilityPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

