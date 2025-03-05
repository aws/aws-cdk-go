package awss3


// The container element for optionally specifying the default Object Lock retention settings for new objects placed in the specified bucket.
//
// > - The `DefaultRetention` settings require both a mode and a period.
// > - The `DefaultRetention` period can be either `Days` or `Years` but you must select one. You cannot specify `Days` and `Years` at the same time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultRetentionProperty := &DefaultRetentionProperty{
//   	Days: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   	Years: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html
//
type CfnBucket_DefaultRetentionProperty struct {
	// The number of days that you want to specify for the default retention period.
	//
	// If Object Lock is turned on, you must specify `Mode` and specify either `Days` or `Years` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html#cfn-s3-bucket-defaultretention-days
	//
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// The default Object Lock retention mode you want to apply to new objects placed in the specified bucket.
	//
	// If Object Lock is turned on, you must specify `Mode` and specify either `Days` or `Years` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html#cfn-s3-bucket-defaultretention-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The number of years that you want to specify for the default retention period.
	//
	// If Object Lock is turned on, you must specify `Mode` and specify either `Days` or `Years` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html#cfn-s3-bucket-defaultretention-years
	//
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

