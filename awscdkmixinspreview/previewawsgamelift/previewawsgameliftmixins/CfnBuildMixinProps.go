package previewawsgameliftmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnBuildPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBuildMixinProps := &CfnBuildMixinProps{
//   	Name: jsii.String("name"),
//   	OperatingSystem: jsii.String("operatingSystem"),
//   	ServerSdkVersion: jsii.String("serverSdkVersion"),
//   	StorageLocation: &StorageLocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		ObjectVersion: jsii.String("objectVersion"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html
//
type CfnBuildMixinProps struct {
	// A descriptive label that is associated with a build.
	//
	// Build names do not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The operating system that your game server binaries run on.
	//
	// This value determines the type of fleet resources that you use for this build. If your game build contains multiple executables, they all must run on the same operating system. You must specify a valid operating system in this request. There is no default value. You can't change a build's operating system later.
	//
	// > Amazon Linux 2 (AL2) will reach end of support on 6/30/2025. See more details in the [Amazon Linux 2 FAQs](https://docs.aws.amazon.com/amazon-linux-2/faqs/) . For game servers that are hosted on AL2 and use server SDK version 4.x for Amazon GameLift Servers, first update the game server build to server SDK 5.x, and then deploy to AL2023 instances. See [Migrate to server SDK version 5.](https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk5-migration.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-operatingsystem
	//
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// A server SDK version you used when integrating your game server build with Amazon GameLift Servers.
	//
	// For more information see [Integrate games with custom game servers](https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-custom-intro.html) . By default Amazon GameLift Servers sets this value to `4.0.2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-serversdkversion
	//
	ServerSdkVersion *string `field:"optional" json:"serverSdkVersion" yaml:"serverSdkVersion"`
	// Information indicating where your game build files are stored.
	//
	// Use this parameter only when creating a build with files stored in an Amazon S3 bucket that you own. The storage location must specify an Amazon S3 bucket name and key. The location must also specify a role ARN that you set up to allow Amazon GameLift Servers to access your Amazon S3 bucket. The S3 bucket and your new build must be in the same Region.
	//
	// If a `StorageLocation` is specified, the size of your file can be found in your Amazon S3 bucket. Amazon GameLift Servers will report a `SizeOnDisk` of 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-storagelocation
	//
	StorageLocation interface{} `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Version information that is associated with this build.
	//
	// Version strings do not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

