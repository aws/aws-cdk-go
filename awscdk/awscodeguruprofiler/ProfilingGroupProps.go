package awscodeguruprofiler


// Properties for creating a new Profiling Group.
//
// Example:
//   profilingGroup := codeguruprofiler.NewProfilingGroup(this, jsii.String("MyProfilingGroup"), &ProfilingGroupProps{
//   	ComputePlatform: codeguruprofiler.ComputePlatform_AWS_LAMBDA,
//   })
//
type ProfilingGroupProps struct {
	// The compute platform of the profiling group.
	// Default: ComputePlatform.DEFAULT
	//
	ComputePlatform ComputePlatform `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// A name for the profiling group.
	// Default: - automatically generated name.
	//
	ProfilingGroupName *string `field:"optional" json:"profilingGroupName" yaml:"profilingGroupName"`
}

