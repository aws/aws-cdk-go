package awsfsx


// The different kinds of file system deployments used by Lustre.
//
// Example:
//   var vpc vpc
//
//
//   fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &lustreFileSystemProps{
//   	lustreConfiguration: &lustreConfiguration{
//   		deploymentType: fsx.lustreDeploymentType_SCRATCH_2,
//   	},
//   	storageCapacityGiB: jsii.Number(1200),
//   	vpc: vpc,
//   	vpcSubnet: vpc.privateSubnets[jsii.Number(0)],
//   })
//
// Experimental.
type LustreDeploymentType string

const (
	// Original type for shorter term data processing.
	//
	// Data is not replicated and does not persist on server fail.
	// Experimental.
	LustreDeploymentType_SCRATCH_1 LustreDeploymentType = "SCRATCH_1"
	// Newer type for shorter term data processing.
	//
	// Data is not replicated and does not persist on server fail.
	// Provides better support for spiky workloads.
	// Experimental.
	LustreDeploymentType_SCRATCH_2 LustreDeploymentType = "SCRATCH_2"
	// Long term storage.
	//
	// Data is replicated and file servers are replaced if they fail.
	// Experimental.
	LustreDeploymentType_PERSISTENT_1 LustreDeploymentType = "PERSISTENT_1"
	// Newer type of long term storage with higher throughput tiers.
	//
	// Data is replicated and file servers are replaced if they fail.
	// Experimental.
	LustreDeploymentType_PERSISTENT_2 LustreDeploymentType = "PERSISTENT_2"
)

