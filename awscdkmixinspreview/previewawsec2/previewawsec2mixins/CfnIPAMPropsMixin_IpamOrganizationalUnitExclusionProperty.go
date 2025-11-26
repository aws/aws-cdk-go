package previewawsec2mixins


// If your IPAM is integrated with AWS Organizations, you can exclude an [organizational unit (OU)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#organizationalunit) from being managed by IPAM. When you exclude an OU, IPAM will not manage the IP addresses in accounts in that OU. For more information, see [Exclude organizational units from IPAM](https://docs.aws.amazon.com/vpc/latest/ipam/exclude-ous.html) in the *Amazon Virtual Private Cloud IP Address Manager User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ipamOrganizationalUnitExclusionProperty := &IpamOrganizationalUnitExclusionProperty{
//   	OrganizationsEntityPath: jsii.String("organizationsEntityPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipam-ipamorganizationalunitexclusion.html
//
type CfnIPAMPropsMixin_IpamOrganizationalUnitExclusionProperty struct {
	// An AWS Organizations entity path.
	//
	// For more information on the entity path, see [Understand the AWS Organizations entity path](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_last-accessed-view-data-orgs.html#access_policies_access-advisor-viewing-orgs-entity-path) in the *AWS Identity and Access Management User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipam-ipamorganizationalunitexclusion.html#cfn-ec2-ipam-ipamorganizationalunitexclusion-organizationsentitypath
	//
	OrganizationsEntityPath *string `field:"optional" json:"organizationsEntityPath" yaml:"organizationsEntityPath"`
}

