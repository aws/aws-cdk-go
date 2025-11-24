package mixinsawsmediapackagev2


// Details about the content that you want MediaPackage to pass through in the manifest to the playback device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dashProgramInformationProperty := &DashProgramInformationProperty{
//   	Copyright: jsii.String("copyright"),
//   	LanguageCode: jsii.String("languageCode"),
//   	MoreInformationUrl: jsii.String("moreInformationUrl"),
//   	Source: jsii.String("source"),
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashprograminformation.html
//
type CfnOriginEndpointPropsMixin_DashProgramInformationProperty struct {
	// A copyright statement about the content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashprograminformation.html#cfn-mediapackagev2-originendpoint-dashprograminformation-copyright
	//
	Copyright *string `field:"optional" json:"copyright" yaml:"copyright"`
	// The language code for this manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashprograminformation.html#cfn-mediapackagev2-originendpoint-dashprograminformation-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// An absolute URL that contains more information about this content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashprograminformation.html#cfn-mediapackagev2-originendpoint-dashprograminformation-moreinformationurl
	//
	MoreInformationUrl *string `field:"optional" json:"moreInformationUrl" yaml:"moreInformationUrl"`
	// Information about the content provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashprograminformation.html#cfn-mediapackagev2-originendpoint-dashprograminformation-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The title for the manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashprograminformation.html#cfn-mediapackagev2-originendpoint-dashprograminformation-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

