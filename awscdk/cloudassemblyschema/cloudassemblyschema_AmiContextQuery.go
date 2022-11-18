package cloudassemblyschema


// Query to AMI context provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amiContextQuery := &amiContextQuery{
//   	account: jsii.String("account"),
//   	filters: map[string][]*string{
//   		"filtersKey": []*string{
//   			jsii.String("filters"),
//   		},
//   	},
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   	owners: []*string{
//   		jsii.String("owners"),
//   	},
//   }
//
type AmiContextQuery struct {
	// Account to query.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Filters to DescribeImages call.
	Filters *map[string]*[]*string `field:"required" json:"filters" yaml:"filters"`
	// Region to query.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Owners to DescribeImages call.
	Owners *[]*string `field:"optional" json:"owners" yaml:"owners"`
}

