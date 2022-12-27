package awss3


// Places an Object Lock configuration on the specified bucket.
//
// The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectLockConfigurationProperty := &objectLockConfigurationProperty{
//   	objectLockEnabled: jsii.String("objectLockEnabled"),
//   	rule: &objectLockRuleProperty{
//   		defaultRetention: &defaultRetentionProperty{
//   			days: jsii.Number(123),
//   			mode: jsii.String("mode"),
//   			years: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnBucket_ObjectLockConfigurationProperty struct {
	// Indicates whether this bucket has an Object Lock configuration enabled.
	//
	// Enable `ObjectLockEnabled` when you apply `ObjectLockConfiguration` to a bucket.
	ObjectLockEnabled *string `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// Specifies the Object Lock rule for the specified object.
	//
	// Enable the this rule when you apply `ObjectLockConfiguration` to a bucket. If Object Lock is turned on, bucket settings require both `Mode` and a period of either `Days` or `Years` . You cannot specify `Days` and `Years` at the same time. For more information, see [ObjectLockRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-objectlockrule.html) and [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html) .
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

