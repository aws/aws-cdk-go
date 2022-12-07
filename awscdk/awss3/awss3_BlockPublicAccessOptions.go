package awss3


// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &bucketProps{
//   	blockPublicAccess: s3.NewBlockPublicAccess(&blockPublicAccessOptions{
//   		blockPublicPolicy: jsii.Boolean(true),
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

