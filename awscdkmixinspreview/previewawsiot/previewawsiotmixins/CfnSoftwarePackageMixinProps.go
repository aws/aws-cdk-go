package previewawsiotmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSoftwarePackagePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSoftwarePackageMixinProps := &CfnSoftwarePackageMixinProps{
//   	Description: jsii.String("description"),
//   	PackageName: jsii.String("packageName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackage.html
//
type CfnSoftwarePackageMixinProps struct {
	// A summary of the package being created.
	//
	// This can be used to outline the package's contents or purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackage.html#cfn-iot-softwarepackage-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the new software package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackage.html#cfn-iot-softwarepackage-packagename
	//
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Metadata that can be used to manage the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackage.html#cfn-iot-softwarepackage-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

