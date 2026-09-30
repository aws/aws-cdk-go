package awss3


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaulteventhold.html#cfn-s3-bucket-defaulteventhold-days
	//
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaulteventhold.html#cfn-s3-bucket-defaulteventhold-years
	//
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

