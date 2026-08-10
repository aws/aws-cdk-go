package previewawstranscribeevents


// Type definition for Settings_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   settings1 := &Settings1{
//   	ChannelIdentification: []*string{
//   		jsii.String("channelIdentification"),
//   	},
//   	VocabularyName: []*string{
//   		jsii.String("vocabularyName"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Settings1 struct {
	// channelIdentification property.
	//
	// Specify an array of string values to match this event if the actual value of channelIdentification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChannelIdentification *[]*string `field:"optional" json:"channelIdentification" yaml:"channelIdentification"`
	// vocabularyName property.
	//
	// Specify an array of string values to match this event if the actual value of vocabularyName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VocabularyName *[]*string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}

