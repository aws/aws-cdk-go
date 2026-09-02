package awsmedialivealpha


// Options for an ancillary caption source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   ancillaryCaptionSourceOptions := &AncillaryCaptionSourceOptions{
//   	SourceChannelNumber: jsii.Number(123),
//   }
//
// Experimental.
type AncillaryCaptionSourceOptions struct {
	// The captions channel (1-4) to extract from the ancillary captions.
	//
	// Required when
	// converting to another caption format; ignored when passing through as embedded.
	// Default: - MediaLive ignores the channel (passthrough).
	//
	// Experimental.
	SourceChannelNumber *float64 `field:"optional" json:"sourceChannelNumber" yaml:"sourceChannelNumber"`
}

