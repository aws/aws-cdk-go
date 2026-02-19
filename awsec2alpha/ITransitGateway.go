package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Transit Gateway.
// Experimental.
type ITransitGateway interface {
	awscdk.IResource
	IRouteTarget
	// The default route table associated with the Transit Gateway.
	//
	// This route table is created by the CDK and is used to manage the routes
	// for attachments that do not have an explicitly defined route table association.
	// Experimental.
	DefaultRouteTable() ITransitGatewayRouteTable
	// Indicates whether new attachments are automatically associated with the default route table.
	//
	// If set to `true`, any VPC or VPN attachment will be automatically associated with
	// the default route table unless otherwise specified.
	// Experimental.
	DefaultRouteTableAssociation() *bool
	// Indicates whether route propagation to the default route table is enabled.
	//
	// When set to `true`, routes from attachments will be automatically propagated
	// to the default route table unless propagation is explicitly disabled.
	// Experimental.
	DefaultRouteTablePropagation() *bool
	// Whether or not DNS support is enabled on the Transit Gateway.
	// Experimental.
	DnsSupport() *bool
	// Whether or not security group referencing support is enabled on the Transit Gateway.
	// Experimental.
	SecurityGroupReferencingSupport() *bool
	// The Amazon Resource Name (ARN) of the Transit Gateway.
	//
	// The ARN uniquely identifies the Transit Gateway across AWS and is commonly
	// used for permissions and resource tracking.
	// Experimental.
	TransitGatewayArn() *string
	// The unique identifier of the Transit Gateway.
	//
	// This ID is automatically assigned by AWS upon creation of the Transit Gateway
	// and is used to reference it in various configurations and operations.
	// Experimental.
	TransitGatewayId() *string
}

// The jsii proxy for ITransitGateway
type jsiiProxy_ITransitGateway struct {
	internal.Type__awscdkIResource
	jsiiProxy_IRouteTarget
}

func (i *jsiiProxy_ITransitGateway) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ITransitGateway) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITransitGateway) DefaultRouteTable() ITransitGatewayRouteTable {
	var returns ITransitGatewayRouteTable
	_jsii_.Get(
		j,
		"defaultRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) DefaultRouteTableAssociation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"defaultRouteTableAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) DefaultRouteTablePropagation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"defaultRouteTablePropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) DnsSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"dnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) SecurityGroupReferencingSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"securityGroupReferencingSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) TransitGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) RouterTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerTargetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) RouterType() awsec2.RouterType {
	var returns awsec2.RouterType
	_jsii_.Get(
		j,
		"routerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGateway) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

