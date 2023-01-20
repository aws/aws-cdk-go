package awsmediaconnect


// Properties for defining a `CfnFlowOutput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowOutputProps := &cfnFlowOutputProps{
//   	flowArn: jsii.String("flowArn"),
//   	protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	cidrAllowList: []*string{
//   		jsii.String("cidrAllowList"),
//   	},
//   	description: jsii.String("description"),
//   	destination: jsii.String("destination"),
//   	encryption: &encryptionProperty{
//   		roleArn: jsii.String("roleArn"),
//   		secretArn: jsii.String("secretArn"),
//
//   		// the properties below are optional
//   		algorithm: jsii.String("algorithm"),
//   		keyType: jsii.String("keyType"),
//   	},
//   	maxLatency: jsii.Number(123),
//   	minLatency: jsii.Number(123),
//   	name: jsii.String("name"),
//   	port: jsii.Number(123),
//   	remoteId: jsii.String("remoteId"),
//   	smoothingLatency: jsii.Number(123),
//   	streamId: jsii.String("streamId"),
//   	vpcInterfaceAttachment: &vpcInterfaceAttachmentProperty{
//   		vpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	},
//   }
//
type CfnFlowOutputProps struct {
	// The Amazon Resource Name (ARN) of the flow.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The protocol to use for the output.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The range of IP addresses that are allowed to initiate output requests to this flow.
	//
	// Format the IP addresses as a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	CidrAllowList *[]*string `field:"optional" json:"cidrAllowList" yaml:"cidrAllowList"`
	// A description of the output.
	//
	// This description is not visible outside of the current AWS account even if the account grants entitlements to other accounts.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IP address where you want to send the output.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The encryption credentials that you want to use for the output.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// The maximum latency in milliseconds for Zixi-based streams.
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents the minimal potential latency of that connection. The latency of the stream is set to the highest number between the sender’s minimum latency and the receiver’s minimum latency.
	MinLatency *float64 `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The name of the VPC interface.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port to use when MediaConnect distributes content to the output.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The identifier that is assigned to the Zixi receiver.
	//
	// This parameter applies only to outputs that use Zixi pull.
	RemoteId *string `field:"optional" json:"remoteId" yaml:"remoteId"`
	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	SmoothingLatency *float64 `field:"optional" json:"smoothingLatency" yaml:"smoothingLatency"`
	// The stream ID that you want to use for the transport.
	//
	// This parameter applies only to Zixi-based streams.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The VPC interface that you want to send your output to.
	VpcInterfaceAttachment interface{} `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

