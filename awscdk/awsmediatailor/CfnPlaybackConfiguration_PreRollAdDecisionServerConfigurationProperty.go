package awsmediatailor


// The configuration for the request to the pre-roll Ad Decision Server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   preRollAdDecisionServerConfigurationProperty := &PreRollAdDecisionServerConfigurationProperty{
//   	VastResponse: &PreRollVastResponseProperty{
//   		AdSequencingMode: jsii.String("adSequencingMode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-prerolladdecisionserverconfiguration.html
//
type CfnPlaybackConfiguration_PreRollAdDecisionServerConfigurationProperty struct {
	// The configuration for how MediaTailor processes the VAST response returned by the pre-roll Ad Decision Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-prerolladdecisionserverconfiguration.html#cfn-mediatailor-playbackconfiguration-prerolladdecisionserverconfiguration-vastresponse
	//
	VastResponse interface{} `field:"optional" json:"vastResponse" yaml:"vastResponse"`
}

