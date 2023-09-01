package awsec2


// Properties to create NetworkAcl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   networkAclProps := &NetworkAclProps{
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	NetworkAclName: jsii.String("networkAclName"),
//   	SubnetSelection: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type NetworkAclProps struct {
	// The VPC in which to create the NetworkACL.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The name of the NetworkAcl.
	//
	// Since the NetworkAcl resource doesn't support providing a physical name, the value provided here will be recorded in the `Name` tag.
	// Default: CDK generated name.
	//
	NetworkAclName *string `field:"optional" json:"networkAclName" yaml:"networkAclName"`
	// Subnets in the given VPC to associate the ACL with.
	//
	// More subnets can always be added later by calling
	// `associateWithSubnets()`.
	// Default: - No subnets associated.
	//
	SubnetSelection *SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
}

