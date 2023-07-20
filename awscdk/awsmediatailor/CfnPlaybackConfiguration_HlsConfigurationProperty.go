package awsmediatailor


// The configuration for HLS content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsConfigurationProperty := &HlsConfigurationProperty{
//   	ManifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-hlsconfiguration.html
//
type CfnPlaybackConfiguration_HlsConfigurationProperty struct {
	// The URL that is used to initiate a playback session for devices that support Apple HLS.
	//
	// The session uses server-side reporting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-hlsconfiguration.html#cfn-mediatailor-playbackconfiguration-hlsconfiguration-manifestendpointprefix
	//
	ManifestEndpointPrefix *string `field:"optional" json:"manifestEndpointPrefix" yaml:"manifestEndpointPrefix"`
}

