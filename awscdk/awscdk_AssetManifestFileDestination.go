// Version 2 of the AWS Cloud Development Kit library
package awscdk


// The destination for a file asset, when it is given to the AssetManifestBuilder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   assetManifestFileDestination := &assetManifestFileDestination{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	role: &roleOptions{
//   		assumeRoleArn: jsii.String("assumeRoleArn"),
//
//   		// the properties below are optional
//   		assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	},
//   }
//
type AssetManifestFileDestination struct {
	// Bucket name where the file asset should be written.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Prefix to prepend to the asset hash.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Role to use for uploading.
	Role *RoleOptions `field:"optional" json:"role" yaml:"role"`
}

