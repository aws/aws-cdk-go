package awsdevicefarm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInstanceProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceProfileProps := &CfnInstanceProfileProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExcludeAppPackagesFromCleanup: []*string{
//   		jsii.String("excludeAppPackagesFromCleanup"),
//   	},
//   	PackageCleanup: jsii.Boolean(false),
//   	RebootAfterUse: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-instanceprofile.html
//
type CfnInstanceProfileProps struct {
	// The name of the instance profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-instanceprofile.html#cfn-devicefarm-instanceprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the instance profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-instanceprofile.html#cfn-devicefarm-instanceprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of strings containing the list of app packages that should not be cleaned up from the device after a test run completes.
	//
	// The list of packages is considered only if you set `packageCleanup` to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-instanceprofile.html#cfn-devicefarm-instanceprofile-excludeapppackagesfromcleanup
	//
	ExcludeAppPackagesFromCleanup *[]*string `field:"optional" json:"excludeAppPackagesFromCleanup" yaml:"excludeAppPackagesFromCleanup"`
	// When set to `true` , Device Farm removes app packages after a test run.
	//
	// The default value is `false` for private devices.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-instanceprofile.html#cfn-devicefarm-instanceprofile-packagecleanup
	//
	PackageCleanup interface{} `field:"optional" json:"packageCleanup" yaml:"packageCleanup"`
	// When set to `true` , Device Farm reboots the instance after a test run.
	//
	// The default value is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-instanceprofile.html#cfn-devicefarm-instanceprofile-rebootafteruse
	//
	RebootAfterUse interface{} `field:"optional" json:"rebootAfterUse" yaml:"rebootAfterUse"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-instanceprofile.html#cfn-devicefarm-instanceprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

