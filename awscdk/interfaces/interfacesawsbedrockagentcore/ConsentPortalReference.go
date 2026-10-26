package interfacesawsbedrockagentcore


// A reference to a ConsentPortal resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   consentPortalReference := &ConsentPortalReference{
//   	ConsentPortalArn: jsii.String("consentPortalArn"),
//   }
//
type ConsentPortalReference struct {
	// The ConsentPortalArn of the ConsentPortal resource.
	ConsentPortalArn *string `field:"required" json:"consentPortalArn" yaml:"consentPortalArn"`
}

