package awsworkspaces


// Information about a WorkSpace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workspacePropertiesProperty := &workspacePropertiesProperty{
//   	computeTypeName: jsii.String("computeTypeName"),
//   	rootVolumeSizeGib: jsii.Number(123),
//   	runningMode: jsii.String("runningMode"),
//   	runningModeAutoStopTimeoutInMinutes: jsii.Number(123),
//   	userVolumeSizeGib: jsii.Number(123),
//   }
//
type CfnWorkspace_WorkspacePropertiesProperty struct {
	// The compute type.
	//
	// For more information, see [Amazon WorkSpaces Bundles](https://docs.aws.amazon.com/workspaces/details/#Amazon_WorkSpaces_Bundles) .
	ComputeTypeName *string `field:"optional" json:"computeTypeName" yaml:"computeTypeName"`
	// The size of the root volume.
	//
	// For important information about how to modify the size of the root and user volumes, see [Modify a WorkSpace](https://docs.aws.amazon.com/workspaces/latest/adminguide/modify-workspaces.html) .
	RootVolumeSizeGib *float64 `field:"optional" json:"rootVolumeSizeGib" yaml:"rootVolumeSizeGib"`
	// The running mode.
	//
	// For more information, see [Manage the WorkSpace Running Mode](https://docs.aws.amazon.com/workspaces/latest/adminguide/running-mode.html) .
	RunningMode *string `field:"optional" json:"runningMode" yaml:"runningMode"`
	// The time after a user logs off when WorkSpaces are automatically stopped.
	//
	// Configured in 60-minute intervals.
	RunningModeAutoStopTimeoutInMinutes *float64 `field:"optional" json:"runningModeAutoStopTimeoutInMinutes" yaml:"runningModeAutoStopTimeoutInMinutes"`
	// The size of the user storage.
	//
	// For important information about how to modify the size of the root and user volumes, see [Modify a WorkSpace](https://docs.aws.amazon.com/workspaces/latest/adminguide/modify-workspaces.html) .
	UserVolumeSizeGib *float64 `field:"optional" json:"userVolumeSizeGib" yaml:"userVolumeSizeGib"`
}

