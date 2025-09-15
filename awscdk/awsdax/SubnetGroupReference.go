package awsdax


// A reference to a SubnetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetGroupReference := &SubnetGroupReference{
//   	SubnetGroupId: jsii.String("subnetGroupId"),
//   }
//
type SubnetGroupReference struct {
	// The Id of the SubnetGroup resource.
	SubnetGroupId *string `field:"required" json:"subnetGroupId" yaml:"subnetGroupId"`
}

