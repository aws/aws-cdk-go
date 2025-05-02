package awss3


// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &BucketProps{
//   	BlockPublicAccess: s3.NewBlockPublicAccess(&BlockPublicAccessOptions{
//   		BlockPublicPolicy: jsii.Boolean(true),
//   	}),
//   })
//
type BlockPublicAccessOptions struct {
	// Whether to block public ACLs.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-options
	//
	BlockPublicAcls *bool `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Whether to block public policy.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-options
	//
	BlockPublicPolicy *bool `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Whether to ignore public ACLs.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-options
	//
	IgnorePublicAcls *bool `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Whether to restrict public access.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-options
	//
	RestrictPublicBuckets *bool `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

