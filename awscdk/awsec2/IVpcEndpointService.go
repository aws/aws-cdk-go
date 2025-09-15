package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A VPC endpoint service.
type IVpcEndpointService interface {
	awscdk.IResource
	IVPCEndpointServiceRef
	// The id of the VPC Endpoint Service that clients use to connect to, like vpce-svc-xxxxxxxxxxxxxxxx.
	VpcEndpointServiceId() *string
	// The service name of the VPC Endpoint Service that clients use to connect to, like com.amazonaws.vpce.<region>.vpce-svc-xxxxxxxxxxxxxxxx.
	VpcEndpointServiceName() *string
}

// The jsii proxy for IVpcEndpointService
type jsiiProxy_IVpcEndpointService struct {
	internal.Type__awscdkIResource
	jsiiProxy_IVPCEndpointServiceRef
}

func (i *jsiiProxy_IVpcEndpointService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IVpcEndpointService) VpcEndpointServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcEndpointService) VpcEndpointServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcEndpointService) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcEndpointService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcEndpointService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcEndpointService) VpcEndpointServiceRef() *VPCEndpointServiceReference {
	var returns *VPCEndpointServiceReference
	_jsii_.Get(
		j,
		"vpcEndpointServiceRef",
		&returns,
	)
	return returns
}

