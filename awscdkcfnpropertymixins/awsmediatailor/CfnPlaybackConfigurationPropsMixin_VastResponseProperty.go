package awsmediatailor


// The configuration for how MediaTailor processes the VAST response returned by the Ad Decision Server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   vastResponseProperty := &VastResponseProperty{
//   	AdSequencingMode: jsii.String("adSequencingMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-vastresponse.html
//
type CfnPlaybackConfigurationPropsMixin_VastResponseProperty struct {
	// Determines how MediaTailor sequences ads returned in the VAST response from the Ad Decision Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-vastresponse.html#cfn-mediatailor-playbackconfiguration-vastresponse-adsequencingmode
	//
	AdSequencingMode *string `field:"optional" json:"adSequencingMode" yaml:"adSequencingMode"`
}

