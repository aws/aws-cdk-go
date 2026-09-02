package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for HLS Basic PUT CDN settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsBasicPutCdnProps := &HlsBasicPutCdnProps{
//   	ConnectionRetryInterval: jsii.Number(123),
//   	FilecacheDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   	NumRetries: jsii.Number(123),
//   	RestartDelay: jsii.Number(123),
//   }
//
// Experimental.
type HlsBasicPutCdnProps struct {
	// The number of seconds to wait before retrying a connection to the CDN.
	// Default: 1.
	//
	// Experimental.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The size of the file cache for streaming outputs.
	// Default: Duration.seconds(300)
	//
	// Experimental.
	FilecacheDuration awscdk.Duration `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// The number of retry attempts.
	// Default: 10.
	//
	// Experimental.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// The number of seconds to wait before restarting after a failure.
	// Default: 1.
	//
	// Experimental.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

