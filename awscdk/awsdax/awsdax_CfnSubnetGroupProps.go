package awsdax


// Properties for defining a `CfnSubnetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubnetGroupProps := &cfnSubnetGroupProps{
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	subnetGroupName: jsii.String("subnetGroupName"),
//   }
//
type CfnSubnetGroupProps struct {
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The description of the subnet group.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the subnet group.
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
}

