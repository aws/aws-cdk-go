package previewawsiotmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSoftwarePackageVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSoftwarePackageVersionMixinProps := &CfnSoftwarePackageVersionMixinProps{
//   	Artifact: &PackageVersionArtifactProperty{
//   		S3Location: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	Description: jsii.String("description"),
//   	PackageName: jsii.String("packageName"),
//   	Recipe: jsii.String("recipe"),
//   	Sbom: &SbomProperty{
//   		S3Location: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VersionName: jsii.String("versionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html
//
type CfnSoftwarePackageVersionMixinProps struct {
	// The artifact location of the package version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-artifact
	//
	Artifact interface{} `field:"optional" json:"artifact" yaml:"artifact"`
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
	// The name of the associated software package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-packagename
	//
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// The inline json job document associated with a software package version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-recipe
	//
	Recipe *string `field:"optional" json:"recipe" yaml:"recipe"`
	// The sbom zip archive location of the package version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-sbom
	//
	Sbom interface{} `field:"optional" json:"sbom" yaml:"sbom"`
	// Metadata that can be used to manage the package version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the new package version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-softwarepackageversion.html#cfn-iot-softwarepackageversion-versionname
	//
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
}

