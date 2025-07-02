package awsworkspacesinstances


// Properties for defining a `CfnVolumeAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVolumeAssociationProps := &CfnVolumeAssociationProps{
//   	Device: jsii.String("device"),
//   	VolumeId: jsii.String("volumeId"),
//   	WorkspaceInstanceId: jsii.String("workspaceInstanceId"),
//
//   	// the properties below are optional
//   	DisassociateMode: jsii.String("disassociateMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-volumeassociation.html
//
type CfnVolumeAssociationProps struct {
	// The device name for the volume attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-volumeassociation.html#cfn-workspacesinstances-volumeassociation-device
	//
	Device *string `field:"required" json:"device" yaml:"device"`
	// ID of the volume to attach to the workspace instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-volumeassociation.html#cfn-workspacesinstances-volumeassociation-volumeid
	//
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// ID of the workspace instance to associate with the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-volumeassociation.html#cfn-workspacesinstances-volumeassociation-workspaceinstanceid
	//
	WorkspaceInstanceId *string `field:"required" json:"workspaceInstanceId" yaml:"workspaceInstanceId"`
	// Mode to use when disassociating the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-volumeassociation.html#cfn-workspacesinstances-volumeassociation-disassociatemode
	//
	DisassociateMode *string `field:"optional" json:"disassociateMode" yaml:"disassociateMode"`
}

