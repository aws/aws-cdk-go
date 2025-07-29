package awsevs


// > Amazon EVS is in public preview release and is subject to change.
//
// The initial VLAN subnets for the environment. Amazon EVS VLAN subnets have a minimum CIDR block size of /28 and a maximum size of /24. Amazon EVS VLAN subnet CIDR blocks must not overlap with other subnets in the VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnEnvironment_InitialVlansProperty struct {
	// The edge VTEP VLAN subnet.
	//
	// This VLAN subnet manages traffic flowing between the internal network and external networks, including internet access and other site connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-edgevtep
	//
	EdgeVTep interface{} `field:"required" json:"edgeVTep" yaml:"edgeVTep"`
	// An additional VLAN subnet that can be used to extend VCF capabilities once configured.
	//
	// For example, you can configure an expansion VLAN subnet to use NSX Federation for centralized management and synchronization of multiple NSX deployments across different locations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-expansionvlan1
	//
	ExpansionVlan1 interface{} `field:"required" json:"expansionVlan1" yaml:"expansionVlan1"`
	// An additional VLAN subnet that can be used to extend VCF capabilities once configured.
	//
	// For example, you can configure an expansion VLAN subnet to use NSX Federation for centralized management and synchronization of multiple NSX deployments across different locations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-expansionvlan2
	//
	ExpansionVlan2 interface{} `field:"required" json:"expansionVlan2" yaml:"expansionVlan2"`
	// The HCX VLAN subnet.
	//
	// This VLAN subnet allows the HCX Interconnnect (IX) and HCX Network Extension (NE) to reach their peers and enable HCX Service Mesh creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-hcx
	//
	Hcx interface{} `field:"required" json:"hcx" yaml:"hcx"`
	// The NSX uplink VLAN subnet.
	//
	// This VLAN subnet allows connectivity to the NSX overlay network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-nsxuplink
	//
	NsxUpLink interface{} `field:"required" json:"nsxUpLink" yaml:"nsxUpLink"`
	// The host VMkernel management VLAN subnet.
	//
	// This VLAN subnet carries traffic for managing ESXi hosts and communicating with VMware vCenter Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-vmkmanagement
	//
	VmkManagement interface{} `field:"required" json:"vmkManagement" yaml:"vmkManagement"`
	// The VM management VLAN subnet.
	//
	// This VLAN subnet carries traffic for vSphere virtual machines.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-vmmanagement
	//
	VmManagement interface{} `field:"required" json:"vmManagement" yaml:"vmManagement"`
	// The vMotion VLAN subnet.
	//
	// This VLAN subnet carries traffic for vSphere vMotion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-vmotion
	//
	VMotion interface{} `field:"required" json:"vMotion" yaml:"vMotion"`
	// The vSAN VLAN subnet.
	//
	// This VLAN subnet carries the communication between ESXi hosts to implement a vSAN shared storage pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-vsan
	//
	VSan interface{} `field:"required" json:"vSan" yaml:"vSan"`
	// The VTEP VLAN subnet.
	//
	// This VLAN subnet handles internal network traffic between virtual machines within a VCF instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-initialvlans.html#cfn-evs-environment-initialvlans-vtep
	//
	VTep interface{} `field:"required" json:"vTep" yaml:"vTep"`
}

