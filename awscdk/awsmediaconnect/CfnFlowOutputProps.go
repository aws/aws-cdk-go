package awsmediaconnect


// Properties for defining a `CfnFlowOutput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowOutputProps := &CfnFlowOutputProps{
//   	FlowArn: jsii.String("flowArn"),
//   	Protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	CidrAllowList: []*string{
//   		jsii.String("cidrAllowList"),
//   	},
//   	Description: jsii.String("description"),
//   	Destination: jsii.String("destination"),
//   	Encryption: &EncryptionProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//
//   		// the properties below are optional
//   		Algorithm: jsii.String("algorithm"),
//   		KeyType: jsii.String("keyType"),
//   	},
//   	MaxLatency: jsii.Number(123),
//   	MediaStreamOutputConfigurations: []interface{}{
//   		&MediaStreamOutputConfigurationProperty{
//   			EncodingName: jsii.String("encodingName"),
//   			MediaStreamName: jsii.String("mediaStreamName"),
//
//   			// the properties below are optional
//   			DestinationConfigurations: []interface{}{
//   				&DestinationConfigurationProperty{
//   					DestinationIp: jsii.String("destinationIp"),
//   					DestinationPort: jsii.Number(123),
//   					Interface: &InterfaceProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			EncodingParameters: &EncodingParametersProperty{
//   				CompressionFactor: jsii.Number(123),
//
//   				// the properties below are optional
//   				EncoderProfile: jsii.String("encoderProfile"),
//   			},
//   		},
//   	},
//   	MinLatency: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	OutputStatus: jsii.String("outputStatus"),
//   	Port: jsii.Number(123),
//   	RemoteId: jsii.String("remoteId"),
//   	SmoothingLatency: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   	VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html
//
type CfnFlowOutputProps struct {
	// The Amazon Resource Name (ARN) of the flow this output is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-flowarn
	//
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The protocol to use for the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The range of IP addresses that are allowed to initiate output requests to this flow.
	//
	// Format the IP addresses as a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-cidrallowlist
	//
	CidrAllowList *[]*string `field:"optional" json:"cidrAllowList" yaml:"cidrAllowList"`
	// A description of the output.
	//
	// This description is not visible outside of the current AWS account even if the account grants entitlements to other accounts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IP address where you want to send the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The encryption credentials that you want to use for the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// The maximum latency in milliseconds.
	//
	// This parameter applies only to RIST-based, Zixi-based, and Fujitsu-based streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-maxlatency
	//
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The definition for each media stream that is associated with the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-mediastreamoutputconfigurations
	//
	MediaStreamOutputConfigurations interface{} `field:"optional" json:"mediaStreamOutputConfigurations" yaml:"mediaStreamOutputConfigurations"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents the minimal potential latency of that connection. The latency of the stream is set to the highest number between the sender’s minimum latency and the receiver’s minimum latency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-minlatency
	//
	MinLatency *float64 `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The name of the output.
	//
	// This value must be unique within the current flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An indication of whether the new output should be enabled or disabled as soon as it is created.
	//
	// If you don't specify the outputStatus field in your request, MediaConnect sets it to ENABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-outputstatus
	//
	OutputStatus *string `field:"optional" json:"outputStatus" yaml:"outputStatus"`
	// The port to use when MediaConnect distributes content to the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The identifier that is assigned to the Zixi receiver.
	//
	// This parameter applies only to outputs that use Zixi pull.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-remoteid
	//
	RemoteId *string `field:"optional" json:"remoteId" yaml:"remoteId"`
	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-smoothinglatency
	//
	SmoothingLatency *float64 `field:"optional" json:"smoothingLatency" yaml:"smoothingLatency"`
	// The stream ID that you want to use for this transport.
	//
	// This parameter applies only to Zixi and SRT caller-based streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The VPC interface that you want to send your output to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-vpcinterfaceattachment
	//
	VpcInterfaceAttachment interface{} `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

