package awsdatasync


// Represents the options that are available to control the behavior of a [StartTaskExecution](https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html) operation. This behavior includes preserving metadata, such as user ID (UID), group ID (GID), and file permissions; overwriting files in the destination; data integrity verification; and so on.
//
// A task has a set of default options associated with it. If you don't specify an option in [StartTaskExecution](https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html) , the default value is used. You can override the default options on each task execution by specifying an overriding `Options` value to [StartTaskExecution](https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionsProperty := &OptionsProperty{
//   	Atime: jsii.String("atime"),
//   	BytesPerSecond: jsii.Number(123),
//   	Gid: jsii.String("gid"),
//   	LogLevel: jsii.String("logLevel"),
//   	Mtime: jsii.String("mtime"),
//   	ObjectTags: jsii.String("objectTags"),
//   	OverwriteMode: jsii.String("overwriteMode"),
//   	PosixPermissions: jsii.String("posixPermissions"),
//   	PreserveDeletedFiles: jsii.String("preserveDeletedFiles"),
//   	PreserveDevices: jsii.String("preserveDevices"),
//   	SecurityDescriptorCopyFlags: jsii.String("securityDescriptorCopyFlags"),
//   	TaskQueueing: jsii.String("taskQueueing"),
//   	TransferMode: jsii.String("transferMode"),
//   	Uid: jsii.String("uid"),
//   	VerifyMode: jsii.String("verifyMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html
//
type CfnTask_OptionsProperty struct {
	// A file metadata value that shows the last time that a file was accessed (that is, when the file was read or written to).
	//
	// If you set `Atime` to `BEST_EFFORT` , AWS DataSync attempts to preserve the original `Atime` attribute on all source files (that is, the version before the PREPARING phase). However, `Atime` 's behavior is not fully standard across platforms, so AWS DataSync can only do this on a best-effort basis.
	//
	// Default value: `BEST_EFFORT`
	//
	// `BEST_EFFORT` : Attempt to preserve the per-file `Atime` value (recommended).
	//
	// `NONE` : Ignore `Atime` .
	//
	// > If `Atime` is set to `BEST_EFFORT` , `Mtime` must be set to `PRESERVE` .
	// >
	// > If `Atime` is set to `NONE` , `Mtime` must also be `NONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-atime
	//
	Atime *string `field:"optional" json:"atime" yaml:"atime"`
	// A value that limits the bandwidth used by AWS DataSync .
	//
	// For example, if you want AWS DataSync to use a maximum of 1 MB, set this value to `1048576` (=1024*1024).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-bytespersecond
	//
	BytesPerSecond *float64 `field:"optional" json:"bytesPerSecond" yaml:"bytesPerSecond"`
	// The group ID (GID) of the file's owners.
	//
	// Default value: `INT_VALUE`
	//
	// `INT_VALUE` : Preserve the integer value of the user ID (UID) and group ID (GID) (recommended).
	//
	// `NAME` : Currently not supported.
	//
	// `NONE` : Ignore the UID and GID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-gid
	//
	Gid *string `field:"optional" json:"gid" yaml:"gid"`
	// Specifies the type of logs that DataSync publishes to a Amazon CloudWatch Logs log group.
	//
	// To specify the log group, see [CloudWatchLogGroupArn](https://docs.aws.amazon.com/datasync/latest/userguide/API_CreateTask.html#DataSync-CreateTask-request-CloudWatchLogGroupArn) .
	//
	// - `BASIC` - Publishes logs with only basic information (such as transfer errors).
	// - `TRANSFER` - Publishes logs for all files or objects that your DataSync task transfers and performs data-integrity checks on.
	// - `OFF` - No logs are published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-loglevel
	//
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// A value that indicates the last time that a file was modified (that is, a file was written to) before the PREPARING phase.
	//
	// This option is required for cases when you need to run the same task more than one time.
	//
	// Default value: `PRESERVE`
	//
	// `PRESERVE` : Preserve original `Mtime` (recommended)
	//
	// `NONE` : Ignore `Mtime` .
	//
	// > If `Mtime` is set to `PRESERVE` , `Atime` must be set to `BEST_EFFORT` .
	// >
	// > If `Mtime` is set to `NONE` , `Atime` must also be set to `NONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-mtime
	//
	Mtime *string `field:"optional" json:"mtime" yaml:"mtime"`
	// Specifies whether you want DataSync to `PRESERVE` object tags (default behavior) when transferring between object storage systems.
	//
	// If you want your DataSync task to ignore object tags, specify the `NONE` value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-objecttags
	//
	ObjectTags *string `field:"optional" json:"objectTags" yaml:"objectTags"`
	// Specifies whether DataSync should modify or preserve data at the destination location.
	//
	// - `ALWAYS` (default) - DataSync modifies data in the destination location when source data (including metadata) has changed.
	//
	// If DataSync overwrites objects, you might incur additional charges for certain Amazon S3 storage classes (for example, for retrieval or early deletion). For more information, see [Storage class considerations with Amazon S3 transfers](https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes) .
	// - `NEVER` - DataSync doesn't overwrite data in the destination location even if the source data has changed. You can use this option to protect against overwriting changes made to files or objects in the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-overwritemode
	//
	OverwriteMode *string `field:"optional" json:"overwriteMode" yaml:"overwriteMode"`
	// A value that determines which users or groups can access a file for a specific purpose, such as reading, writing, or execution of the file.
	//
	// This option should be set only for Network File System (NFS), Amazon EFS, and Amazon S3 locations. For more information about what metadata is copied by DataSync, see [Metadata Copied by DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/special-files.html#metadata-copied) .
	//
	// Default value: `PRESERVE`
	//
	// `PRESERVE` : Preserve POSIX-style permissions (recommended).
	//
	// `NONE` : Ignore permissions.
	//
	// > AWS DataSync can preserve extant permissions of a source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-posixpermissions
	//
	PosixPermissions *string `field:"optional" json:"posixPermissions" yaml:"posixPermissions"`
	// A value that specifies whether files in the destination that don't exist in the source file system are preserved.
	//
	// This option can affect your storage costs. If your task deletes objects, you might incur minimum storage duration charges for certain storage classes. For detailed information, see [Considerations when working with Amazon S3 storage classes in DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes) in the *AWS DataSync User Guide* .
	//
	// Default value: `PRESERVE`
	//
	// `PRESERVE` : Ignore destination files that aren't present in the source (recommended).
	//
	// `REMOVE` : Delete destination files that aren't present in the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-preservedeletedfiles
	//
	PreserveDeletedFiles *string `field:"optional" json:"preserveDeletedFiles" yaml:"preserveDeletedFiles"`
	// A value that determines whether AWS DataSync should preserve the metadata of block and character devices in the source file system, and re-create the files with that device name and metadata on the destination.
	//
	// DataSync does not copy the contents of such devices, only the name and metadata.
	//
	// > AWS DataSync can't sync the actual contents of such devices, because they are nonterminal and don't return an end-of-file (EOF) marker.
	//
	// Default value: `NONE`
	//
	// `NONE` : Ignore special devices (recommended).
	//
	// `PRESERVE` : Preserve character and block device metadata. This option isn't currently supported for Amazon EFS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-preservedevices
	//
	PreserveDevices *string `field:"optional" json:"preserveDevices" yaml:"preserveDevices"`
	// A value that determines which components of the SMB security descriptor are copied from source to destination objects.
	//
	// This value is only used for transfers between SMB and Amazon FSx for Windows File Server locations, or between two Amazon FSx for Windows File Server locations. For more information about how DataSync handles metadata, see [How DataSync Handles Metadata and Special Files](https://docs.aws.amazon.com/datasync/latest/userguide/special-files.html) .
	//
	// Default value: `OWNER_DACL`
	//
	// `OWNER_DACL` : For each copied object, DataSync copies the following metadata:
	//
	// - Object owner.
	// - NTFS discretionary access control lists (DACLs), which determine whether to grant access to an object.
	//
	// When you use option, DataSync does NOT copy the NTFS system access control lists (SACLs), which are used by administrators to log attempts to access a secured object.
	//
	// `OWNER_DACL_SACL` : For each copied object, DataSync copies the following metadata:
	//
	// - Object owner.
	// - NTFS discretionary access control lists (DACLs), which determine whether to grant access to an object.
	// - NTFS system access control lists (SACLs), which are used by administrators to log attempts to access a secured object.
	//
	// Copying SACLs requires granting additional permissions to the Windows user that DataSync uses to access your SMB location. For information about choosing a user that ensures sufficient permissions to files, folders, and metadata, see [user](https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#SMBuser) .
	//
	// `NONE` : None of the SMB security descriptor components are copied. Destination objects are owned by the user that was provided for accessing the destination location. DACLs and SACLs are set based on the destination server’s configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-securitydescriptorcopyflags
	//
	SecurityDescriptorCopyFlags *string `field:"optional" json:"securityDescriptorCopyFlags" yaml:"securityDescriptorCopyFlags"`
	// Specifies whether your transfer tasks should be put into a queue during certain scenarios when [running multiple tasks](https://docs.aws.amazon.com/datasync/latest/userguide/run-task.html#running-multiple-tasks) . This is `ENABLED` by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-taskqueueing
	//
	TaskQueueing *string `field:"optional" json:"taskQueueing" yaml:"taskQueueing"`
	// A value that determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing it to the destination location.
	//
	// `CHANGED` : DataSync copies only data or metadata that is new or different from the source location to the destination location.
	//
	// `ALL` : DataSync copies all source location content to the destination, without comparing it to existing content on the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-transfermode
	//
	TransferMode *string `field:"optional" json:"transferMode" yaml:"transferMode"`
	// The user ID (UID) of the file's owner.
	//
	// Default value: `INT_VALUE`
	//
	// `INT_VALUE` : Preserve the integer value of the UID and group ID (GID) (recommended).
	//
	// `NAME` : Currently not supported
	//
	// `NONE` : Ignore the UID and GID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-uid
	//
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
	// A value that determines whether a data integrity verification is performed at the end of a task execution after all data and metadata have been transferred.
	//
	// For more information, see [Configure task settings](https://docs.aws.amazon.com/datasync/latest/userguide/create-task.html) .
	//
	// Default value: `POINT_IN_TIME_CONSISTENT`
	//
	// `ONLY_FILES_TRANSFERRED` (recommended): Perform verification only on files that were transferred.
	//
	// `POINT_IN_TIME_CONSISTENT` : Scan the entire source and entire destination at the end of the transfer to verify that the source and destination are fully synchronized. This option isn't supported when transferring to S3 Glacier or S3 Glacier Deep Archive storage classes.
	//
	// `NONE` : No additional verification is done at the end of the transfer, but all data transmissions are integrity-checked with checksum verification during the transfer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-options.html#cfn-datasync-task-options-verifymode
	//
	VerifyMode *string `field:"optional" json:"verifyMode" yaml:"verifyMode"`
}

