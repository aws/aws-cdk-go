package mixinsawsservicediscovery

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsservicediscovery/mixinsawsservicediscovery/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// A complex type that contains information about the specified service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var serviceAttributes interface{}
//
//   cfnServicePropsMixin := awscdkmixinspreview.Mixins.NewCfnServicePropsMixin(&CfnServiceMixinProps{
//   	Description: jsii.String("description"),
//   	DnsConfig: &DnsConfigProperty{
//   		DnsRecords: []interface{}{
//   			&DnsRecordProperty{
//   				Ttl: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		NamespaceId: jsii.String("namespaceId"),
//   		RoutingPolicy: jsii.String("routingPolicy"),
//   	},
//   	HealthCheckConfig: &HealthCheckConfigProperty{
//   		FailureThreshold: jsii.Number(123),
//   		ResourcePath: jsii.String("resourcePath"),
//   		Type: jsii.String("type"),
//   	},
//   	HealthCheckCustomConfig: &HealthCheckCustomConfigProperty{
//   		FailureThreshold: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	NamespaceId: jsii.String("namespaceId"),
//   	ServiceAttributes: serviceAttributes,
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html
//
type CfnServicePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnServiceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServicePropsMixin
type jsiiProxy_CfnServicePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnServicePropsMixin) Props() *CfnServiceMixinProps {
	var returns *CfnServiceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServicePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ServiceDiscovery::Service`.
func NewCfnServicePropsMixin(props *CfnServiceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnServicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_servicediscovery.mixins.CfnServicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ServiceDiscovery::Service`.
func NewCfnServicePropsMixin_Override(c CfnServicePropsMixin, props *CfnServiceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_servicediscovery.mixins.CfnServicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnServicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_servicediscovery.mixins.CfnServicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_servicediscovery.mixins.CfnServicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServicePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnServicePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

