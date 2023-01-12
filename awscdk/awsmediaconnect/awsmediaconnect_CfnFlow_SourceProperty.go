package awsmediaconnect


// The details of the sources of the flow.
//
// If you are creating a flow with a VPC source, you must first create the flow with a temporary standard source by doing the following:
//
// - Use CloudFormation to create a flow with a standard source that uses the flow’s public IP address.
// - Use CloudFormation to create the VPC interface to add to this flow. This can also be done as part of the previous step.
// - After CloudFormation has created the flow and the VPC interface, update the source to point to the VPC interface that you created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &sourceProperty{
//   	decryption: &encryptionProperty{
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		algorithm: jsii.String("algorithm"),
//   		constantInitializationVector: jsii.String("constantInitializationVector"),
//   		deviceId: jsii.String("deviceId"),
//   		keyType: jsii.String("keyType"),
//   		region: jsii.String("region"),
//   		resourceId: jsii.String("resourceId"),
//   		secretArn: jsii.String("secretArn"),
//   		url: jsii.String("url"),
//   	},
//   	description: jsii.String("description"),
//   	entitlementArn: jsii.String("entitlementArn"),
//   	ingestIp: jsii.String("ingestIp"),
//   	ingestPort: jsii.Number(123),
//   	maxBitrate: jsii.Number(123),
//   	maxLatency: jsii.Number(123),
//   	minLatency: jsii.Number(123),
//   	name: jsii.String("name"),
//   	protocol: jsii.String("protocol"),
//   	senderControlPort: jsii.Number(123),
//   	senderIpAddress: jsii.String("senderIpAddress"),
//   	sourceArn: jsii.String("sourceArn"),
//   	sourceIngestPort: jsii.String("sourceIngestPort"),
//   	sourceListenerAddress: jsii.String("sourceListenerAddress"),
//   	sourceListenerPort: jsii.Number(123),
//   	streamId: jsii.String("streamId"),
//   	vpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	whitelistCidr: jsii.String("whitelistCidr"),
//   }
//
type CfnFlow_SourceProperty struct {
	// The type of encryption that is used on the content ingested from the source.
	Decryption interface{} `field:"optional" json:"decryption" yaml:"decryption"`
	// A description of the source.
	//
	// This description is not visible outside of the current AWS account.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account.
	//
	// The entitlement is set by the content originator and the ARN is generated as part of the originator’s flow.
	EntitlementArn *string `field:"optional" json:"entitlementArn" yaml:"entitlementArn"`
	// The IP address that the flow listens on for incoming content.
	IngestIp *string `field:"optional" json:"ingestIp" yaml:"ingestIp"`
	// The port that the flow listens on for incoming content.
	//
	// If the protocol of the source is Zixi, the port must be set to 2088.
	IngestPort *float64 `field:"optional" json:"ingestPort" yaml:"ingestPort"`
	// The maximum bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum latency in milliseconds for a RIST or Zixi-based source.
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents the minimal potential latency of that connection. The latency of the stream is set to the highest number between the sender’s minimum latency and the receiver’s minimum latency.
	MinLatency *float64 `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The name of the source.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol that is used by the source.
	//
	// For a full list of available protocols, see: [Source protocols](https://docs.aws.amazon.com/mediaconnect/latest/api/v1-flows-flowarn-source.html#v1-flows-flowarn-source-prop-setsourcerequest-protocol) in the *AWS Elemental MediaConnect API Reference* .
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// `CfnFlow.SourceProperty.SenderControlPort`.
	SenderControlPort *float64 `field:"optional" json:"senderControlPort" yaml:"senderControlPort"`
	// `CfnFlow.SourceProperty.SenderIpAddress`.
	SenderIpAddress *string `field:"optional" json:"senderIpAddress" yaml:"senderIpAddress"`
	// The ARN of the source.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// The port that the flow will be listening on for incoming content.
	SourceIngestPort *string `field:"optional" json:"sourceIngestPort" yaml:"sourceIngestPort"`
	// `CfnFlow.SourceProperty.SourceListenerAddress`.
	SourceListenerAddress *string `field:"optional" json:"sourceListenerAddress" yaml:"sourceListenerAddress"`
	// `CfnFlow.SourceProperty.SourceListenerPort`.
	SourceListenerPort *float64 `field:"optional" json:"sourceListenerPort" yaml:"sourceListenerPort"`
	// The stream ID that you want to use for the transport.
	//
	// This parameter applies only to Zixi-based streams.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The name of the VPC interface that the source content comes from.
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
	// The range of IP addresses that are allowed to contribute content to your source.
	//
	// Format the IP addresses as a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr *string `field:"optional" json:"whitelistCidr" yaml:"whitelistCidr"`
}

