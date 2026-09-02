package interfacesawscloudformation


// A reference to a ChangeSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   changeSetReference := &ChangeSetReference{
//   	ChangeSetId: jsii.String("changeSetId"),
//   }
//
type ChangeSetReference struct {
	// The ChangeSetId of the ChangeSet resource.
	ChangeSetId *string `field:"required" json:"changeSetId" yaml:"changeSetId"`
}

