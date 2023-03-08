package awscodebuild


// Information about a file system created by Amazon Elastic File System (EFS).
//
// For more information, see [What Is Amazon Elastic File System?](https://docs.aws.amazon.com/efs/latest/ug/whatisefs.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectFileSystemLocationProperty := &projectFileSystemLocationProperty{
//   	identifier: jsii.String("identifier"),
//   	location: jsii.String("location"),
//   	mountPoint: jsii.String("mountPoint"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	mountOptions: jsii.String("mountOptions"),
//   }
//
type CfnProject_ProjectFileSystemLocationProperty struct {
	// The name used to access a file system created by Amazon EFS.
	//
	// CodeBuild creates an environment variable by appending the `identifier` in all capital letters to `CODEBUILD_` . For example, if you specify `my_efs` for `identifier` , a new environment variable is create named `CODEBUILD_MY_EFS` .
	//
	// The `identifier` is used to mount your file system.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// A string that specifies the location of the file system created by Amazon EFS.
	//
	// Its format is `efs-dns-name:/directory-path` . You can find the DNS name of file system when you view it in the Amazon EFS console. The directory path is a path to a directory in the file system that CodeBuild mounts. For example, if the DNS name of a file system is `fs-abcd1234.efs.us-west-2.amazonaws.com` , and its mount directory is `my-efs-mount-directory` , then the `location` is `fs-abcd1234.efs.us-west-2.amazonaws.com:/my-efs-mount-directory` .
	//
	// The directory path in the format `efs-dns-name:/directory-path` is optional. If you do not specify a directory path, the location is only the DNS name and CodeBuild mounts the entire file system.
	Location *string `field:"required" json:"location" yaml:"location"`
	// The location in the container where you mount the file system.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
	// The type of the file system.
	//
	// The one supported type is `EFS` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// The mount options for a file system created by Amazon EFS.
	//
	// The default mount options used by CodeBuild are `nfsvers=4.1,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2` . For more information, see [Recommended NFS Mount Options](https://docs.aws.amazon.com/efs/latest/ug/mounting-fs-nfs-mount-settings.html) .
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

