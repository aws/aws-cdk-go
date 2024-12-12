package awsmediaconnect


// Properties for defining a `CfnFlowSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowSourceProps := &CfnFlowSourceProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Decryption: &EncryptionProperty{
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		Algorithm: jsii.String("algorithm"),
//   		ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   		DeviceId: jsii.String("deviceId"),
//   		KeyType: jsii.String("keyType"),
//   		Region: jsii.String("region"),
//   		ResourceId: jsii.String("resourceId"),
//   		SecretArn: jsii.String("secretArn"),
//   		Url: jsii.String("url"),
//   	},
//   	EntitlementArn: jsii.String("entitlementArn"),
//   	FlowArn: jsii.String("flowArn"),
//   	GatewayBridgeSource: &GatewayBridgeSourceProperty{
//   		BridgeArn: jsii.String("bridgeArn"),
//
//   		// the properties below are optional
//   		VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   			VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		},
//   	},
//   	IngestPort: jsii.Number(123),
//   	MaxBitrate: jsii.Number(123),
//   	MaxLatency: jsii.Number(123),
//   	MinLatency: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	SenderControlPort: jsii.Number(123),
//   	SenderIpAddress: jsii.String("senderIpAddress"),
//   	SourceListenerAddress: jsii.String("sourceListenerAddress"),
//   	SourceListenerPort: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   	VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	WhitelistCidr: jsii.String("whitelistCidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html
//
type CfnFlowSourceProps struct {
	// A description of the source.
	//
	// This description is not visible outside of the current AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of encryption that is used on the content ingested from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-decryption
	//
	Decryption interface{} `field:"optional" json:"decryption" yaml:"decryption"`
	// The ARN of the entitlement that allows you to subscribe to the flow.
	//
	// The entitlement is set by the content originator, and the ARN is generated as part of the originator's flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-entitlementarn
	//
	EntitlementArn *string `field:"optional" json:"entitlementArn" yaml:"entitlementArn"`
	// The Amazon Resource Name (ARN) of the flow this source is connected to.
	//
	// The flow must have Failover enabled to add an additional source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-flowarn
	//
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The source configuration for cloud flows receiving a stream from a bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-gatewaybridgesource
	//
	GatewayBridgeSource interface{} `field:"optional" json:"gatewayBridgeSource" yaml:"gatewayBridgeSource"`
	// The port that the flow listens on for incoming content.
	//
	// If the protocol of the source is Zixi, the port must be set to 2088.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-ingestport
	//
	IngestPort *float64 `field:"optional" json:"ingestPort" yaml:"ingestPort"`
	// The maximum bitrate for RIST, RTP, and RTP-FEC streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-maxbitrate
	//
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum latency in milliseconds.
	//
	// This parameter applies only to RIST-based and Zixi-based streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-maxlatency
	//
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents the minimal potential latency of that connection. The latency of the stream is set to the highest number between the sender’s minimum latency and the receiver’s minimum latency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-minlatency
	//
	MinLatency *float64 `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The protocol that the source uses to deliver the content to MediaConnect.
	//
	// Adding additional sources to an existing flow requires Failover to be enabled. When you enable Failover, the additional source must use the same protocol as the existing source. Only the following protocols support failover: Zixi-push, RTP-FEC, RTP, RIST and SRT protocols.
	//
	// If you use failover with SRT caller or listener, the `FailoverMode` property must be set to `FAILOVER` . The `FailoverMode` property is found in the `FailoverConfig` resource of the same flow ARN you used for the source's `FlowArn` property. SRT caller/listener does not support merge mode failover.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The port that the flow uses to send outbound requests to initiate connection with the sender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-sendercontrolport
	//
	SenderControlPort *float64 `field:"optional" json:"senderControlPort" yaml:"senderControlPort"`
	// The IP address that the flow communicates with to initiate connection with the sender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-senderipaddress
	//
	SenderIpAddress *string `field:"optional" json:"senderIpAddress" yaml:"senderIpAddress"`
	// Source IP or domain name for SRT-caller protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-sourcelisteneraddress
	//
	SourceListenerAddress *string `field:"optional" json:"sourceListenerAddress" yaml:"sourceListenerAddress"`
	// Source port for SRT-caller protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-sourcelistenerport
	//
	SourceListenerPort *float64 `field:"optional" json:"sourceListenerPort" yaml:"sourceListenerPort"`
	// The stream ID that you want to use for this transport.
	//
	// This parameter applies only to Zixi and SRT caller-based streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The name of the VPC interface that you want to send your output to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-vpcinterfacename
	//
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
	// The range of IP addresses that are allowed to contribute content to your source.
	//
	// Format the IP addresses as a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html#cfn-mediaconnect-flowsource-whitelistcidr
	//
	WhitelistCidr *string `field:"optional" json:"whitelistCidr" yaml:"whitelistCidr"`
}

