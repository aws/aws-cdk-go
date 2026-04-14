package awss3files


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expirationDataRuleProperty := &ExpirationDataRuleProperty{
//   	DaysAfterLastAccess: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-filesystem-expirationdatarule.html
//
type CfnFileSystem_ExpirationDataRuleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-filesystem-expirationdatarule.html#cfn-s3files-filesystem-expirationdatarule-daysafterlastaccess
	//
	DaysAfterLastAccess *float64 `field:"required" json:"daysAfterLastAccess" yaml:"daysAfterLastAccess"`
}

