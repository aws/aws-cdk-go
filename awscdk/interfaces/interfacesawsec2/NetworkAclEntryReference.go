package interfacesawsec2


// A reference to a NetworkAclEntry resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkAclEntryReference := &NetworkAclEntryReference{
//   	NetworkAclEntryId: jsii.String("networkAclEntryId"),
//   }
//
type NetworkAclEntryReference struct {
	// The Id of the NetworkAclEntry resource.
	NetworkAclEntryId *string `field:"required" json:"networkAclEntryId" yaml:"networkAclEntryId"`
}

