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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-packagename
	//
	PackageName *string `field:"required" json:"packageName" yaml:"packageName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-versionname
	//
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
}

