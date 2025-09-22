package awssecurityhub


// A reference to a DelegatedAdmin resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   delegatedAdminReference := &DelegatedAdminReference{
//   	DelegatedAdminIdentifier: jsii.String("delegatedAdminIdentifier"),
//   }
//
type DelegatedAdminReference struct {
	// The DelegatedAdminIdentifier of the DelegatedAdmin resource.
	DelegatedAdminIdentifier *string `field:"required" json:"delegatedAdminIdentifier" yaml:"delegatedAdminIdentifier"`
}

