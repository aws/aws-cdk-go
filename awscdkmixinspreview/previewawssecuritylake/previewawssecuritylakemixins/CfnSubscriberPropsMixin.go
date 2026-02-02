package previewawssecuritylakemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssecuritylake/previewawssecuritylakemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a subscriber for accounts that are already enabled in Amazon Security Lake.
//
// You can create a subscriber with access to data in the current AWS Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSubscriberPropsMixin := awscdkmixinspreview.Mixins.NewCfnSubscriberPropsMixin(&CfnSubscriberMixinProps{
//   	AccessTypes: []*string{
//   		jsii.String("accessTypes"),
//   	},
//   	DataLakeArn: jsii.String("dataLakeArn"),
//   	Sources: []interface{}{
//   		&SourceProperty{
//   			AwsLogSource: &AwsLogSourceProperty{
//   				SourceName: jsii.String("sourceName"),
//   				SourceVersion: jsii.String("sourceVersion"),
//   			},
//   			CustomLogSource: &CustomLogSourceProperty{
//   				SourceName: jsii.String("sourceName"),
//   				SourceVersion: jsii.String("sourceVersion"),
//   			},
//   		},
//   	},
//   	SubscriberDescription: jsii.String("subscriberDescription"),
//   	SubscriberIdentity: &SubscriberIdentityProperty{
//   		ExternalId: jsii.String("externalId"),
//   		Principal: jsii.String("principal"),
//   	},
//   	SubscriberName: jsii.String("subscriberName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html
//
type CfnSubscriberPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSubscriberMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSubscriberPropsMixin
type jsiiProxy_CfnSubscriberPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSubscriberPropsMixin) Props() *CfnSubscriberMixinProps {
	var returns *CfnSubscriberMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriberPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityLake::Subscriber`.
func NewCfnSubscriberPropsMixin(props *CfnSubscriberMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSubscriberPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSubscriberPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSubscriberPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityLake::Subscriber`.
func NewCfnSubscriberPropsMixin_Override(c CfnSubscriberPropsMixin, props *CfnSubscriberMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSubscriberPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSubscriberPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubscriberPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubscriberPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSubscriberPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

