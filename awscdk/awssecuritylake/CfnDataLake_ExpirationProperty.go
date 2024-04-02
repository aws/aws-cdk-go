package awssecuritylake


// Provides data expiration details of Amazon Security Lake object.
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
	// Number of days before data expires in the Amazon Security Lake object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-expiration.html#cfn-securitylake-datalake-expiration-days
	//
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

