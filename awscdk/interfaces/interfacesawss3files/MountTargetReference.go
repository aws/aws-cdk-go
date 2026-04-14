package interfacesawss3files


// A reference to a MountTarget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mountTargetReference := &MountTargetReference{
//   	MountTargetId: jsii.String("mountTargetId"),
//   }
//
type MountTargetReference struct {
	// The MountTargetId of the MountTarget resource.
	MountTargetId *string `field:"required" json:"mountTargetId" yaml:"mountTargetId"`
}

