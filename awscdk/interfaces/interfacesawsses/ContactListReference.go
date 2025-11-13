package interfacesawsses


// A reference to a ContactList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactListReference := &ContactListReference{
//   	ContactListName: jsii.String("contactListName"),
//   }
//
type ContactListReference struct {
	// The ContactListName of the ContactList resource.
	ContactListName *string `field:"required" json:"contactListName" yaml:"contactListName"`
}

