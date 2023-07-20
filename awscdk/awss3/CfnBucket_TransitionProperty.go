package awss3


// Specifies when an object transitions to a specified storage class.
//
// For more information about Amazon S3 lifecycle configuration rules, see [Transitioning Objects Using Amazon S3 Lifecycle](https://docs.aws.amazon.com/AmazonS3/latest/dev/lifecycle-transition-general-considerations.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitionProperty := &TransitionProperty{
//   	StorageClass: jsii.String("storageClass"),
//
//   	// the properties below are optional
//   	TransitionDate: NewDate(),
//   	TransitionInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-transition.html
//
type CfnBucket_TransitionProperty struct {
	// The storage class to which you want the object to transition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-transition.html#cfn-s3-bucket-transition-storageclass
	//
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Indicates when objects are transitioned to the specified storage class.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-transition.html#cfn-s3-bucket-transition-transitiondate
	//
	TransitionDate interface{} `field:"optional" json:"transitionDate" yaml:"transitionDate"`
	// Indicates the number of days after creation when objects are transitioned to the specified storage class.
	//
	// The value must be a positive integer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-transition.html#cfn-s3-bucket-transition-transitionindays
	//
	TransitionInDays *float64 `field:"optional" json:"transitionInDays" yaml:"transitionInDays"`
}

