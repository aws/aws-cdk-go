package awsmediaconnect


// Properties for defining a `CfnFlowOutput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   cfnFlowOutputProps := &CfnFlowOutputProps{
//   	FlowArn: jsii.String("flowArn"),
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
//   	NdiProgramName: jsii.String("ndiProgramName"),
//   	NdiSpeedHqQuality: jsii.Number(123),
//   	OutputStatus: jsii.String("outputStatus"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	RemoteId: jsii.String("remoteId"),
//   	RouterIntegrationState: jsii.String("routerIntegrationState"),
//   	RouterIntegrationTransitEncryption: &FlowTransitEncryptionProperty{
//   		EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		EncryptionKeyType: jsii.String("encryptionKeyType"),
//   	},
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
	// The range of IP addresses that should be allowed to initiate output requests to this flow.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-cidrallowlist
	//
	CidrAllowList *[]*string `field:"optional" json:"cidrAllowList" yaml:"cidrAllowList"`
	// A description of the output.
	//
	// This description appears only on the MediaConnect console and will not be seen by the end user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IP address where you want to send the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The type of key used for the encryption.
	//
	// If no `keyType` is provided, the service will use the default setting (static-key). Allowable encryption types: static-key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// The maximum latency in milliseconds.
	//
	// This parameter applies only to RIST-based and Zixi-based streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-maxlatency
	//
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The media streams that are associated with the output, and the parameters for those associations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-mediastreamoutputconfigurations
	//
	MediaStreamOutputConfigurations interface{} `field:"optional" json:"mediaStreamOutputConfigurations" yaml:"mediaStreamOutputConfigurations"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents the minimal potential latency of that connection. The latency of the stream is set to the highest number between the sender’s minimum latency and the receiver’s minimum latency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-minlatency
	//
	MinLatency *float64 `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The name of the bridge's output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A suffix for the names of the NDI sources that the flow creates.
	//
	// If a custom name isn't specified, MediaConnect uses the output name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-ndiprogramname
	//
	NdiProgramName *string `field:"optional" json:"ndiProgramName" yaml:"ndiProgramName"`
	// A quality setting for the NDI Speed HQ encoder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-ndispeedhqquality
	//
	NdiSpeedHqQuality *float64 `field:"optional" json:"ndiSpeedHqQuality" yaml:"ndiSpeedHqQuality"`
	// An indication of whether the output should transmit data or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-outputstatus
	//
	OutputStatus *string `field:"optional" json:"outputStatus" yaml:"outputStatus"`
	// The port to use when content is distributed to this output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use for the output.
	//
	// > AWS Elemental MediaConnect no longer supports the Fujitsu QoS protocol. This reference is maintained for legacy purposes only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The remote ID for the Zixi-pull stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-remoteid
	//
	RemoteId *string `field:"optional" json:"remoteId" yaml:"remoteId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-routerintegrationstate
	//
	RouterIntegrationState *string `field:"optional" json:"routerIntegrationState" yaml:"routerIntegrationState"`
	// Encryption information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-routerintegrationtransitencryption
	//
	RouterIntegrationTransitEncryption interface{} `field:"optional" json:"routerIntegrationTransitEncryption" yaml:"routerIntegrationTransitEncryption"`
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
	// The name of the VPC interface attachment to use for this output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html#cfn-mediaconnect-flowoutput-vpcinterfaceattachment
	//
	VpcInterfaceAttachment interface{} `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

