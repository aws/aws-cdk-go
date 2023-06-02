package awspanorama

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPackage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackageProps := &CfnPackageProps{
//   	PackageName: jsii.String("packageName"),
//
//   	// the properties below are optional
//   	StorageLocation: &StorageLocationProperty{
//   		BinaryPrefixLocation: jsii.String("binaryPrefixLocation"),
//   		Bucket: jsii.String("bucket"),
//   		GeneratedPrefixLocation: jsii.String("generatedPrefixLocation"),
//   		ManifestPrefixLocation: jsii.String("manifestPrefixLocation"),
//   		RepoPrefixLocation: jsii.String("repoPrefixLocation"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPackageProps struct {
	// A name for the package.
	PackageName *string `field:"required" json:"packageName" yaml:"packageName"`
	// `AWS::Panorama::Package.StorageLocation`.
	StorageLocation interface{} `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Tags for the package.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

