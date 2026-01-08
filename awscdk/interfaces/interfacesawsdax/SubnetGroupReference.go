package interfacesawsdax


// A reference to a SubnetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetGroupReference := &SubnetGroupReference{
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   }
//
type SubnetGroupReference struct {
	// The SubnetGroupName of the SubnetGroup resource.
	SubnetGroupName *string `field:"required" json:"subnetGroupName" yaml:"subnetGroupName"`
}

