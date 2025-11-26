package previewawsvoiceidevents


// Type definition for Data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   data := &Data{
//   	EnrollmentSource: []*string{
//   		jsii.String("enrollmentSource"),
//   	},
//   	EnrollmentSourceId: []*string{
//   		jsii.String("enrollmentSourceId"),
//   	},
//   	EnrollmentStatus: []*string{
//   		jsii.String("enrollmentStatus"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdSpeakerAction_Data struct {
	// enrollmentSource property.
	//
	// Specify an array of string values to match this event if the actual value of enrollmentSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnrollmentSource *[]*string `field:"optional" json:"enrollmentSource" yaml:"enrollmentSource"`
	// enrollmentSourceId property.
	//
	// Specify an array of string values to match this event if the actual value of enrollmentSourceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnrollmentSourceId *[]*string `field:"optional" json:"enrollmentSourceId" yaml:"enrollmentSourceId"`
	// enrollmentStatus property.
	//
	// Specify an array of string values to match this event if the actual value of enrollmentStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnrollmentStatus *[]*string `field:"optional" json:"enrollmentStatus" yaml:"enrollmentStatus"`
}

