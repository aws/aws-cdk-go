package awsfsx


// The Lustre version for the file system.
//
// Example:
//   var vpc vpc
//
//
//   fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
//   	LustreConfiguration: &LustreConfiguration{
//   		DeploymentType: fsx.LustreDeploymentType_SCRATCH_2,
//   	},
//   	StorageCapacityGiB: jsii.Number(1200),
//   	Vpc: Vpc,
//   	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
//   	FileSystemTypeVersion: fsx.FileSystemTypeVersion_V_2_15,
//   })
//
type FileSystemTypeVersion string

const (
	// Version 2.10.
	FileSystemTypeVersion_V_2_10 FileSystemTypeVersion = "V_2_10"
	// Version 2.12.
	FileSystemTypeVersion_V_2_12 FileSystemTypeVersion = "V_2_12"
	// Version 2.15.
	FileSystemTypeVersion_V_2_15 FileSystemTypeVersion = "V_2_15"
)

