package awss3


// Specifies the Object Lock rule for the specified object.
//
// Enable the this rule when you apply `ObjectLockConfiguration` to a bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectLockRuleProperty := &ObjectLockRuleProperty{
//   	DefaultRetention: &DefaultRetentionProperty{
//   		Days: jsii.Number(123),
//   		Mode: jsii.String("mode"),
//   		Years: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-objectlockrule.html
//
type CfnBucket_ObjectLockRuleProperty struct {
	// The default Object Lock retention mode and period that you want to apply to new objects placed in the specified bucket.
	//
	// If Object Lock is turned on, bucket settings require both `Mode` and a period of either `Days` or `Years` . You cannot specify `Days` and `Years` at the same time. For more information about allowable values for mode and period, see [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-objectlockrule.html#cfn-s3-bucket-objectlockrule-defaultretention
	//
	DefaultRetention interface{} `field:"optional" json:"defaultRetention" yaml:"defaultRetention"`
}

