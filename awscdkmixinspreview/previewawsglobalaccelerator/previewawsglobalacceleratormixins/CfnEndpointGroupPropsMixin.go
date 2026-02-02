package previewawsglobalacceleratormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglobalaccelerator/previewawsglobalacceleratormixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::GlobalAccelerator::EndpointGroup` resource is a Global Accelerator resource type that contains information about how you create an endpoint group for the specified listener.
//
// An endpoint group is a collection of endpoints in one AWS Region .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEndpointGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnEndpointGroupPropsMixin(&CfnEndpointGroupMixinProps{
//   	EndpointConfigurations: []interface{}{
//   		&EndpointConfigurationProperty{
//   			AttachmentArn: jsii.String("attachmentArn"),
//   			ClientIpPreservationEnabled: jsii.Boolean(false),
//   			EndpointId: jsii.String("endpointId"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	EndpointGroupRegion: jsii.String("endpointGroupRegion"),
//   	HealthCheckIntervalSeconds: jsii.Number(123),
//   	HealthCheckPath: jsii.String("healthCheckPath"),
//   	HealthCheckPort: jsii.Number(123),
//   	HealthCheckProtocol: jsii.String("healthCheckProtocol"),
//   	ListenerArn: jsii.String("listenerArn"),
//   	PortOverrides: []interface{}{
//   		&PortOverrideProperty{
//   			EndpointPort: jsii.Number(123),
//   			ListenerPort: jsii.Number(123),
//   		},
//   	},
//   	ThresholdCount: jsii.Number(123),
//   	TrafficDialPercentage: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html
//
type CfnEndpointGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEndpointGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEndpointGroupPropsMixin
type jsiiProxy_CfnEndpointGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEndpointGroupPropsMixin) Props() *CfnEndpointGroupMixinProps {
	var returns *CfnEndpointGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GlobalAccelerator::EndpointGroup`.
func NewCfnEndpointGroupPropsMixin(props *CfnEndpointGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEndpointGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEndpointGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEndpointGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_globalaccelerator.mixins.CfnEndpointGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GlobalAccelerator::EndpointGroup`.
func NewCfnEndpointGroupPropsMixin_Override(c CfnEndpointGroupPropsMixin, props *CfnEndpointGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_globalaccelerator.mixins.CfnEndpointGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEndpointGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpointGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_globalaccelerator.mixins.CfnEndpointGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_globalaccelerator.mixins.CfnEndpointGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpointGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEndpointGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

