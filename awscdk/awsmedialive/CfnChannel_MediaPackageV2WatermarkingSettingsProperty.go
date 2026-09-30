package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaPackageV2WatermarkingSettingsProperty := &MediaPackageV2WatermarkingSettingsProperty{
//   	MediaPackageV2AbWatermarkerIrdetoSettings: &CmafIngestMediaPackageV2AbWatermarkerIrdetoSettingsProperty{
//   		AdditionalDestinationsAlternateDestinations: []interface{}{
//   			&OutputLocationRefProperty{
//   				DestinationRefId: jsii.String("destinationRefId"),
//   			},
//   		},
//   		AlternateDestination: &OutputLocationRefProperty{
//   			DestinationRefId: jsii.String("destinationRefId"),
//   		},
//   		CustomProfile: &AbWatermarkingCustomProfileProperty{
//   			EmbeddingFrequency: jsii.Number(123),
//   			SceneCut: jsii.Number(123),
//   			TargetPsnr: jsii.Number(123),
//   		},
//   		License: jsii.String("license"),
//   		OperatorId: jsii.Number(123),
//   		PolyPeriod: jsii.Number(123),
//   		Profile: jsii.String("profile"),
//   		WatermarkIdLength: jsii.String("watermarkIdLength"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2watermarkingsettings.html
//
type CfnChannel_MediaPackageV2WatermarkingSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackagev2watermarkingsettings.html#cfn-medialive-channel-mediapackagev2watermarkingsettings-mediapackagev2abwatermarkerirdetosettings
	//
	MediaPackageV2AbWatermarkerIrdetoSettings interface{} `field:"optional" json:"mediaPackageV2AbWatermarkerIrdetoSettings" yaml:"mediaPackageV2AbWatermarkerIrdetoSettings"`
}

