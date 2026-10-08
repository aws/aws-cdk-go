package awsredshift


// The Lake Formation scope.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lakeFormationQueryProperty := &LakeFormationQueryProperty{
//   	Authorization: jsii.String("authorization"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-lakeformationquery.html
//
type CfnRedshiftIdcApplication_LakeFormationQueryProperty struct {
	// Determines whether the query scope is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-lakeformationquery.html#cfn-redshift-redshiftidcapplication-lakeformationquery-authorization
	//
	Authorization *string `field:"required" json:"authorization" yaml:"authorization"`
}

