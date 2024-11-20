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
//   	ServerSdkVersion: jsii.String("serverSdkVersion"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html
//
type CfnBuildProps struct {
	// A descriptive label that is associated with a build.
	//
	// Build names do not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The environment that your game server binaries run on.
	//
	// This value determines the type of fleet resources that you use for this build. If your game build contains multiple executables, they all must run on the same operating system. This parameter is required, and there's no default value. You can't change a build's operating system later.
	//
	// > Amazon Linux 2 (AL2) will reach end of support on 6/30/2025. See more details in the [Amazon Linux 2 FAQs](https://docs.aws.amazon.com/https://aws.amazon.com/amazon-linux-2/faqs/) . For game servers that are hosted on AL2 and use Amazon GameLift server SDK 4.x., first update the game server build to server SDK 5.x, and then deploy to AL2023 instances. See [Migrate to Amazon GameLift server SDK version 5.](https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk5-migration.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-operatingsystem
	//
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// A server SDK version you used when integrating your game server build with Amazon GameLift.
	//
	// For more information see [Integrate games with custom game servers](https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-custom-intro.html) . By default Amazon GameLift sets this value to `4.0.2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-serversdkversion
	//
	ServerSdkVersion *string `field:"optional" json:"serverSdkVersion" yaml:"serverSdkVersion"`
	// Information indicating where your game build files are stored.
	//
	// Use this parameter only when creating a build with files stored in an Amazon S3 bucket that you own. The storage location must specify an Amazon S3 bucket name and key. The location must also specify a role ARN that you set up to allow Amazon GameLift to access your Amazon S3 bucket. The S3 bucket and your new build must be in the same Region.
	//
	// If a `StorageLocation` is specified, the size of your file can be found in your Amazon S3 bucket. Amazon GameLift will report a `SizeOnDisk` of 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-storagelocation
	//
	StorageLocation interface{} `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Version information that is associated with this build.
	//
	// Version strings do not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

