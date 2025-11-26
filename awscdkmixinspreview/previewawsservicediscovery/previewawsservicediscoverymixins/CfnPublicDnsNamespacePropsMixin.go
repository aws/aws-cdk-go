package previewawsservicediscoverymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsservicediscovery/previewawsservicediscoverymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a public namespace based on DNS, which is visible on the internet.
//
// The namespace defines your service naming scheme. For example, if you name your namespace `example.com` and name your service `backend` , the resulting DNS name for the service is `backend.example.com` . You can discover instances that were registered with a public DNS namespace by using either a `DiscoverInstances` request or using DNS. For the current quota on the number of namespaces that you can create using the same AWS account , see [AWS Cloud Map quotas](https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html) in the *AWS Cloud Map Developer Guide* .
//
// > The `CreatePublicDnsNamespace` API operation is not supported in the AWS GovCloud (US) Regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPublicDnsNamespacePropsMixin := awscdkmixinspreview.Mixins.NewCfnPublicDnsNamespacePropsMixin(&CfnPublicDnsNamespaceMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Properties: &PropertiesProperty{
//   		DnsProperties: &PublicDnsPropertiesMutableProperty{
//   			Soa: &SOAProperty{
//   				Ttl: jsii.Number(123),
//   			},
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html
//
type CfnPublicDnsNamespacePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPublicDnsNamespaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPublicDnsNamespacePropsMixin
type jsiiProxy_CfnPublicDnsNamespacePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPublicDnsNamespacePropsMixin) Props() *CfnPublicDnsNamespaceMixinProps {
	var returns *CfnPublicDnsNamespaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicDnsNamespacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ServiceDiscovery::PublicDnsNamespace`.
func NewCfnPublicDnsNamespacePropsMixin(props *CfnPublicDnsNamespaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPublicDnsNamespacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPublicDnsNamespacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPublicDnsNamespacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_servicediscovery.mixins.CfnPublicDnsNamespacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ServiceDiscovery::PublicDnsNamespace`.
func NewCfnPublicDnsNamespacePropsMixin_Override(c CfnPublicDnsNamespacePropsMixin, props *CfnPublicDnsNamespaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_servicediscovery.mixins.CfnPublicDnsNamespacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPublicDnsNamespacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPublicDnsNamespacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_servicediscovery.mixins.CfnPublicDnsNamespacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPublicDnsNamespacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_servicediscovery.mixins.CfnPublicDnsNamespacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPublicDnsNamespacePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPublicDnsNamespacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

