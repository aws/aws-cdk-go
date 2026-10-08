package awsredshift


// The S3 Access Grants scope.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   readWriteAccessProperty := &ReadWriteAccessProperty{
//   	Authorization: jsii.String("authorization"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-readwriteaccess.html
//
type CfnRedshiftIdcApplication_ReadWriteAccessProperty struct {
	// Determines whether the read/write scope is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-readwriteaccess.html#cfn-redshift-redshiftidcapplication-readwriteaccess-authorization
	//
	Authorization *string `field:"required" json:"authorization" yaml:"authorization"`
}

