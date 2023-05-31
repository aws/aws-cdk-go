package awsshield


// Contact information that the SRT can use to contact you if you have proactive engagement enabled, for escalations to the SRT and to initiate proactive customer support.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emergencyContactProperty := &EmergencyContactProperty{
//   	EmailAddress: jsii.String("emailAddress"),
//
//   	// the properties below are optional
//   	ContactNotes: jsii.String("contactNotes"),
//   	PhoneNumber: jsii.String("phoneNumber"),
//   }
//
type CfnProactiveEngagement_EmergencyContactProperty struct {
	// The email address for the contact.
	EmailAddress *string `field:"required" json:"emailAddress" yaml:"emailAddress"`
	// Additional notes regarding the contact.
	ContactNotes *string `field:"optional" json:"contactNotes" yaml:"contactNotes"`
	// The phone number for the contact.
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

