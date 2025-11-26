package previewawsvoiceidevents


// Type definition for Data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   data := &Data{
//   	RegistrationSource: []*string{
//   		jsii.String("registrationSource"),
//   	},
//   	RegistrationSourceId: []*string{
//   		jsii.String("registrationSourceId"),
//   	},
//   	RegistrationStatus: []*string{
//   		jsii.String("registrationStatus"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdFraudsterAction_Data struct {
	// registrationSource property.
	//
	// Specify an array of string values to match this event if the actual value of registrationSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RegistrationSource *[]*string `field:"optional" json:"registrationSource" yaml:"registrationSource"`
	// registrationSourceId property.
	//
	// Specify an array of string values to match this event if the actual value of registrationSourceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RegistrationSourceId *[]*string `field:"optional" json:"registrationSourceId" yaml:"registrationSourceId"`
	// registrationStatus property.
	//
	// Specify an array of string values to match this event if the actual value of registrationStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RegistrationStatus *[]*string `field:"optional" json:"registrationStatus" yaml:"registrationStatus"`
}

