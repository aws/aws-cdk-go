package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for an SRT caller input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret Secret
//   var srtDecryptionAlgorithm SrtDecryptionAlgorithm
//
//   srtCallerSourceProps := &SrtCallerSourceProps{
//   	SrtListenerAddress: jsii.String("srtListenerAddress"),
//   	SrtListenerPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	Decryption: &SrtDecryptionProps{
//   		Algorithm: srtDecryptionAlgorithm,
//   		PassphraseSecret: secret,
//   	},
//   	MinimumLatency: cdk.Duration_Minutes(jsii.Number(30)),
//   	StreamId: jsii.String("streamId"),
//   }
//
// Experimental.
type SrtCallerSourceProps struct {
	// The address of the SRT listener to connect to.
	// Experimental.
	SrtListenerAddress *string `field:"required" json:"srtListenerAddress" yaml:"srtListenerAddress"`
	// The port of the SRT listener.
	// Experimental.
	SrtListenerPort *float64 `field:"required" json:"srtListenerPort" yaml:"srtListenerPort"`
	// Decryption settings for the SRT connection.
	// Default: - no decryption.
	//
	// Experimental.
	Decryption *SrtDecryptionProps `field:"optional" json:"decryption" yaml:"decryption"`
	// The minimum latency.
	// Default: - service default.
	//
	// Experimental.
	MinimumLatency awscdk.Duration `field:"optional" json:"minimumLatency" yaml:"minimumLatency"`
	// The stream ID for the SRT connection.
	// Default: - no stream ID.
	//
	// Experimental.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

