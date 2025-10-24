package awsfsx


// The storage type for the file system.
//
// Example:
//   var vpc Vpc
//
//
//   fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
//   	LustreConfiguration: &LustreConfiguration{
//   		DeploymentType: fsx.LustreDeploymentType_PERSISTENT_1,
//   	},
//   	StorageCapacityGiB: jsii.Number(1200),
//   	Vpc: Vpc,
//   	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
//   	StorageType: fsx.StorageType_HDD,
//   })
//
type StorageType string

const (
	// Solid State Drive storage.
	StorageType_SSD StorageType = "SSD"
	// Hard Disk Drive storage.
	StorageType_HDD StorageType = "HDD"
	// Intelligent Tiering storage.
	StorageType_INTELLIGENT_TIERING StorageType = "INTELLIGENT_TIERING"
)

