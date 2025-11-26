package previewawsrekognitionmixins


// Allows you to opt in or opt out to share data with Rekognition to improve model performance.
//
// You can choose this option at the account level or on a per-stream basis. Note that if you opt out at the account level, this setting is ignored on individual streams. For more information, see [StreamProcessorDataSharingPreference](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessorDataSharingPreference) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSharingPreferenceProperty := &DataSharingPreferenceProperty{
//   	OptIn: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-datasharingpreference.html
//
type CfnStreamProcessorPropsMixin_DataSharingPreferenceProperty struct {
	// Describes the opt-in status applied to a stream processor's data sharing policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-datasharingpreference.html#cfn-rekognition-streamprocessor-datasharingpreference-optin
	//
	OptIn interface{} `field:"optional" json:"optIn" yaml:"optIn"`
}

