package awsfsx


// The different kinds of file system deployments used by Lustre.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//   var bucket bucket
//
//
//   lustreConfiguration := map[string]interface{}{
//   	"deploymentType": fsx.LustreDeploymentType_SCRATCH_2,
//   	"exportPath": bucket.s3UrlForObject(),
//   	"importPath": bucket.s3UrlForObject(),
//   	"autoImportPolicy": fsx.LustreAutoImportPolicy_NEW_CHANGED_DELETED,
//   }
//
//   fs := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
//   	Vpc: vpc,
//   	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
//   	StorageCapacityGiB: jsii.Number(1200),
//   	LustreConfiguration: LustreConfiguration,
//   })
//
type LustreDeploymentType string

const (
	// Original type for shorter term data processing.
	//
	// Data is not replicated and does not persist on server fail.
	LustreDeploymentType_SCRATCH_1 LustreDeploymentType = "SCRATCH_1"
	// Newer type for shorter term data processing.
	//
	// Data is not replicated and does not persist on server fail.
	// Provides better support for spiky workloads.
	LustreDeploymentType_SCRATCH_2 LustreDeploymentType = "SCRATCH_2"
	// Long term storage.
	//
	// Data is replicated and file servers are replaced if they fail.
	LustreDeploymentType_PERSISTENT_1 LustreDeploymentType = "PERSISTENT_1"
	// Newer type of long term storage with higher throughput tiers.
	//
	// Data is replicated and file servers are replaced if they fail.
	LustreDeploymentType_PERSISTENT_2 LustreDeploymentType = "PERSISTENT_2"
)

