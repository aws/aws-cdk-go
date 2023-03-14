package awsmediatailor


// The configuration for HLS content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsConfigurationProperty := &hlsConfigurationProperty{
//   	manifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   }
//
type CfnPlaybackConfiguration_HlsConfigurationProperty struct {
	// The URL that is used to initiate a playback session for devices that support Apple HLS.
	//
	// The session uses server-side reporting.
	ManifestEndpointPrefix *string `field:"optional" json:"manifestEndpointPrefix" yaml:"manifestEndpointPrefix"`
}

