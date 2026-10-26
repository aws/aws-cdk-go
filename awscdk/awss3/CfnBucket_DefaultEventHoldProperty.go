package awss3


// The default Object Lock event hold duration that you want to apply to new objects placed in the specified bucket, unless an object's upload request specifies its own Object Lock settings.
//
// New object versions are placed under an event hold and remain protected while the hold is active. When the event hold is released, Amazon S3 retains the object version for this duration, measured from the time the hold is released.
//   The ``DefaultEventHold`` duration can be specified in either ``Days`` or ``Years``, but you must select one. You can't specify ``Days`` and ``Years`` at the same time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultEventHoldProperty := &DefaultEventHoldProperty{
//   	Days: jsii.Number(123),
//   	Years: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaulteventhold.html
//
type CfnBucket_DefaultEventHoldProperty struct {
	// The number of days that you want to specify for the default event hold duration.
	//
	// Valid values are ``1`` to ``36500`` days. You must specify either ``Days`` or ``Years``, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaulteventhold.html#cfn-s3-bucket-defaulteventhold-days
	//
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// The number of years that you want to specify for the default event hold duration.
	//
	// Valid values are ``1`` to ``100`` years. You must specify either ``Days`` or ``Years``, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaulteventhold.html#cfn-s3-bucket-defaulteventhold-years
	//
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

