package awscdkamplifyalpha


// Specifies the size of the build instance.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
//   	BuildComputeType: amplify.BuildComputeType_LARGE_16GB,
//   })
//
// Experimental.
type BuildComputeType string

const (
	// vCPUs: 4, Memory: 8 GiB, Disk space: 128 GB.
	// Experimental.
	BuildComputeType_STANDARD_8GB BuildComputeType = "STANDARD_8GB"
	// vCPUs: 8, Memory: 16 GiB, Disk space: 128 GB.
	// Experimental.
	BuildComputeType_LARGE_16GB BuildComputeType = "LARGE_16GB"
	// vCPUs: 36, Memory: 72 GiB, Disk space: 256 GB.
	// Experimental.
	BuildComputeType_XLARGE_72GB BuildComputeType = "XLARGE_72GB"
)

