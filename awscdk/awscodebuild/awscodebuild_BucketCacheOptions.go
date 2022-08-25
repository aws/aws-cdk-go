package awscodebuild


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketCacheOptions := &bucketCacheOptions{
//   	prefix: jsii.String("prefix"),
//   }
//
type BucketCacheOptions struct {
	// The prefix to use to store the cache in the bucket.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

