package awsshield


// Properties for defining a `CfnProactiveEngagement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProactiveEngagementProps := &CfnProactiveEngagementProps{
//   	EmergencyContactList: []interface{}{
//   		&EmergencyContactProperty{
//   			EmailAddress: jsii.String("emailAddress"),
//
//   			// the properties below are optional
//   			ContactNotes: jsii.String("contactNotes"),
//   			PhoneNumber: jsii.String("phoneNumber"),
//   		},
//   	},
//   	ProactiveEngagementStatus: jsii.String("proactiveEngagementStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-proactiveengagement.html
//
type CfnProactiveEngagementProps struct {
	// The list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you for escalations to the SRT and to initiate proactive customer support, plus any relevant notes.
	//
	// To enable proactive engagement, the contact list must include at least one phone number.
	//
	// If you provide more than one contact, in the notes, indicate the circumstances under which each contact should be used. Include primary and secondary contact designations, and provide the hours of availability and time zones for each contact.
	//
	// Example contact notes:
	//
	// - This is a hotline that's staffed 24x7x365. Please work with the responding analyst and they will get the appropriate person on the call.
	// - Please contact the secondary phone number if the hotline doesn't respond within 5 minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-proactiveengagement.html#cfn-shield-proactiveengagement-emergencycontactlist
	//
	EmergencyContactList interface{} `field:"required" json:"emergencyContactList" yaml:"emergencyContactList"`
	// Specifies whether proactive engagement is enabled or disabled.
	//
	// Valid values:
	//
	// `ENABLED` - The Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
	//
	// `DISABLED` - The SRT will not proactively notify contacts about escalations or to initiate proactive customer support.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-proactiveengagement.html#cfn-shield-proactiveengagement-proactiveengagementstatus
	//
	ProactiveEngagementStatus *string `field:"required" json:"proactiveEngagementStatus" yaml:"proactiveEngagementStatus"`
}

