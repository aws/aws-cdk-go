package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cmafIngestWatermarkingSettingsProperty := &CmafIngestWatermarkingSettingsProperty{
//   	CmafIngestAbWatermarkerIrdetoSettings: &CmafIngestMediaPackageV2AbWatermarkerIrdetoSettingsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestwatermarkingsettings.html
//
type CfnChannelPropsMixin_CmafIngestWatermarkingSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestwatermarkingsettings.html#cfn-medialive-channel-cmafingestwatermarkingsettings-cmafingestabwatermarkerirdetosettings
	//
	CmafIngestAbWatermarkerIrdetoSettings interface{} `field:"optional" json:"cmafIngestAbWatermarkerIrdetoSettings" yaml:"cmafIngestAbWatermarkerIrdetoSettings"`
}

