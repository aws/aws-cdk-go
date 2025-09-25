package awsdocdb


// A reference to a DBSubnetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBSubnetGroupReference := &DBSubnetGroupReference{
//   	DbSubnetGroupId: jsii.String("dbSubnetGroupId"),
//   }
//
type DBSubnetGroupReference struct {
	// The Id of the DBSubnetGroup resource.
	DbSubnetGroupId *string `field:"required" json:"dbSubnetGroupId" yaml:"dbSubnetGroupId"`
}

