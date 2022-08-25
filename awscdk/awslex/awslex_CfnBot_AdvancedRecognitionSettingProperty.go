package awslex


// Specifies settings that enable advanced audio recognition for slot values.
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
	// Specifies that Amazon Lex should use slot values as a custom vocabulary when recognizing user utterances.
	AudioRecognitionStrategy *string `field:"optional" json:"audioRecognitionStrategy" yaml:"audioRecognitionStrategy"`
}

