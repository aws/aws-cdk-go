package awss3


// Specifies tags to use to identify a subset of objects for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagFilterProperty := &tagFilterProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnBucket_TagFilterProperty struct {
	// The tag key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

