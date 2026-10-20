package awsworkspaces


// Describes the self-service permissions for a directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfservicePermissionsProperty := &SelfservicePermissionsProperty{
//   	ChangeComputeType: jsii.String("changeComputeType"),
//   	IncreaseVolumeSize: jsii.String("increaseVolumeSize"),
//   	RebuildWorkspace: jsii.String("rebuildWorkspace"),
//   	RestartWorkspace: jsii.String("restartWorkspace"),
//   	SwitchRunningMode: jsii.String("switchRunningMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-selfservicepermissions.html
//
type CfnDirectory_SelfservicePermissionsProperty struct {
	// Specifies whether users can change the compute type (bundle) for their WorkSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-selfservicepermissions.html#cfn-workspaces-directory-selfservicepermissions-changecomputetype
	//
	ChangeComputeType *string `field:"optional" json:"changeComputeType" yaml:"changeComputeType"`
	// Specifies whether users can increase the volume size of the drives on their WorkSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-selfservicepermissions.html#cfn-workspaces-directory-selfservicepermissions-increasevolumesize
	//
	IncreaseVolumeSize *string `field:"optional" json:"increaseVolumeSize" yaml:"increaseVolumeSize"`
	// Specifies whether users can rebuild the operating system of a WorkSpace to its original state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-selfservicepermissions.html#cfn-workspaces-directory-selfservicepermissions-rebuildworkspace
	//
	RebuildWorkspace *string `field:"optional" json:"rebuildWorkspace" yaml:"rebuildWorkspace"`
	// Specifies whether users can restart their WorkSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-selfservicepermissions.html#cfn-workspaces-directory-selfservicepermissions-restartworkspace
	//
	RestartWorkspace *string `field:"optional" json:"restartWorkspace" yaml:"restartWorkspace"`
	// Specifies whether users can switch the running mode of their WorkSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-selfservicepermissions.html#cfn-workspaces-directory-selfservicepermissions-switchrunningmode
	//
	SwitchRunningMode *string `field:"optional" json:"switchRunningMode" yaml:"switchRunningMode"`
}

