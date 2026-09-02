package awsmediatailor


// The configuration for how MediaTailor processes the VAST response returned by the pre-roll Ad Decision Server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   preRollVastResponseProperty := &PreRollVastResponseProperty{
//   	AdSequencingMode: jsii.String("adSequencingMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-prerollvastresponse.html
//
type CfnPlaybackConfiguration_PreRollVastResponseProperty struct {
	// Determines how MediaTailor sequences ads returned in the pre-roll VAST response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-prerollvastresponse.html#cfn-mediatailor-playbackconfiguration-prerollvastresponse-adsequencingmode
	//
	AdSequencingMode *string `field:"optional" json:"adSequencingMode" yaml:"adSequencingMode"`
}

