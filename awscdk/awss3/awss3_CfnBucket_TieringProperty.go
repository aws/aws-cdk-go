package awss3


// The S3 Intelligent-Tiering storage class is designed to optimize storage costs by automatically moving data to the most cost-effective storage access tier, without additional operational overhead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tieringProperty := &tieringProperty{
//   	accessTier: jsii.String("accessTier"),
//   	days: jsii.Number(123),
//   }
//
type CfnBucket_TieringProperty struct {
	// S3 Intelligent-Tiering access tier.
	//
	// See [Storage class for automatically optimizing frequently and infrequently accessed objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access) for a list of access tiers in the S3 Intelligent-Tiering storage class.
	AccessTier *string `field:"required" json:"accessTier" yaml:"accessTier"`
	// The number of consecutive days of no access after which an object will be eligible to be transitioned to the corresponding tier.
	//
	// The minimum number of days specified for Archive Access tier must be at least 90 days and Deep Archive Access tier must be at least 180 days. The maximum can be up to 2 years (730 days).
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

