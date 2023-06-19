package awsec2


// Options when executing a file.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instance instance
//
//
//   asset := awscdk.NewAsset(this, jsii.String("Asset"), &AssetProps{
//   	Path: jsii.String("./configure.sh"),
//   })
//
//   localPath := instance.UserData.AddS3DownloadCommand(&S3DownloadOptions{
//   	Bucket: asset.Bucket,
//   	BucketKey: asset.S3ObjectKey,
//   	Region: jsii.String("us-east-1"),
//   })
//   instance.UserData.AddExecuteFileCommand(&ExecuteFileOptions{
//   	FilePath: localPath,
//   	Arguments: jsii.String("--verbose -y"),
//   })
//   asset.GrantRead(instance.Role)
//
// Experimental.
type ExecuteFileOptions struct {
	// The path to the file.
	// Experimental.
	FilePath *string `field:"required" json:"filePath" yaml:"filePath"`
	// The arguments to be passed to the file.
	// Experimental.
	Arguments *string `field:"optional" json:"arguments" yaml:"arguments"`
}

