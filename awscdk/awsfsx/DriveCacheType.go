package awsfsx


// The type of drive cache used by PERSISTENT_1 file systems that are provisioned with HDD storage devices.
//
// Example:
//   var vpc Vpc
//
//
//   fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
//   	LustreConfiguration: &LustreConfiguration{
//   		DeploymentType: fsx.LustreDeploymentType_PERSISTENT_1,
//   		DriveCacheType: fsx.DriveCacheType_READ,
//   	},
//   	StorageCapacityGiB: jsii.Number(1200),
//   	Vpc: Vpc,
//   	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
//   	StorageType: fsx.StorageType_HDD,
//   })
//
type DriveCacheType string

const (
	// The Lustre file system is configured with no data cache.
	DriveCacheType_NONE DriveCacheType = "NONE"
	// The Lustre file system is configured with a read cache.
	DriveCacheType_READ DriveCacheType = "READ"
)

