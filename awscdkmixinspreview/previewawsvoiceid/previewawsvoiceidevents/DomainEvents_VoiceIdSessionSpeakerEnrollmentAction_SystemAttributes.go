package previewawsvoiceidevents


// Type definition for SystemAttributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   systemAttributes := &SystemAttributes{
//   	AwsConnectOriginalContactArn: []*string{
//   		jsii.String("awsConnectOriginalContactArn"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdSessionSpeakerEnrollmentAction_SystemAttributes struct {
	// aws-connect-OriginalContactArn property.
	//
	// Specify an array of string values to match this event if the actual value of aws-connect-OriginalContactArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsConnectOriginalContactArn *[]*string `field:"optional" json:"awsConnectOriginalContactArn" yaml:"awsConnectOriginalContactArn"`
}

