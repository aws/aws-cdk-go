package awssecuritylake


// Provides data expiration details of the Amazon Security Lake object.
//
// You can specify your preferred Amazon S3 storage class and the time period for S3 objects to stay in that storage class before they expire. For more information about Amazon S3 Lifecycle configurations, see [Managing your storage lifecycle](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html) in the *Amazon Simple Storage Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expirationProperty := &ExpirationProperty{
//   	Days: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-expiration.html
//
type CfnDataLake_ExpirationProperty struct {
	// The number of days before data expires in the Amazon Security Lake object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-expiration.html#cfn-securitylake-datalake-expiration-days
	//
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

