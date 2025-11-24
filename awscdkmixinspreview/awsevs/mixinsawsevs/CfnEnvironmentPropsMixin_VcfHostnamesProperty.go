package mixinsawsevs


// The DNS hostnames that Amazon EVS uses to install VMware vCenter Server, NSX, SDDC Manager, and Cloud Builder.
//
// Each hostname must be unique, and resolve to a domain name that you've registered in your DNS service of choice. Hostnames cannot be changed.
//
// VMware VCF requires the deployment of two NSX Edge nodes, and three NSX Manager virtual machines.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vcfHostnamesProperty := &VcfHostnamesProperty{
//   	CloudBuilder: jsii.String("cloudBuilder"),
//   	Nsx: jsii.String("nsx"),
//   	NsxEdge1: jsii.String("nsxEdge1"),
//   	NsxEdge2: jsii.String("nsxEdge2"),
//   	NsxManager1: jsii.String("nsxManager1"),
//   	NsxManager2: jsii.String("nsxManager2"),
//   	NsxManager3: jsii.String("nsxManager3"),
//   	SddcManager: jsii.String("sddcManager"),
//   	VCenter: jsii.String("vCenter"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-vcfhostnames.html
//
type CfnEnvironmentPropsMixin_VcfHostnamesProperty struct {
	// The hostname for VMware Cloud Builder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-vcfhostnames.html#cfn-evs-environment-vcfhostnames-cloudbuilder
	//
	CloudBuilder *string `field:"optional" json:"cloudBuilder" yaml:"cloudBuilder"`
	// The VMware NSX hostname.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-vcfhostnames.html#cfn-evs-environment-vcfhostnames-nsx
	//
	Nsx *string `field:"optional" json:"nsx" yaml:"nsx"`
	// The hostname for the first NSX Edge node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-vcfhostnames.html#cfn-evs-environment-vcfhostnames-nsxedge1
	//
	NsxEdge1 *string `field:"optional" json:"nsxEdge1" yaml:"nsxEdge1"`
	// The hostname for the second NSX Edge node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-vcfhostnames.html#cfn-evs-environment-vcfhostnames-nsxedge2
	//
	NsxEdge2 *string `field:"optional" json:"nsxEdge2" yaml:"nsxEdge2"`
	// The hostname for the first VMware NSX Manager virtual machine (VM).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-vcfhostnames.html#cfn-evs-environment-vcfhostnames-nsxmanager1
	//
	NsxManager1 *string `field:"optional" json:"nsxManager1" yaml:"nsxManager1"`
	// The hostname for the second VMware NSX Manager virtual machine (VM).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-vcfhostnames.html#cfn-evs-environment-vcfhostnames-nsxmanager2
	//
	NsxManager2 *string `field:"optional" json:"nsxManager2" yaml:"nsxManager2"`
	// The hostname for the third VMware NSX Manager virtual machine (VM).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-vcfhostnames.html#cfn-evs-environment-vcfhostnames-nsxmanager3
	//
	NsxManager3 *string `field:"optional" json:"nsxManager3" yaml:"nsxManager3"`
	// The hostname for SDDC Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-vcfhostnames.html#cfn-evs-environment-vcfhostnames-sddcmanager
	//
	SddcManager *string `field:"optional" json:"sddcManager" yaml:"sddcManager"`
	// The VMware vCenter hostname.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-vcfhostnames.html#cfn-evs-environment-vcfhostnames-vcenter
	//
	VCenter *string `field:"optional" json:"vCenter" yaml:"vCenter"`
}

