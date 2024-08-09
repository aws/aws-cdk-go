package awscdkcloudassemblyschema


// Where in S3 a file asset needs to be published.
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

