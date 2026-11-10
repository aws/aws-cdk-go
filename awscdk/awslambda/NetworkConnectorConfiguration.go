package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The network configuration of a connector.
//
// Exactly one kind of configuration applies to a connector; pick it with one of the static
// factories.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc Vpc
//   var code WebFunctionCode
//
//
//   securityGroup := ec2.NewSecurityGroup(this, jsii.String("EgressSecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   })
//
//   connector := lambda.NewNetworkConnector(this, jsii.String("Egress"), &NetworkConnectorProps{
//   	Configuration: lambda.NetworkConnectorConfiguration_VpcEgress(&VpcEgressConfiguration{
//   		ComputeType: lambda.NetworkConnectorComputeType_WEB_FUNCTION,
//   		Vpc: *Vpc,
//   		VpcSubnets: &SubnetSelection{
//   			SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   		},
//   		SecurityGroups: []ISecurityGroupRef{
//   			securityGroup,
//   		},
//   	}),
//   })
//
//   lambda.NewWebFunction(this, jsii.String("MyWebFunction"), &WebFunctionProps{
//   	Code: Code,
//   	Runtime: lambda.WebFunctionRuntime_NODEJS_24_X(),
//   	DefaultEndpointAuth: lambda.WebFunctionAuthType_IAM_AUTH,
//   	EgressNetworkConnector: connector,
//   })
//
type NetworkConnectorConfiguration interface {
}

// The jsii proxy struct for NetworkConnectorConfiguration
type jsiiProxy_NetworkConnectorConfiguration struct {
	_ byte // padding
}

func NewNetworkConnectorConfiguration_Override(n NetworkConnectorConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.NetworkConnectorConfiguration",
		nil, // no parameters
		n,
	)
}

// Egress from the compute into a VPC, through network interfaces the connector manages.
func NetworkConnectorConfiguration_VpcEgress(config *VpcEgressConfiguration) NetworkConnectorConfiguration {
	_init_.Initialize()

	if err := validateNetworkConnectorConfiguration_VpcEgressParameters(config); err != nil {
		panic(err)
	}
	var returns NetworkConnectorConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.NetworkConnectorConfiguration",
		"vpcEgress",
		[]interface{}{config},
		&returns,
	)

	return returns
}

