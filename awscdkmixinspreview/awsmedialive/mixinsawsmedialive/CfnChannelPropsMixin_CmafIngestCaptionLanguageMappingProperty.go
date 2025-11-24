package mixinsawsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cmafIngestCaptionLanguageMappingProperty := &CmafIngestCaptionLanguageMappingProperty{
//   	CaptionChannel: jsii.Number(123),
//   	LanguageCode: jsii.String("languageCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestcaptionlanguagemapping.html
//
type CfnChannelPropsMixin_CmafIngestCaptionLanguageMappingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestcaptionlanguagemapping.html#cfn-medialive-channel-cmafingestcaptionlanguagemapping-captionchannel
	//
	CaptionChannel *float64 `field:"optional" json:"captionChannel" yaml:"captionChannel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestcaptionlanguagemapping.html#cfn-medialive-channel-cmafingestcaptionlanguagemapping-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
}

