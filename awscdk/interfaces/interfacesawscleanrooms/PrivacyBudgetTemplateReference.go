package interfacesawscleanrooms


// A reference to a PrivacyBudgetTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privacyBudgetTemplateReference := &PrivacyBudgetTemplateReference{
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	PrivacyBudgetTemplateArn: jsii.String("privacyBudgetTemplateArn"),
//   	PrivacyBudgetTemplateIdentifier: jsii.String("privacyBudgetTemplateIdentifier"),
//   }
//
type PrivacyBudgetTemplateReference struct {
	// The MembershipIdentifier of the PrivacyBudgetTemplate resource.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// The ARN of the PrivacyBudgetTemplate resource.
	PrivacyBudgetTemplateArn *string `field:"required" json:"privacyBudgetTemplateArn" yaml:"privacyBudgetTemplateArn"`
	// The PrivacyBudgetTemplateIdentifier of the PrivacyBudgetTemplate resource.
	PrivacyBudgetTemplateIdentifier *string `field:"required" json:"privacyBudgetTemplateIdentifier" yaml:"privacyBudgetTemplateIdentifier"`
}

