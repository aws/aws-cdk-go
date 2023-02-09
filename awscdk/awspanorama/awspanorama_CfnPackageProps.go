package awspanorama

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPackage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackageProps := &cfnPackageProps{
//   	packageName: jsii.String("packageName"),
//
//   	// the properties below are optional
//   	storageLocation: &storageLocationProperty{
//   		binaryPrefixLocation: jsii.String("binaryPrefixLocation"),
//   		bucket: jsii.String("bucket"),
//   		generatedPrefixLocation: jsii.String("generatedPrefixLocation"),
//   		manifestPrefixLocation: jsii.String("manifestPrefixLocation"),
//   		repoPrefixLocation: jsii.String("repoPrefixLocation"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

