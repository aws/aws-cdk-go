package previewawsgroundstationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgroundstation/previewawsgroundstationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Dataflow Endpoint Group request.
//
// Dataflow endpoint groups contain a list of endpoints. When the name of a dataflow endpoint group is specified in a mission profile, the Ground Station service will connect to the endpoints and flow data during a contact.
//
// For more information about dataflow endpoint groups, see [Dataflow Endpoint Groups](https://docs.aws.amazon.com/ground-station/latest/ug/dataflowendpointgroups.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataflowEndpointGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnDataflowEndpointGroupPropsMixin(&CfnDataflowEndpointGroupMixinProps{
//   	ContactPostPassDurationSeconds: jsii.Number(123),
//   	ContactPrePassDurationSeconds: jsii.Number(123),
//   	EndpointDetails: []interface{}{
//   		&EndpointDetailsProperty{
//   			AwsGroundStationAgentEndpoint: &AwsGroundStationAgentEndpointProperty{
//   				AgentStatus: jsii.String("agentStatus"),
//   				AuditResults: jsii.String("auditResults"),
//   				EgressAddress: &ConnectionDetailsProperty{
//   					Mtu: jsii.Number(123),
//   					SocketAddress: &SocketAddressProperty{
//   						Name: jsii.String("name"),
//   						Port: jsii.Number(123),
//   					},
//   				},
//   				IngressAddress: &RangedConnectionDetailsProperty{
//   					Mtu: jsii.Number(123),
//   					SocketAddress: &RangedSocketAddressProperty{
//   						Name: jsii.String("name"),
//   						PortRange: &IntegerRangeProperty{
//   							Maximum: jsii.Number(123),
//   							Minimum: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   			Endpoint: &DataflowEndpointProperty{
//   				Address: &SocketAddressProperty{
//   					Name: jsii.String("name"),
//   					Port: jsii.Number(123),
//   				},
//   				Mtu: jsii.Number(123),
//   				Name: jsii.String("name"),
//   			},
//   			SecurityDetails: &SecurityDetailsProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-dataflowendpointgroup.html
//
type CfnDataflowEndpointGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataflowEndpointGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataflowEndpointGroupPropsMixin
type jsiiProxy_CfnDataflowEndpointGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataflowEndpointGroupPropsMixin) Props() *CfnDataflowEndpointGroupMixinProps {
	var returns *CfnDataflowEndpointGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GroundStation::DataflowEndpointGroup`.
func NewCfnDataflowEndpointGroupPropsMixin(props *CfnDataflowEndpointGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataflowEndpointGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataflowEndpointGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataflowEndpointGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GroundStation::DataflowEndpointGroup`.
func NewCfnDataflowEndpointGroupPropsMixin_Override(c CfnDataflowEndpointGroupPropsMixin, props *CfnDataflowEndpointGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataflowEndpointGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataflowEndpointGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataflowEndpointGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnDataflowEndpointGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

