package awss3


// Specifies when noncurrent object versions expire.
//
// Upon expiration, Amazon S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's lifetime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   noncurrentVersionExpirationProperty := &noncurrentVersionExpirationProperty{
//   	noncurrentDays: jsii.Number(123),
//
//   	// the properties below are optional
//   	newerNoncurrentVersions: jsii.Number(123),
//   }
//
type CfnBucket_NoncurrentVersionExpirationProperty struct {
	// Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	//
	// For information about the noncurrent days calculations, see [How Amazon S3 Calculates When an Object Became Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the *Amazon S3 User Guide* .
	NoncurrentDays *float64 `field:"required" json:"noncurrentDays" yaml:"noncurrentDays"`
	// Specifies how many noncurrent versions Amazon S3 will retain.
	//
	// If there are this many more recent noncurrent versions, Amazon S3 will take the associated action. For more information about noncurrent versions, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html) in the *Amazon S3 User Guide* .
	NewerNoncurrentVersions *float64 `field:"optional" json:"newerNoncurrentVersions" yaml:"newerNoncurrentVersions"`
}

