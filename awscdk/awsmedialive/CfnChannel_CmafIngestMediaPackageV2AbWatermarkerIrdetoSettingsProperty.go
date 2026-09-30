package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmafIngestMediaPackageV2AbWatermarkerIrdetoSettingsProperty := &CmafIngestMediaPackageV2AbWatermarkerIrdetoSettingsProperty{
//   	AdditionalDestinationsAlternateDestinations: []interface{}{
//   		&OutputLocationRefProperty{
//   			DestinationRefId: jsii.String("destinationRefId"),
//   		},
//   	},
//   	AlternateDestination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   	CustomProfile: &AbWatermarkingCustomProfileProperty{
//   		EmbeddingFrequency: jsii.Number(123),
//   		SceneCut: jsii.Number(123),
//   		TargetPsnr: jsii.Number(123),
//   	},
//   	License: jsii.String("license"),
//   	OperatorId: jsii.Number(123),
//   	PolyPeriod: jsii.Number(123),
//   	Profile: jsii.String("profile"),
//   	WatermarkIdLength: jsii.String("watermarkIdLength"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings.html
//
type CfnChannel_CmafIngestMediaPackageV2AbWatermarkerIrdetoSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings.html#cfn-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings-additionaldestinationsalternatedestinations
	//
	AdditionalDestinationsAlternateDestinations interface{} `field:"optional" json:"additionalDestinationsAlternateDestinations" yaml:"additionalDestinationsAlternateDestinations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings.html#cfn-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings-alternatedestination
	//
	AlternateDestination interface{} `field:"optional" json:"alternateDestination" yaml:"alternateDestination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings.html#cfn-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings-customprofile
	//
	CustomProfile interface{} `field:"optional" json:"customProfile" yaml:"customProfile"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings.html#cfn-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings-license
	//
	License *string `field:"optional" json:"license" yaml:"license"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings.html#cfn-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings-operatorid
	//
	OperatorId *float64 `field:"optional" json:"operatorId" yaml:"operatorId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings.html#cfn-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings-polyperiod
	//
	PolyPeriod *float64 `field:"optional" json:"polyPeriod" yaml:"polyPeriod"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings.html#cfn-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings-profile
	//
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings.html#cfn-medialive-channel-cmafingestmediapackagev2abwatermarkerirdetosettings-watermarkidlength
	//
	WatermarkIdLength *string `field:"optional" json:"watermarkIdLength" yaml:"watermarkIdLength"`
}

