package awsec2


// If your IPAM is integrated with AWS Organizations and you add an organizational unit (OU) exclusion, IPAM will not manage the IP addresses in accounts in that OU exclusion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipamResourceDiscoveryOrganizationalUnitExclusionProperty := &IpamResourceDiscoveryOrganizationalUnitExclusionProperty{
//   	OrganizationsEntityPath: jsii.String("organizationsEntityPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamresourcediscovery-ipamresourcediscoveryorganizationalunitexclusion.html
//
type CfnIPAMResourceDiscovery_IpamResourceDiscoveryOrganizationalUnitExclusionProperty struct {
	// An AWS Organizations entity path.
	//
	// Build the path for the OU(s) using AWS Organizations IDs separated by a '/'. Include all child OUs by ending the path with '/*'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamresourcediscovery-ipamresourcediscoveryorganizationalunitexclusion.html#cfn-ec2-ipamresourcediscovery-ipamresourcediscoveryorganizationalunitexclusion-organizationsentitypath
	//
	OrganizationsEntityPath *string `field:"required" json:"organizationsEntityPath" yaml:"organizationsEntityPath"`
}

