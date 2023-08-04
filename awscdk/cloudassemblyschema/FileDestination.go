package cloudassemblyschema


// Where in S3 a file asset needs to be published.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileDestination := &FileDestination{
//   	BucketName: jsii.String("bucketName"),
//   	ObjectKey: jsii.String("objectKey"),
//
//   	// the properties below are optional
//   	AssumeRoleArn: jsii.String("assumeRoleArn"),
//   	AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	Region: jsii.String("region"),
//   }
//
type FileDestination struct {
	// The role that needs to be assumed while publishing this asset.
	// Default: - No role will be assumed.
	//
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	// Default: - No ExternalId will be supplied.
	//
	AssumeRoleExternalId *string `field:"optional" json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	// Default: - Current region.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The name of the bucket.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The destination object key.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
}

