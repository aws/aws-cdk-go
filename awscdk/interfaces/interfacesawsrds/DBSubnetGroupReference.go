package interfacesawsrds


// A reference to a DBSubnetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBSubnetGroupReference := &DBSubnetGroupReference{
//   	DbSubnetGroupArn: jsii.String("dbSubnetGroupArn"),
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   }
//
type DBSubnetGroupReference struct {
	// The ARN of the DBSubnetGroup resource.
	DbSubnetGroupArn *string `field:"required" json:"dbSubnetGroupArn" yaml:"dbSubnetGroupArn"`
	// The DBSubnetGroupName of the DBSubnetGroup resource.
	DbSubnetGroupName *string `field:"required" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
}

