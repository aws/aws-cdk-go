package awsamplify


// Describes the configuration details that apply to the jobs for an Amplify app.
//
// Use `JobConfig` to apply configuration to jobs, such as customizing the build instance size when you create or update an Amplify app. For more information about customizable build instances, see [Custom build instances](https://docs.aws.amazon.com/amplify/latest/userguide/custom-build-instance.html) in the *Amplify User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobConfigProperty := &JobConfigProperty{
//   	BuildComputeType: jsii.String("buildComputeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-jobconfig.html
//
type CfnApp_JobConfigProperty struct {
	// Specifies the size of the build instance.
	//
	// Amplify supports three instance sizes: `STANDARD_8GB` , `LARGE_16GB` , and `XLARGE_72GB` . If you don't specify a value, Amplify uses the `STANDARD_8GB` default.
	//
	// The following list describes the CPU, memory, and storage capacity for each build instance type:
	//
	// - **STANDARD_8GB** - - vCPUs: 4
	// - Memory: 8 GiB
	// - Disk space: 128 GB
	// - **LARGE_16GB** - - vCPUs: 8
	// - Memory: 16 GiB
	// - Disk space: 128 GB
	// - **XLARGE_72GB** - - vCPUs: 36
	// - Memory: 72 GiB
	// - Disk space: 256 GB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-jobconfig.html#cfn-amplify-app-jobconfig-buildcomputetype
	//
	BuildComputeType *string `field:"required" json:"buildComputeType" yaml:"buildComputeType"`
}

