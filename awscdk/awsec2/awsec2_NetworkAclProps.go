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
//   networkAclProps := &networkAclProps{
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	networkAclName: jsii.String("networkAclName"),
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type NetworkAclProps struct {
	// The VPC in which to create the NetworkACL.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The name of the NetworkAcl.
	//
	// It is not recommended to use an explicit name.
	NetworkAclName *string `field:"optional" json:"networkAclName" yaml:"networkAclName"`
	// Subnets in the given VPC to associate the ACL with.
	//
	// More subnets can always be added later by calling
	// `associateWithSubnets()`.
	SubnetSelection *SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
}

