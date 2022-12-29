package awsgamelift


// Properties for defining a `CfnBuild`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBuildProps := &cfnBuildProps{
//   	name: jsii.String("name"),
//   	operatingSystem: jsii.String("operatingSystem"),
//   	storageLocation: &storageLocationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   	version: jsii.String("version"),
//   }
//
type CfnBuildProps struct {
	// A descriptive label that is associated with a build.
	//
	// Build names do not need to be unique.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The operating system that the game server binaries are built to run on.
	//
	// This value determines the type of fleet resources that you can use for this build. If your game build contains multiple executables, they all must run on the same operating system. If an operating system is not specified when creating a build, Amazon GameLift uses the default value (WINDOWS_2012). This value cannot be changed later.
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// Information indicating where your game build files are stored.
	//
	// Use this parameter only when creating a build with files stored in an Amazon S3 bucket that you own. The storage location must specify an Amazon S3 bucket name and key. The location must also specify a role ARN that you set up to allow Amazon GameLift to access your Amazon S3 bucket. The S3 bucket and your new build must be in the same Region.
	//
	// If a `StorageLocation` is specified, the size of your file can be found in your Amazon S3 bucket. Amazon GameLift will report a `SizeOnDisk` of 0.
	StorageLocation interface{} `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Version information that is associated with this build.
	//
	// Version strings do not need to be unique.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

