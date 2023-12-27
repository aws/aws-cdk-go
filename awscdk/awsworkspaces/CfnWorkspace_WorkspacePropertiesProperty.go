package awsworkspaces


// Information about a WorkSpace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workspacePropertiesProperty := &WorkspacePropertiesProperty{
//   	ComputeTypeName: jsii.String("computeTypeName"),
//   	RootVolumeSizeGib: jsii.Number(123),
//   	RunningMode: jsii.String("runningMode"),
//   	RunningModeAutoStopTimeoutInMinutes: jsii.Number(123),
//   	UserVolumeSizeGib: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html
//
type CfnWorkspace_WorkspacePropertiesProperty struct {
	// The compute type.
	//
	// For more information, see [Amazon WorkSpaces Bundles](https://docs.aws.amazon.com/workspaces/details/#Amazon_WorkSpaces_Bundles) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html#cfn-workspaces-workspace-workspaceproperties-computetypename
	//
	ComputeTypeName *string `field:"optional" json:"computeTypeName" yaml:"computeTypeName"`
	// The size of the root volume.
	//
	// For important information about how to modify the size of the root and user volumes, see [Modify a WorkSpace](https://docs.aws.amazon.com/workspaces/latest/adminguide/modify-workspaces.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html#cfn-workspaces-workspace-workspaceproperties-rootvolumesizegib
	//
	RootVolumeSizeGib *float64 `field:"optional" json:"rootVolumeSizeGib" yaml:"rootVolumeSizeGib"`
	// The running mode. For more information, see [Manage the WorkSpace Running Mode](https://docs.aws.amazon.com/workspaces/latest/adminguide/running-mode.html) .
	//
	// > - The `MANUAL` value is only supported by Amazon WorkSpaces Core. Contact your account team to be allow-listed to use this value. For more information, see [Amazon WorkSpaces Core](https://docs.aws.amazon.com/workspaces/core/) .
	// > - Ensure you review your running mode to ensure you are using a running mode that is optimal for your needs and budget. For more information on switching running modes, see [Can I switch between hourly and monthly billing?](https://docs.aws.amazon.com/https://aws.amazon.com/workspaces/faqs/#:~:text=Q%3A%20Can%20I%20switch%20between%20hourly%20and%20monthly%20billing%3F)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html#cfn-workspaces-workspace-workspaceproperties-runningmode
	//
	RunningMode *string `field:"optional" json:"runningMode" yaml:"runningMode"`
	// The time after a user logs off when WorkSpaces are automatically stopped.
	//
	// Configured in 60-minute intervals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html#cfn-workspaces-workspace-workspaceproperties-runningmodeautostoptimeoutinminutes
	//
	RunningModeAutoStopTimeoutInMinutes *float64 `field:"optional" json:"runningModeAutoStopTimeoutInMinutes" yaml:"runningModeAutoStopTimeoutInMinutes"`
	// The size of the user storage.
	//
	// For important information about how to modify the size of the root and user volumes, see [Modify a WorkSpace](https://docs.aws.amazon.com/workspaces/latest/adminguide/modify-workspaces.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html#cfn-workspaces-workspace-workspaceproperties-uservolumesizegib
	//
	UserVolumeSizeGib *float64 `field:"optional" json:"userVolumeSizeGib" yaml:"userVolumeSizeGib"`
}

