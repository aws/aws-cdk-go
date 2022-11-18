package awsfsx


// The different kinds of file system deployments used by Lustre.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//   var bucket s3.Bucket
//
//
//   lustreConfiguration := map[string]interface{}{
//   	"deploymentType": fsx.LustreDeploymentType_SCRATCH_2,
//   	"exportPath": bucket.s3UrlForObject(),
//   	"importPath": bucket.s3UrlForObject(),
//   	"autoImportPolicy": fsx.LustreAutoImportPolicy_NEW_CHANGED_DELETED,
//   }
//
//   fs := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &lustreFileSystemProps{
//   	vpc: vpc,
//   	vpcSubnet: vpc.privateSubnets[jsii.Number(0)],
//   	storageCapacityGiB: jsii.Number(1200),
//   	lustreConfiguration: lustreConfiguration,
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

