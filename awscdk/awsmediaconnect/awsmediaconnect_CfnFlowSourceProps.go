package awsmediaconnect


// Properties for defining a `CfnFlowSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowSourceProps := &cfnFlowSourceProps{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	decryption: &encryptionProperty{
//   		algorithm: jsii.String("algorithm"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		constantInitializationVector: jsii.String("constantInitializationVector"),
//   		deviceId: jsii.String("deviceId"),
//   		keyType: jsii.String("keyType"),
//   		region: jsii.String("region"),
//   		resourceId: jsii.String("resourceId"),
//   		secretArn: jsii.String("secretArn"),
//   		url: jsii.String("url"),
//   	},
//   	entitlementArn: jsii.String("entitlementArn"),
//   	flowArn: jsii.String("flowArn"),
//   	ingestPort: jsii.Number(123),
//   	maxBitrate: jsii.Number(123),
//   	maxLatency: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   	streamId: jsii.String("streamId"),
//   	vpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	whitelistCidr: jsii.String("whitelistCidr"),
//   }
//
type CfnFlowSourceProps struct {
	// A description of the source.
	//
	// This description is not visible outside of the current AWS account.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the source.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of encryption that is used on the content ingested from the source.
	Decryption interface{} `field:"optional" json:"decryption" yaml:"decryption"`
	// The ARN of the entitlement that allows you to subscribe to the flow.
	//
	// The entitlement is set by the content originator, and the ARN is generated as part of the originator's flow.
	EntitlementArn *string `field:"optional" json:"entitlementArn" yaml:"entitlementArn"`
	// The Amazon Resource Name (ARN) of the flow.
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The port that the flow listens on for incoming content.
	//
	// If the protocol of the source is Zixi, the port must be set to 2088.
	IngestPort *float64 `field:"optional" json:"ingestPort" yaml:"ingestPort"`
	// The maximum bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum latency in milliseconds.
	//
	// This parameter applies only to RIST-based, Zixi-based, and Fujitsu-based streams.
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The protocol that the source uses to deliver the content to MediaConnect.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The stream ID that you want to use for the transport.
	//
	// This parameter applies only to Zixi-based streams.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The name of the VPC interface that you want to send your output to.
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
	// The range of IP addresses that are allowed to contribute content to your source.
	//
	// Format the IP addresses as a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr *string `field:"optional" json:"whitelistCidr" yaml:"whitelistCidr"`
}

