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
//   transitionProperty := &transitionProperty{
//   	storageClass: jsii.String("storageClass"),
//
//   	// the properties below are optional
//   	transitionDate: NewDate(),
//   	transitionInDays: jsii.Number(123),
//   }
//
type CfnBucket_TransitionProperty struct {
	// The storage class to which you want the object to transition.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Indicates when objects are transitioned to the specified storage class.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC.
	TransitionDate interface{} `field:"optional" json:"transitionDate" yaml:"transitionDate"`
	// Indicates the number of days after creation when objects are transitioned to the specified storage class.
	//
	// The value must be a positive integer.
	TransitionInDays *float64 `field:"optional" json:"transitionInDays" yaml:"transitionInDays"`
}

