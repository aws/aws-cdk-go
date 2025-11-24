package mixinsawsmediaconnect


// Properties for CfnBridgePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBridgeMixinProps := &CfnBridgeMixinProps{
//   	EgressGatewayBridge: &EgressGatewayBridgeProperty{
//   		MaxBitrate: jsii.Number(123),
//   	},
//   	IngressGatewayBridge: &IngressGatewayBridgeProperty{
//   		MaxBitrate: jsii.Number(123),
//   		MaxOutputs: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
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
//   	PlacementArn: jsii.String("placementArn"),
//   	SourceFailoverConfig: &FailoverConfigProperty{
//   		FailoverMode: jsii.String("failoverMode"),
//   		SourcePriority: &SourcePriorityProperty{
//   			PrimarySource: jsii.String("primarySource"),
//   		},
//   		State: jsii.String("state"),
//   	},
//   	Sources: []interface{}{
//   		&BridgeSourceProperty{
//   			FlowSource: &BridgeFlowSourceProperty{
//   				FlowArn: jsii.String("flowArn"),
//   				FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   					VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   				},
//   				Name: jsii.String("name"),
//   			},
//   			NetworkSource: &BridgeNetworkSourceProperty{
//   				MulticastIp: jsii.String("multicastIp"),
//   				MulticastSourceSettings: &MulticastSourceSettingsProperty{
//   					MulticastSourceIp: jsii.String("multicastSourceIp"),
//   				},
//   				Name: jsii.String("name"),
//   				NetworkName: jsii.String("networkName"),
//   				Port: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html
//
type CfnBridgeMixinProps struct {
	// An egress bridge is a cloud-to-ground bridge.
	//
	// The content comes from an existing MediaConnect flow and is delivered to your premises.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-egressgatewaybridge
	//
	EgressGatewayBridge interface{} `field:"optional" json:"egressGatewayBridge" yaml:"egressGatewayBridge"`
	// An ingress bridge is a ground-to-cloud bridge.
	//
	// The content originates at your premises and is delivered to the cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-ingressgatewaybridge
	//
	IngressGatewayBridge interface{} `field:"optional" json:"ingressGatewayBridge" yaml:"ingressGatewayBridge"`
	// The name of the bridge.
	//
	// This name can not be modified after the bridge is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The outputs that you want to add to this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-outputs
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// The bridge placement Amazon Resource Number (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-placementarn
	//
	PlacementArn *string `field:"optional" json:"placementArn" yaml:"placementArn"`
	// The settings for source failover.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-sourcefailoverconfig
	//
	SourceFailoverConfig interface{} `field:"optional" json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
	// The sources that you want to add to this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html#cfn-mediaconnect-bridge-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

