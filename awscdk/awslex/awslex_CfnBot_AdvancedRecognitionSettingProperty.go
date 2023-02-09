package awslex


// Provides settings that enable advanced recognition settings for slot values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedRecognitionSettingProperty := &advancedRecognitionSettingProperty{
//   	audioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   }
//
type CfnBot_AdvancedRecognitionSettingProperty struct {
	// Enables using the slot values as a custom vocabulary for recognizing user utterances.
	AudioRecognitionStrategy *string `field:"optional" json:"audioRecognitionStrategy" yaml:"audioRecognitionStrategy"`
}

