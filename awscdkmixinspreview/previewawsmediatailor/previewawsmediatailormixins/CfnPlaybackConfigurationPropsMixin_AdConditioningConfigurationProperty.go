package previewawsmediatailormixins


// The setting that indicates what conditioning MediaTailor will perform on ads that the ad decision server (ADS) returns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   adConditioningConfigurationProperty := &AdConditioningConfigurationProperty{
//   	StreamingMediaFileConditioning: jsii.String("streamingMediaFileConditioning"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adconditioningconfiguration.html
//
type CfnPlaybackConfigurationPropsMixin_AdConditioningConfigurationProperty struct {
	// For ads that have media files with streaming delivery and supported file extensions, indicates what transcoding action MediaTailor takes when it first receives these ads from the ADS.
	//
	// `TRANSCODE` indicates that MediaTailor must transcode the ads. `NONE` indicates that you have already transcoded the ads outside of MediaTailor and don't need them transcoded as part of the ad insertion workflow. For more information about ad conditioning see [Using preconditioned ads](https://docs.aws.amazon.com/mediatailor/latest/ug/precondition-ads.html) in the AWS Elemental MediaTailor user guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adconditioningconfiguration.html#cfn-mediatailor-playbackconfiguration-adconditioningconfiguration-streamingmediafileconditioning
	//
	StreamingMediaFileConditioning *string `field:"optional" json:"streamingMediaFileConditioning" yaml:"streamingMediaFileConditioning"`
}

