package mixinsawsevs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEnvironmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnvironmentMixinProps := &CfnEnvironmentMixinProps{
//   	ConnectivityInfo: &ConnectivityInfoProperty{
//   		PrivateRouteServerPeerings: []*string{
//   			jsii.String("privateRouteServerPeerings"),
//   		},
//   	},
//   	EnvironmentName: jsii.String("environmentName"),
//   	Hosts: []interface{}{
//   		&HostInfoForCreateProperty{
//   			DedicatedHostId: jsii.String("dedicatedHostId"),
//   			HostName: jsii.String("hostName"),
//   			InstanceType: jsii.String("instanceType"),
//   			KeyName: jsii.String("keyName"),
//   			PlacementGroupId: jsii.String("placementGroupId"),
//   		},
//   	},
//   	InitialVlans: &InitialVlansProperty{
//   		EdgeVTep: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		ExpansionVlan1: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		ExpansionVlan2: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		Hcx: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		HcxNetworkAclId: jsii.String("hcxNetworkAclId"),
//   		IsHcxPublic: jsii.Boolean(false),
//   		NsxUpLink: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VmkManagement: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VmManagement: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VMotion: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VSan: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VTep: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LicenseInfo: &LicenseInfoProperty{
//   		SolutionKey: jsii.String("solutionKey"),
//   		VsanKey: jsii.String("vsanKey"),
//   	},
//   	ServiceAccessSecurityGroups: &ServiceAccessSecurityGroupsProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   	},
//   	ServiceAccessSubnetId: jsii.String("serviceAccessSubnetId"),
//   	SiteId: jsii.String("siteId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TermsAccepted: jsii.Boolean(false),
//   	VcfHostnames: &VcfHostnamesProperty{
//   		CloudBuilder: jsii.String("cloudBuilder"),
//   		Nsx: jsii.String("nsx"),
//   		NsxEdge1: jsii.String("nsxEdge1"),
//   		NsxEdge2: jsii.String("nsxEdge2"),
//   		NsxManager1: jsii.String("nsxManager1"),
//   		NsxManager2: jsii.String("nsxManager2"),
//   		NsxManager3: jsii.String("nsxManager3"),
//   		SddcManager: jsii.String("sddcManager"),
//   		VCenter: jsii.String("vCenter"),
//   	},
//   	VcfVersion: jsii.String("vcfVersion"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html
//
type CfnEnvironmentMixinProps struct {
	// The connectivity configuration for the environment.
	//
	// Amazon EVS requires that you specify two route server peer IDs. During environment creation, the route server endpoints peer with the NSX uplink VLAN for connectivity to the NSX overlay network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-connectivityinfo
	//
	ConnectivityInfo interface{} `field:"optional" json:"connectivityInfo" yaml:"connectivityInfo"`
	// The name of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-environmentname
	//
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// Required for environment resource creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-hosts
	//
	Hosts interface{} `field:"optional" json:"hosts" yaml:"hosts"`
	// > Amazon EVS is in public preview release and is subject to change.
	//
	// The initial VLAN subnets for the environment. Amazon EVS VLAN subnets have a minimum CIDR block size of /28 and a maximum size of /24. Amazon EVS VLAN subnet CIDR blocks must not overlap with other subnets in the VPC.
	//
	// Required for environment resource creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-initialvlans
	//
	InitialVlans interface{} `field:"optional" json:"initialVlans" yaml:"initialVlans"`
	// The AWS KMS key ID that AWS Secrets Manager uses to encrypt secrets that are associated with the environment.
	//
	// These secrets contain the VCF credentials that are needed to install vCenter Server, NSX, and SDDC Manager.
	//
	// By default, Amazon EVS use the AWS Secrets Manager managed key `aws/secretsmanager` . You can also specify a customer managed key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The license information that Amazon EVS requires to create an environment.
	//
	// Amazon EVS requires two license keys: a VCF solution key and a vSAN license key. The VCF solution key must cover a minimum of 256 cores. The vSAN license key must provide at least 110 TiB of vSAN capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-licenseinfo
	//
	LicenseInfo interface{} `field:"optional" json:"licenseInfo" yaml:"licenseInfo"`
	// The security groups that allow traffic between the Amazon EVS control plane and your VPC for service access.
	//
	// If a security group is not specified, Amazon EVS uses the default security group in your account for service access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-serviceaccesssecuritygroups
	//
	ServiceAccessSecurityGroups interface{} `field:"optional" json:"serviceAccessSecurityGroups" yaml:"serviceAccessSecurityGroups"`
	// The subnet that is used to establish connectivity between the Amazon EVS control plane and VPC.
	//
	// Amazon EVS uses this subnet to perform validations and create the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-serviceaccesssubnetid
	//
	ServiceAccessSubnetId *string `field:"optional" json:"serviceAccessSubnetId" yaml:"serviceAccessSubnetId"`
	// The Broadcom Site ID that is associated with your Amazon EVS environment.
	//
	// Amazon EVS uses the Broadcom Site ID that you provide to meet Broadcom VCF license usage reporting requirements for Amazon EVS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-siteid
	//
	SiteId *string `field:"optional" json:"siteId" yaml:"siteId"`
	// Metadata that assists with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Tags don't propagate to any other cluster or AWS resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Customer confirmation that the customer has purchased and will continue to maintain the required number of VCF software licenses to cover all physical processor cores in the Amazon EVS environment.
	//
	// Information about your VCF software in Amazon EVS will be shared with Broadcom to verify license compliance. Amazon EVS does not validate license keys. To validate license keys, visit the Broadcom support portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-termsaccepted
	//
	TermsAccepted interface{} `field:"optional" json:"termsAccepted" yaml:"termsAccepted"`
	// The DNS hostnames to be used by the VCF management appliances in your environment.
	//
	// For environment creation to be successful, each hostname entry must resolve to a domain name that you've registered in your DNS service of choice and configured in the DHCP option set of your VPC. DNS hostnames cannot be changed after environment creation has started.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-vcfhostnames
	//
	VcfHostnames interface{} `field:"optional" json:"vcfHostnames" yaml:"vcfHostnames"`
	// The VCF version of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-vcfversion
	//
	VcfVersion *string `field:"optional" json:"vcfVersion" yaml:"vcfVersion"`
	// The VPC associated with the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html#cfn-evs-environment-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

