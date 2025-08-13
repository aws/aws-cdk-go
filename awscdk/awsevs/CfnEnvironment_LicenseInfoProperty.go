package awsevs


// The license information that Amazon EVS requires to create an environment.
//
// Amazon EVS requires two license keys: a VCF solution key and a vSAN license key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseInfoProperty := &LicenseInfoProperty{
//   	SolutionKey: jsii.String("solutionKey"),
//   	VsanKey: jsii.String("vsanKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-licenseinfo.html
//
type CfnEnvironment_LicenseInfoProperty struct {
	// The VCF solution key.
	//
	// This license unlocks VMware VCF product features, including vSphere, NSX, SDDC Manager, and vCenter Server. The VCF solution key must cover a minimum of 256 cores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-licenseinfo.html#cfn-evs-environment-licenseinfo-solutionkey
	//
	SolutionKey *string `field:"required" json:"solutionKey" yaml:"solutionKey"`
	// The VSAN license key.
	//
	// This license unlocks vSAN features. The vSAN license key must provide at least 110 TiB of vSAN capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-licenseinfo.html#cfn-evs-environment-licenseinfo-vsankey
	//
	VsanKey *string `field:"required" json:"vsanKey" yaml:"vsanKey"`
}

