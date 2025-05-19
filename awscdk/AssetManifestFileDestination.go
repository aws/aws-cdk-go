package awscdk


// The destination for a file asset, when it is given to the AssetManifestBuilder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var assumeRoleAdditionalOptions interface{}
//
//   assetManifestFileDestination := &AssetManifestFileDestination{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	Role: &RoleOptions{
//   		AssumeRoleArn: jsii.String("assumeRoleArn"),
//
//   		// the properties below are optional
//   		AssumeRoleAdditionalOptions: map[string]interface{}{
//   			"assumeRoleAdditionalOptionsKey": assumeRoleAdditionalOptions,
//   		},
//   		AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	},
//   }
//
type AssetManifestFileDestination struct {
	// Bucket name where the file asset should be written.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Prefix to prepend to the asset hash.
	// Default: ''.
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Role to use for uploading.
	// Default: - current role.
	//
	Role *RoleOptions `field:"optional" json:"role" yaml:"role"`
}

