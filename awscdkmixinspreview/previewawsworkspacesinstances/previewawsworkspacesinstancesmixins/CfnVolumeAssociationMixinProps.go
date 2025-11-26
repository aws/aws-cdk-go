package previewawsworkspacesinstancesmixins


// Properties for CfnVolumeAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVolumeAssociationMixinProps := &CfnVolumeAssociationMixinProps{
//   	Device: jsii.String("device"),
//   	DisassociateMode: jsii.String("disassociateMode"),
//   	VolumeId: jsii.String("volumeId"),
//   	WorkspaceInstanceId: jsii.String("workspaceInstanceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-volumeassociation.html
//
type CfnVolumeAssociationMixinProps struct {
	// The device name for the volume attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-volumeassociation.html#cfn-workspacesinstances-volumeassociation-device
	//
	Device *string `field:"optional" json:"device" yaml:"device"`
	// Mode to use when disassociating the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-volumeassociation.html#cfn-workspacesinstances-volumeassociation-disassociatemode
	//
	DisassociateMode *string `field:"optional" json:"disassociateMode" yaml:"disassociateMode"`
	// ID of the volume to attach to the workspace instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-volumeassociation.html#cfn-workspacesinstances-volumeassociation-volumeid
	//
	VolumeId *string `field:"optional" json:"volumeId" yaml:"volumeId"`
	// ID of the workspace instance to associate with the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-volumeassociation.html#cfn-workspacesinstances-volumeassociation-workspaceinstanceid
	//
	WorkspaceInstanceId *string `field:"optional" json:"workspaceInstanceId" yaml:"workspaceInstanceId"`
}

