package awsssmcontacts


// The contact that Incident Manager is engaging during an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactTargetInfoProperty := &contactTargetInfoProperty{
//   	contactId: jsii.String("contactId"),
//   	isEssential: jsii.Boolean(false),
//   }
//
type CfnContact_ContactTargetInfoProperty struct {
	// The Amazon Resource Name (ARN) of the contact.
	ContactId *string `field:"required" json:"contactId" yaml:"contactId"`
	// A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.
	IsEssential interface{} `field:"required" json:"isEssential" yaml:"isEssential"`
}

