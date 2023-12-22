package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSoftwarePackage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSoftwarePackageProps := &CfnSoftwarePackageProps{
//   	Description: jsii.String("description"),
//   	PackageName: jsii.String("packageName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackage.html
//
type CfnSoftwarePackageProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackage.html#cfn-iot-softwarepackage-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackage.html#cfn-iot-softwarepackage-packagename
	//
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackage.html#cfn-iot-softwarepackage-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

