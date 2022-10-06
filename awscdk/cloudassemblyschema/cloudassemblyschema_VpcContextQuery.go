package cloudassemblyschema


// Query input for looking up a VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcContextQuery := &vpcContextQuery{
//   	account: jsii.String("account"),
//   	filter: map[string]*string{
//   		"filterKey": jsii.String("filter"),
//   	},
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   	returnAsymmetricSubnets: jsii.Boolean(false),
//   	subnetGroupNameTag: jsii.String("subnetGroupNameTag"),
//   }
//
type VpcContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Filters to apply to the VPC.
	//
	// Filter parameters are the same as passed to DescribeVpcs.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html
	//
	Filter *map[string]*string `field:"required" json:"filter" yaml:"filter"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Whether to populate the subnetGroups field of the {@link VpcContextResponse}, which contains potentially asymmetric subnet groups.
	ReturnAsymmetricSubnets *bool `field:"optional" json:"returnAsymmetricSubnets" yaml:"returnAsymmetricSubnets"`
	// Optional tag for subnet group name.
	//
	// If not provided, we'll look at the aws-cdk:subnet-name tag.
	// If the subnet does not have the specified tag,
	// we'll use its type as the name.
	SubnetGroupNameTag *string `field:"optional" json:"subnetGroupNameTag" yaml:"subnetGroupNameTag"`
}

