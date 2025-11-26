package previewawsec2events


// Type definition for CpuOptions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cpuOptions := &CpuOptions{
//   	CoreCount: []*string{
//   		jsii.String("coreCount"),
//   	},
//   	ThreadsPerCore: []*string{
//   		jsii.String("threadsPerCore"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_CpuOptions struct {
	// coreCount property.
	//
	// Specify an array of string values to match this event if the actual value of coreCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CoreCount *[]*string `field:"optional" json:"coreCount" yaml:"coreCount"`
	// threadsPerCore property.
	//
	// Specify an array of string values to match this event if the actual value of threadsPerCore is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ThreadsPerCore *[]*string `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

