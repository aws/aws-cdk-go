package interfacesawsiot


// A reference to a PolicyPrincipalAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyPrincipalAttachmentReference := &PolicyPrincipalAttachmentReference{
//   	PolicyName: jsii.String("policyName"),
//   	Principal: jsii.String("principal"),
//   }
//
type PolicyPrincipalAttachmentReference struct {
	// The PolicyName of the PolicyPrincipalAttachment resource.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The Principal of the PolicyPrincipalAttachment resource.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

