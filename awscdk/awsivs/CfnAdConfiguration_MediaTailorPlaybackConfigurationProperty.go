package awsivs


// Object specifying a configuration for integration with an AWS Elemental MediaTailor (EMT).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaTailorPlaybackConfigurationProperty := &MediaTailorPlaybackConfigurationProperty{
//   	PlaybackConfigurationArn: jsii.String("playbackConfigurationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-adconfiguration-mediatailorplaybackconfiguration.html
//
type CfnAdConfiguration_MediaTailorPlaybackConfigurationProperty struct {
	// ARN of the customer-created EMT PlaybackConfiguration resource in the same region and account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-adconfiguration-mediatailorplaybackconfiguration.html#cfn-ivs-adconfiguration-mediatailorplaybackconfiguration-playbackconfigurationarn
	//
	PlaybackConfigurationArn *string `field:"optional" json:"playbackConfigurationArn" yaml:"playbackConfigurationArn"`
}

