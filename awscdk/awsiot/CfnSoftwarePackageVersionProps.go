package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSoftwarePackageVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSoftwarePackageVersionProps := &CfnSoftwarePackageVersionProps{
//   	PackageName: jsii.String("packageName"),
//
//   	// the properties below are optional
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VersionName: jsii.String("versionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html
//
type CfnSoftwarePackageVersionProps struct {
	// The name of the associated software package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-packagename
	//
	PackageName *string `field:"required" json:"packageName" yaml:"packageName"`
	// Metadata that can be used to define a package version’s configuration.
	//
	// For example, the S3 file location, configuration options that are being sent to the device or fleet.
	//
	// The combined size of all the attributes on a package version is limited to 3KB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// A summary of the package version being created.
	//
	// This can be used to outline the package's contents or purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata that can be used to manage the package version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the new package version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-versionname
	//
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
}

