package previewawsgroundstationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgroundstation/previewawsgroundstationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a `DataflowEndpoint` group containing the specified list of Ground Station Agent based endpoints.
//
// The `name` field in each endpoint is used in your mission profile `DataflowEndpointConfig` to specify which endpoints to use during a contact.
//
// When a contact uses multiple `DataflowEndpointConfig` objects, each `Config` must match a `DataflowEndpoint` in the same group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataflowEndpointGroupV2PropsMixin := awscdkmixinspreview.Mixins.NewCfnDataflowEndpointGroupV2PropsMixin(&CfnDataflowEndpointGroupV2MixinProps{
//   	ContactPostPassDurationSeconds: jsii.Number(123),
//   	ContactPrePassDurationSeconds: jsii.Number(123),
//   	Endpoints: []interface{}{
//   		&CreateEndpointDetailsProperty{
//   			DownlinkAwsGroundStationAgentEndpoint: &DownlinkAwsGroundStationAgentEndpointProperty{
//   				DataflowDetails: &DownlinkDataflowDetailsProperty{
//   					AgentConnectionDetails: &DownlinkConnectionDetailsProperty{
//   						AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   							Mtu: jsii.Number(123),
//   							SocketAddress: &RangedSocketAddressProperty{
//   								Name: jsii.String("name"),
//   								PortRange: &IntegerRangeProperty{
//   									Maximum: jsii.Number(123),
//   									Minimum: jsii.Number(123),
//   								},
//   							},
//   						},
//   						EgressAddressAndPort: &ConnectionDetailsProperty{
//   							Mtu: jsii.Number(123),
//   							SocketAddress: &SocketAddressProperty{
//   								Name: jsii.String("name"),
//   								Port: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   			UplinkAwsGroundStationAgentEndpoint: &UplinkAwsGroundStationAgentEndpointProperty{
//   				DataflowDetails: &UplinkDataflowDetailsProperty{
//   					AgentConnectionDetails: &UplinkConnectionDetailsProperty{
//   						AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   							Mtu: jsii.Number(123),
//   							SocketAddress: &RangedSocketAddressProperty{
//   								Name: jsii.String("name"),
//   								PortRange: &IntegerRangeProperty{
//   									Maximum: jsii.Number(123),
//   									Minimum: jsii.Number(123),
//   								},
//   							},
//   						},
//   						IngressAddressAndPort: &ConnectionDetailsProperty{
//   							Mtu: jsii.Number(123),
//   							SocketAddress: &SocketAddressProperty{
//   								Name: jsii.String("name"),
//   								Port: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-dataflowendpointgroupv2.html
//
type CfnDataflowEndpointGroupV2PropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataflowEndpointGroupV2MixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataflowEndpointGroupV2PropsMixin
type jsiiProxy_CfnDataflowEndpointGroupV2PropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2PropsMixin) Props() *CfnDataflowEndpointGroupV2MixinProps {
	var returns *CfnDataflowEndpointGroupV2MixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2PropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GroundStation::DataflowEndpointGroupV2`.
func NewCfnDataflowEndpointGroupV2PropsMixin(props *CfnDataflowEndpointGroupV2MixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataflowEndpointGroupV2PropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataflowEndpointGroupV2PropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataflowEndpointGroupV2PropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GroundStation::DataflowEndpointGroupV2`.
func NewCfnDataflowEndpointGroupV2PropsMixin_Override(c CfnDataflowEndpointGroupV2PropsMixin, props *CfnDataflowEndpointGroupV2MixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataflowEndpointGroupV2PropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataflowEndpointGroupV2PropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataflowEndpointGroupV2PropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupV2PropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2PropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2PropsMixin) Supports(construct constructs.IConstruct) *bool {
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

