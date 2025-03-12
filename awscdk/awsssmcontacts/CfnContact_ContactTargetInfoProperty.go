package awsssmcontacts


// The contact that Incident Manager is engaging during an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactTargetInfoProperty := &ContactTargetInfoProperty{
//   	ContactId: jsii.String("contactId"),
//   	IsEssential: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-contact-contacttargetinfo.html
//
type CfnContact_ContactTargetInfoProperty struct {
	// The Amazon Resource Name (ARN) of the contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-contact-contacttargetinfo.html#cfn-ssmcontacts-contact-contacttargetinfo-contactid
	//
	ContactId *string `field:"required" json:"contactId" yaml:"contactId"`
	// A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-contact-contacttargetinfo.html#cfn-ssmcontacts-contact-contacttargetinfo-isessential
	//
	IsEssential interface{} `field:"required" json:"isEssential" yaml:"isEssential"`
}

