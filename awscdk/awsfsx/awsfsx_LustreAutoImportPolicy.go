package awsfsx


// The different auto import policies which are allowed.
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
type LustreAutoImportPolicy string

const (
	// AutoImport is off.
	//
	// Amazon FSx only updates file and directory listings from the linked S3 bucket when the file system is created. FSx does not update file and directory listings for any new or changed objects after choosing this option.
	LustreAutoImportPolicy_NONE LustreAutoImportPolicy = "NONE"
	// AutoImport is on.
	//
	// Amazon FSx automatically imports directory listings of any new objects added to the linked S3 bucket that do not currently exist in the FSx file system.
	LustreAutoImportPolicy_NEW LustreAutoImportPolicy = "NEW"
	// AutoImport is on.
	//
	// Amazon FSx automatically imports file and directory listings of any new objects added to the S3 bucket and any existing objects that are changed in the S3 bucket after you choose this option.
	LustreAutoImportPolicy_NEW_CHANGED LustreAutoImportPolicy = "NEW_CHANGED"
	// AutoImport is on.
	//
	// Amazon FSx automatically imports file and directory listings of any new objects added to the S3 bucket, any existing objects that are changed in the S3 bucket, and any objects that were deleted in the S3 bucket.
	LustreAutoImportPolicy_NEW_CHANGED_DELETED LustreAutoImportPolicy = "NEW_CHANGED_DELETED"
)

