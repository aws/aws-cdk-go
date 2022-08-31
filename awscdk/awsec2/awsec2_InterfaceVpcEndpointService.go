package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A custom-hosted service for an interface VPC endpoint.
//
// Example:
//   var vpc vpc
//
//
//   ec2.NewInterfaceVpcEndpoint(this, jsii.String("VPC Endpoint"), &interfaceVpcEndpointProps{
//   	vpc: vpc,
//   	service: ec2.NewInterfaceVpcEndpointService(jsii.String("com.amazonaws.vpce.us-east-1.vpce-svc-uuddlrlrbastrtsvc"), jsii.Number(443)),
//   	// Choose which availability zones to place the VPC endpoint in, based on
//   	// available AZs
//   	subnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("us-east-1a"),
//   			jsii.String("us-east-1c"),
//   		},
//   	},
//   })
//
// Experimental.
type InterfaceVpcEndpointService interface {
	IInterfaceVpcEndpointService
	// The name of the service.
	// Experimental.
	Name() *string
	// The port of the service.
	// Experimental.
	Port() *float64
	// Whether Private DNS is supported by default.
	// Experimental.
	PrivateDnsDefault() *bool
}

// The jsii proxy struct for InterfaceVpcEndpointService
type jsiiProxy_InterfaceVpcEndpointService struct {
	jsiiProxy_IInterfaceVpcEndpointService
}

func (j *jsiiProxy_InterfaceVpcEndpointService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpointService) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpointService) PrivateDnsDefault() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privateDnsDefault",
		&returns,
	)
	return returns
}


// Experimental.
func NewInterfaceVpcEndpointService(name *string, port *float64) InterfaceVpcEndpointService {
	_init_.Initialize()

	j := jsiiProxy_InterfaceVpcEndpointService{}

	_jsii_.Create(
		"monocdk.aws_ec2.InterfaceVpcEndpointService",
		[]interface{}{name, port},
		&j,
	)

	return &j
}

// Experimental.
func NewInterfaceVpcEndpointService_Override(i InterfaceVpcEndpointService, name *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.InterfaceVpcEndpointService",
		[]interface{}{name, port},
		i,
	)
}

