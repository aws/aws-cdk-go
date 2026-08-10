package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration options for RTP-FEC outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   rtpFecOutputConfig := &RtpFecOutputConfig{
//   	Destination: jsii.String("destination"),
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	SmoothingLatency: cdk.Duration_Minutes(jsii.Number(30)),
//   	VpcInterfaceAttachmentName: jsii.String("vpcInterfaceAttachmentName"),
//   }
//
// Experimental.
type RtpFecOutputConfig struct {
	// The IP address where you want to send the output.
	// Experimental.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The port to use when content is distributed to this output.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The smoothing latency for RIST, RTP, and RTP-FEC streams.
	// Default: - no smoothing latency.
	//
	// Experimental.
	SmoothingLatency awscdk.Duration `field:"optional" json:"smoothingLatency" yaml:"smoothingLatency"`
	// The name of the VPC interface attachment to use for this output.
	// Default: - no VPC interface attachment.
	//
	// Experimental.
	VpcInterfaceAttachmentName *string `field:"optional" json:"vpcInterfaceAttachmentName" yaml:"vpcInterfaceAttachmentName"`
}

