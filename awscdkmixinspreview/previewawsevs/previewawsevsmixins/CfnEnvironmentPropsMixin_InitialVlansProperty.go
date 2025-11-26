package previewawsevsmixins


// The initial VLAN subnets for the environment.
//
// Amazon EVS VLAN subnets have a minimum CIDR block size of /28 and a maximum size of /24. Amazon EVS VLAN subnet CIDR blocks must not overlap with other subnets in the VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   initialVlansProperty := &InitialVlansProperty{
//   	EdgeVTep: &InitialVlanInfoProperty{
//   		Cidr: jsii.String("cidr"),
//   	},
//   	ExpansionVlan1: &InitialVlanInfoProperty{
//   		Cidr: jsii.String("cidr"),
//   	},
//   	ExpansionVlan2: &InitialVlanInfoProperty{
//   		Cidr: jsii.String("cidr"),
//   	},
//   	Hcx: &InitialVlanInfoProperty{
//   		Cidr: jsii.String("cidr"),
//   	},
//   	HcxNetworkAclId: jsii.String("hcxNetworkAclId"),
//   	IsHcxPublic: jsii.Boolean(false),
//   	NsxUpLink: &InitialVlanInfoProperty{
//   		Cidr: jsii.String("cidr"),
//   	},
//   	VmkManagement: &InitialVlanInfoProperty{
//   		Cidr: jsii.String("cidr"),
//   	},
//   	VmManagement: &InitialVlanInfoProperty{
//   		Cidr: jsii.String("cidr"),
//   	},
//   	VMotion: &InitialVlanInfoProperty{
//   		Cidr: jsii.String("cidr"),
//   	},
//   	VSan: &InitialVlanInfoProperty{
//   		Cidr: jsii.String("cidr"),
//   	},
//   	VTep: &InitialVlanInfoProperty{
//   		Cidr: jsii.String("cidr"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html
//
type CfnEnvironmentPropsMixin_InitialVlansProperty struct {
	// The edge VTEP VLAN subnet.
	//
	// This VLAN subnet manages traffic flowing between the internal network and external networks, including internet access and other site connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-edgevtep
	//
	EdgeVTep interface{} `field:"optional" json:"edgeVTep" yaml:"edgeVTep"`
	// An additional VLAN subnet that can be used to extend VCF capabilities once configured.
	//
	// For example, you can configure an expansion VLAN subnet to use NSX Federation for centralized management and synchronization of multiple NSX deployments across different locations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-expansionvlan1
	//
	ExpansionVlan1 interface{} `field:"optional" json:"expansionVlan1" yaml:"expansionVlan1"`
	// An additional VLAN subnet that can be used to extend VCF capabilities once configured.
	//
	// For example, you can configure an expansion VLAN subnet to use NSX Federation for centralized management and synchronization of multiple NSX deployments across different locations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-expansionvlan2
	//
	ExpansionVlan2 interface{} `field:"optional" json:"expansionVlan2" yaml:"expansionVlan2"`
	// The HCX VLAN subnet.
	//
	// This VLAN subnet allows the HCX Interconnnect (IX) and HCX Network Extension (NE) to reach their peers and enable HCX Service Mesh creation.
	//
	// If you plan to use a public HCX VLAN subnet, the following requirements must be met:
	//
	// - Must have a /28 netmask and be allocated from the IPAM public pool. Required for HCX internet access configuration.
	// - The HCX public VLAN CIDR block must be added to the VPC as a secondary CIDR block.
	// - Must have at least two Elastic IP addresses to be allocated from the public IPAM pool for HCX components.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-hcx
	//
	Hcx interface{} `field:"optional" json:"hcx" yaml:"hcx"`
	// A unique ID for a network access control list that the HCX VLAN uses.
	//
	// Required when `isHcxPublic` is set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-hcxnetworkaclid
	//
	HcxNetworkAclId *string `field:"optional" json:"hcxNetworkAclId" yaml:"hcxNetworkAclId"`
	// Determines if the HCX VLAN that Amazon EVS provisions is public or private.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-ishcxpublic
	//
	IsHcxPublic interface{} `field:"optional" json:"isHcxPublic" yaml:"isHcxPublic"`
	// The NSX uplink VLAN subnet.
	//
	// This VLAN subnet allows connectivity to the NSX overlay network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-nsxuplink
	//
	NsxUpLink interface{} `field:"optional" json:"nsxUpLink" yaml:"nsxUpLink"`
	// The host VMkernel management VLAN subnet.
	//
	// This VLAN subnet carries traffic for managing ESXi hosts and communicating with VMware vCenter Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-vmkmanagement
	//
	VmkManagement interface{} `field:"optional" json:"vmkManagement" yaml:"vmkManagement"`
	// The VM management VLAN subnet.
	//
	// This VLAN subnet carries traffic for vSphere virtual machines.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-vmmanagement
	//
	VmManagement interface{} `field:"optional" json:"vmManagement" yaml:"vmManagement"`
	// The vMotion VLAN subnet.
	//
	// This VLAN subnet carries traffic for vSphere vMotion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-vmotion
	//
	VMotion interface{} `field:"optional" json:"vMotion" yaml:"vMotion"`
	// The vSAN VLAN subnet.
	//
	// This VLAN subnet carries the communication between ESXi hosts to implement a vSAN shared storage pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-vsan
	//
	VSan interface{} `field:"optional" json:"vSan" yaml:"vSan"`
	// The VTEP VLAN subnet.
	//
	// This VLAN subnet handles internal network traffic between virtual machines within a VCF instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-vtep
	//
	VTep interface{} `field:"optional" json:"vTep" yaml:"vTep"`
}

