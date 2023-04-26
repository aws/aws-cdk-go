package awsgamelift


// Properties for defining a `CfnBuild`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBuildProps := &CfnBuildProps{
//   	Name: jsii.String("name"),
//   	OperatingSystem: jsii.String("operatingSystem"),
//   	StorageLocation: &StorageLocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   	Version: jsii.String("version"),
//   }
//
type CfnBuildProps struct {
	// A descriptive label that is associated with a build.
	//
	// Build names do not need to be unique.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The operating system that your game server binaries run on.
	//
	// This value determines the type of fleet resources that you use for this build. If your game build contains multiple executables, they all must run on the same operating system. You must specify a valid operating system in this request. There is no default value. You can't change a build's operating system later.
	//
	// > If you have active fleets using the Windows Server 2012 operating system, you can continue to create new builds using this OS until October 10, 2023, when Microsoft ends its support. All others must use Windows Server 2016 when creating new Windows-based builds.
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

