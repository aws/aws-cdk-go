package awss3


// The journal table record expiration settings for a journal table in an S3 Metadata configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordExpirationProperty := &RecordExpirationProperty{
//   	Expiration: jsii.String("expiration"),
//
//   	// the properties below are optional
//   	Days: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-recordexpiration.html
//
type CfnBucket_RecordExpirationProperty struct {
	// Specifies whether journal table record expiration is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-recordexpiration.html#cfn-s3-bucket-recordexpiration-expiration
	//
	Expiration *string `field:"required" json:"expiration" yaml:"expiration"`
	// If you enable journal table record expiration, you can set the number of days to retain your journal table records.
	//
	// Journal table records must be retained for a minimum of 7 days. To set this value, specify any whole number from `7` to `2147483647` . For example, to retain your journal table records for one year, set this value to `365` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-recordexpiration.html#cfn-s3-bucket-recordexpiration-days
	//
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

