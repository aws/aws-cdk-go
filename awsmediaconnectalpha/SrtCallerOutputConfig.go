package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration options for SRT Caller outputs.
//
// Example:
//   var stack Stack
//   var flow Flow
//   var role IRole
//   var secret ISecret
//
//
//   awsmediaconnectalpha.NewFlowOutput(stack, jsii.String("SrtOutput"), &FlowOutputProps{
//   	Flow: Flow,
//   	Output: awsmediaconnectalpha.OutputConfiguration_SrtCaller(&SrtCallerOutputConfig{
//   		Destination: jsii.String("203.0.113.100"),
//   		Port: jsii.Number(7000),
//   		Encryption: &SrtPasswordEncryption{
//   			Role: *Role,
//   			Secret: *Secret,
//   		},
//   	}),
//   })
//
// Experimental.
type SrtCallerOutputConfig struct {
	// The IP address where you want to send the output.
	// Experimental.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The port to use when content is distributed to this output.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// SRT password encryption for this output.
	// Default: - no encryption.
	//
	// Experimental.
	Encryption *SrtPasswordEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents
	// the minimal potential latency of that connection. The latency of the stream is set to the highest number between
	// the sender’s minimum latency and the receiver’s minimum latency.
	// Default: - no minimum latency.
	//
	// Experimental.
	MinLatency awscdk.Duration `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The stream ID that you want to use for this transport.
	// Default: - no stream ID.
	//
	// Experimental.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The name of the VPC interface attachment to use for this output.
	// Default: - no VPC interface attachment.
	//
	// Experimental.
	VpcInterfaceAttachmentName *string `field:"optional" json:"vpcInterfaceAttachmentName" yaml:"vpcInterfaceAttachmentName"`
}

