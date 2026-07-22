package interfacesawsrds


// A reference to a ReservedDBInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reservedDBInstanceReference := &ReservedDBInstanceReference{
//   	ReservedDbInstanceArn: jsii.String("reservedDbInstanceArn"),
//   	ReservedDbInstanceId: jsii.String("reservedDbInstanceId"),
//   }
//
type ReservedDBInstanceReference struct {
	// The ARN of the ReservedDBInstance resource.
	ReservedDbInstanceArn *string `field:"required" json:"reservedDbInstanceArn" yaml:"reservedDbInstanceArn"`
	// The ReservedDBInstanceId of the ReservedDBInstance resource.
	ReservedDbInstanceId *string `field:"required" json:"reservedDbInstanceId" yaml:"reservedDbInstanceId"`
}

