package awsrds


// Properties for defining a `CfnReservedDBInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReservedDBInstanceProps := &CfnReservedDBInstanceProps{
//   	DbInstanceCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-reserveddbinstance.html
//
type CfnReservedDBInstanceProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-reserveddbinstance.html#cfn-rds-reserveddbinstance-dbinstancecount
	//
	DbInstanceCount *float64 `field:"optional" json:"dbInstanceCount" yaml:"dbInstanceCount"`
}

