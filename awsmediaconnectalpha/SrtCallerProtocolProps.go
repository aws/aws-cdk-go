package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for SRT Caller protocol configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//   var secret Secret
//
//   srtCallerProtocolProps := &SrtCallerProtocolProps{
//   	MinimumLatency: cdk.Duration_Minutes(jsii.Number(30)),
//   	SourceAddress: jsii.String("sourceAddress"),
//   	SourcePort: jsii.Number(123),
//
//   	// the properties below are optional
//   	DecryptionConfiguration: &RouterSrtEncryption{
//   		Secret: secret,
//
//   		// the properties below are optional
//   		Role: role,
//   	},
//   	StreamId: jsii.String("streamId"),
//   }
//
// Experimental.
type SrtCallerProtocolProps struct {
	// Minimum latency for SRT.
	// Experimental.
	MinimumLatency awscdk.Duration `field:"required" json:"minimumLatency" yaml:"minimumLatency"`
	// Source IP address to connect to.
	// Experimental.
	SourceAddress *string `field:"required" json:"sourceAddress" yaml:"sourceAddress"`
	// Source port to connect to.
	// Experimental.
	SourcePort *float64 `field:"required" json:"sourcePort" yaml:"sourcePort"`
	// Optional decryption configuration for encrypted SRT streams.
	// Default: - No decryption.
	//
	// Experimental.
	DecryptionConfiguration *RouterSrtEncryption `field:"optional" json:"decryptionConfiguration" yaml:"decryptionConfiguration"`
	// Optional stream ID for SRT connection.
	// Default: - No stream ID.
	//
	// Experimental.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

