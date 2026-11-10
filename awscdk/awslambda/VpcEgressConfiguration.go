package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
)

// Configuration for a connector that gives the compute egress into a VPC.
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
type VpcEgressConfiguration struct {
	// The Lambda compute resource type this connector serves.
	//
	// The compute type cannot be changed after the connector is created.
	ComputeType NetworkConnectorComputeType `field:"required" json:"computeType" yaml:"computeType"`
	// The security groups to attach to the network interfaces.
	//
	// Between 1 and 5 security groups, in the same VPC as the subnets. At least one is required,
	// so the interfaces never fall back to the VPC's default security group. The connector only
	// initiates connections, so inbound rules are irrelevant; the outbound rules decide what the
	// compute can reach.
	SecurityGroups *[]interfacesawsec2.ISecurityGroupRef `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The VPC to create network interfaces in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The IP protocol the network interfaces use.
	//
	// The protocol cannot be changed after the connector is created.
	// Default: NetworkConnectorProtocol.IPV4
	//
	NetworkProtocol NetworkConnectorProtocol `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// Which subnets of the VPC to create network interfaces in.
	//
	// The selection must yield between 1 and 16 subnets.
	// Default: - the VPC's default selection: its private subnets, if it has any.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

