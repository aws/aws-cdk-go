package awsdocdb


// A reference to a DBInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBInstanceReference := &DBInstanceReference{
//   	DbInstanceId: jsii.String("dbInstanceId"),
//   }
//
type DBInstanceReference struct {
	// The Id of the DBInstance resource.
	DbInstanceId *string `field:"required" json:"dbInstanceId" yaml:"dbInstanceId"`
}

