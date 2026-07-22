package awsrds


// Properties for CfnReservedDBInstancePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnReservedDBInstanceMixinProps := &CfnReservedDBInstanceMixinProps{
//   	DbInstanceCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-reserveddbinstance.html
//
type CfnReservedDBInstanceMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-reserveddbinstance.html#cfn-rds-reserveddbinstance-dbinstancecount
	//
	DbInstanceCount *float64 `field:"optional" json:"dbInstanceCount" yaml:"dbInstanceCount"`
}

