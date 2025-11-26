package previewawss3mixins


// Specifies when an object transitions to a specified storage class.
//
// For more information about Amazon S3 lifecycle configuration rules, see [Transitioning Objects Using Amazon S3 Lifecycle](https://docs.aws.amazon.com/AmazonS3/latest/dev/lifecycle-transition-general-considerations.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transitionProperty := &TransitionProperty{
//   	StorageClass: jsii.String("storageClass"),
//   	TransitionDate: NewDate(),
//   	TransitionInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-transition.html
//
type CfnBucketPropsMixin_TransitionProperty struct {
	// The storage class to which you want the object to transition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-transition.html#cfn-s3-bucket-transition-storageclass
	//
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Indicates when objects are transitioned to the specified storage class.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-transition.html#cfn-s3-bucket-transition-transitiondate
	//
	TransitionDate interface{} `field:"optional" json:"transitionDate" yaml:"transitionDate"`
	// Indicates the number of days after creation when objects are transitioned to the specified storage class.
	//
	// If the specified storage class is `INTELLIGENT_TIERING` , `GLACIER_IR` , `GLACIER` , or `DEEP_ARCHIVE` , valid values are `0` or positive integers. If the specified storage class is `STANDARD_IA` or `ONEZONE_IA` , valid values are positive integers greater than `30` . Be aware that some storage classes have a minimum storage duration and that you're charged for transitioning objects before their minimum storage duration. For more information, see [Constraints and considerations for transitions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-transition-general-considerations.html#lifecycle-configuration-constraints) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-transition.html#cfn-s3-bucket-transition-transitionindays
	//
	TransitionInDays *float64 `field:"optional" json:"transitionInDays" yaml:"transitionInDays"`
}

