package awss3


// Container for the transition rule that describes when noncurrent objects transition to the `STANDARD_IA` , `ONEZONE_IA` , `INTELLIGENT_TIERING` , `GLACIER_IR` , `GLACIER` , or `DEEP_ARCHIVE` storage class.
//
// If your bucket is versioning-enabled (or versioning is suspended), you can set this action to request that Amazon S3 transition noncurrent object versions to the `STANDARD_IA` , `ONEZONE_IA` , `INTELLIGENT_TIERING` , `GLACIER_IR` , `GLACIER` , or `DEEP_ARCHIVE` storage class at a specific period in the object's lifetime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   noncurrentVersionTransitionProperty := &noncurrentVersionTransitionProperty{
//   	storageClass: jsii.String("storageClass"),
//   	transitionInDays: jsii.Number(123),
//
//   	// the properties below are optional
//   	newerNoncurrentVersions: jsii.Number(123),
//   }
//
type CfnBucket_NoncurrentVersionTransitionProperty struct {
	// The class of storage used to store the object.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	//
	// For information about the noncurrent days calculations, see [How Amazon S3 Calculates How Long an Object Has Been Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the *Amazon S3 User Guide* .
	TransitionInDays *float64 `field:"required" json:"transitionInDays" yaml:"transitionInDays"`
	// Specifies how many noncurrent versions Amazon S3 will retain.
	//
	// If there are this many more recent noncurrent versions, Amazon S3 will take the associated action. For more information about noncurrent versions, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html) in the *Amazon S3 User Guide* .
	NewerNoncurrentVersions *float64 `field:"optional" json:"newerNoncurrentVersions" yaml:"newerNoncurrentVersions"`
}

