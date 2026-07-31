package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// VPC Interface configuration.
//
// Example:
//   var stack Stack
//   var gateway Gateway
//   var flow Flow
//   var vpcInterface VpcInterfaceConfig
//   var productionNetwork GatewayNetwork
//
//
//   egressBridge := awsmediaconnectalpha.NewBridge(stack, jsii.String("MyEgressBridge"), &BridgeProps{
//   	BridgeName: jsii.String("my-egress-bridge"),
//   	Config: awsmediaconnectalpha.BridgeConfiguration_Egress(&EgressBridgeConfiguration{
//   		MaxBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   		FlowSources: []BridgeFlowInput{
//   			&BridgeFlowInput{
//   				Name: jsii.String("cloud-source"),
//   				Source: &BridgeFlowSource{
//   					Flow: flow,
//   					VpcInterface: vpcInterface,
//   				},
//   			},
//   		},
//   		NetworkOutputs: []BridgeNetworkOutput{
//   			&BridgeNetworkOutput{
//   				Name: jsii.String("on-prem-output"),
//   				Output: awsmediaconnectalpha.BridgeOutputConfiguration_Network(&BridgeNetworkOutputProps{
//   					IpAddress: jsii.String("192.168.1.200"),
//   					Port: jsii.Number(5001),
//   					Network: productionNetwork,
//   					Protocol: awsmediaconnectalpha.BridgeProtocol_RTP(),
//   					Ttl: jsii.Number(50),
//   				}),
//   			},
//   		},
//   	}),
//   	Gateway: gateway,
//   })
//
// Experimental.
type VpcInterfaceConfig struct {
	// Unique name for this VPC interface within the flow.
	//
	// Cannot be changed after creation.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// IAM role that MediaConnect assumes to create ENIs in your account.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// Security groups to apply to the ENI.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Subnet where the ENI is created.
	//
	// Must be in the same Availability Zone as the flow.
	// Experimental.
	Subnet awsec2.ISubnet `field:"required" json:"subnet" yaml:"subnet"`
	// IDs of the network interfaces.
	//
	// Set this when importing existing ENIs via
	// `VpcInterface.fromNetworkInterfaces()`; leave unset to have MediaConnect create them.
	// Default: - MediaConnect creates network interfaces automatically.
	//
	// Experimental.
	NetworkInterfaceIds *[]*string `field:"optional" json:"networkInterfaceIds" yaml:"networkInterfaceIds"`
	// The type of network interface.
	//
	// Use `EFA` for CDI workflows.
	// Default: - undefined; when omitted, MediaConnect applies NetworkInterface.ENA at deploy time
	//
	// Experimental.
	NetworkInterfaceType NetworkInterface `field:"optional" json:"networkInterfaceType" yaml:"networkInterfaceType"`
}

