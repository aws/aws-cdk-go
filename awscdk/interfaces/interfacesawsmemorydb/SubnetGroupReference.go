package interfacesawsmemorydb


// A reference to a SubnetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetGroupReference := &SubnetGroupReference{
//   	SubnetGroupArn: jsii.String("subnetGroupArn"),
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   }
//
type SubnetGroupReference struct {
	// The ARN of the SubnetGroup resource.
	SubnetGroupArn *string `field:"required" json:"subnetGroupArn" yaml:"subnetGroupArn"`
	// The SubnetGroupName of the SubnetGroup resource.
	SubnetGroupName *string `field:"required" json:"subnetGroupName" yaml:"subnetGroupName"`
}

