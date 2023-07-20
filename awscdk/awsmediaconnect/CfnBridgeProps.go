package awsmediaconnect


// Properties for defining a `CfnBridge`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBridgeProps := &CfnBridgeProps{
//   	Name: jsii.String("name"),
//   	PlacementArn: jsii.String("placementArn"),
//   	Sources: []interface{}{
//   		&BridgeSourceProperty{
//   			FlowSource: &BridgeFlowSourceProperty{
//   				FlowArn: jsii.String("flowArn"),
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   					VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   				},
//   			},
//   			NetworkSource: &BridgeNetworkSourceProperty{
//   				MulticastIp: jsii.String("multicastIp"),
//   				Name: jsii.String("name"),
//   				NetworkName: jsii.String("networkName"),
//   				Port: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	EgressGatewayBridge: &EgressGatewayBridgeProperty{
//   		MaxBitrate: jsii.Number(123),
//   	},
//   	IngressGatewayBridge: &IngressGatewayBridgeProperty{
//   		MaxBitrate: jsii.Number(123),
//   		MaxOutputs: jsii.Number(123),
//   	},
//   	Outputs: []interface{}{
//   		&BridgeOutputProperty{
//   			NetworkOutput: &BridgeNetworkOutputProperty{
//   				IpAddress: jsii.String("ipAddress"),
//   				Name: jsii.String("name"),
//   				NetworkName: jsii.String("networkName"),
//   				Port: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   				Ttl: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SourceFailoverConfig: &FailoverConfigProperty{
//   		FailoverMode: jsii.String("failoverMode"),
//
//   		// the properties below are optional
//   		SourcePriority: &SourcePriorityProperty{
//   			PrimarySource: jsii.String("primarySource"),
//   		},
//   		State: jsii.String("state"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html
//
type CfnBridgeProps struct {
	// The network output name.
	//
	// This name is used to reference the output and must be unique among outputs in this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The bridge placement Amazon Resource Number (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-placementarn
	//
	PlacementArn *string `field:"required" json:"placementArn" yaml:"placementArn"`
	// The sources that you want to add to this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-sources
	//
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// Create a bridge with the egress bridge type.
	//
	// An egress bridge is a cloud-to-ground bridge. The content comes from an existing MediaConnect flow and is delivered to your premises.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-egressgatewaybridge
	//
	EgressGatewayBridge interface{} `field:"optional" json:"egressGatewayBridge" yaml:"egressGatewayBridge"`
	// Create a bridge with the ingress bridge type.
	//
	// An ingress bridge is a ground-to-cloud bridge. The content originates at your premises and is delivered to the cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-ingressgatewaybridge
	//
	IngressGatewayBridge interface{} `field:"optional" json:"ingressGatewayBridge" yaml:"ingressGatewayBridge"`
	// The outputs that you want to add to this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-outputs
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// The settings for source failover.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-sourcefailoverconfig
	//
	SourceFailoverConfig interface{} `field:"optional" json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
}

