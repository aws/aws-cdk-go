package awsworkspacesinstances


// A reference to a VolumeAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeAssociationReference := &VolumeAssociationReference{
//   	Device: jsii.String("device"),
//   	VolumeId: jsii.String("volumeId"),
//   	WorkspaceInstanceId: jsii.String("workspaceInstanceId"),
//   }
//
type VolumeAssociationReference struct {
	// The Device of the VolumeAssociation resource.
	Device *string `field:"required" json:"device" yaml:"device"`
	// The VolumeId of the VolumeAssociation resource.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// The WorkspaceInstanceId of the VolumeAssociation resource.
	WorkspaceInstanceId *string `field:"required" json:"workspaceInstanceId" yaml:"workspaceInstanceId"`
}

