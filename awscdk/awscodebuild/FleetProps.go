package awscodebuild


// Construction properties of a CodeBuild {@link Fleet}.
//
// Example:
//   fleet := codebuild.NewFleet(this, jsii.String("Fleet"), &FleetProps{
//   	ComputeType: codebuild.FleetComputeType_MEDIUM,
//   	EnvironmentType: codebuild.EnvironmentType_LINUX_CONTAINER,
//   	BaseCapacity: jsii.Number(1),
//   })
//
//   codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		Fleet: *Fleet,
//   		BuildImage: codebuild.LinuxBuildImage_STANDARD_7_0(),
//   	},
//   })
//
type FleetProps struct {
	// The number of machines allocated to the compute ﬂeet. Deﬁnes the number of builds that can run in parallel.
	//
	// Minimum value of 1.
	BaseCapacity *float64 `field:"required" json:"baseCapacity" yaml:"baseCapacity"`
	// The instance type of the compute fleet.
	// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_codebuild.ComputeType.html
	//
	ComputeType FleetComputeType `field:"required" json:"computeType" yaml:"computeType"`
	// The build environment (operating system/architecture/accelerator) type made available to projects using this fleet.
	EnvironmentType EnvironmentType `field:"required" json:"environmentType" yaml:"environmentType"`
	// The name of the Fleet.
	// Default: - CloudFormation generated name.
	//
	FleetName *string `field:"optional" json:"fleetName" yaml:"fleetName"`
}

