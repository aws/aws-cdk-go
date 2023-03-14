package awsses


// Properties for defining a `CfnReceiptFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReceiptFilterProps := &cfnReceiptFilterProps{
//   	filter: &filterProperty{
//   		ipFilter: &ipFilterProperty{
//   			cidr: jsii.String("cidr"),
//   			policy: jsii.String("policy"),
//   		},
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnReceiptFilterProps struct {
	// A data structure that describes the IP address filter to create, which consists of a name, an IP address range, and whether to allow or block mail from it.
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
}

