package awsevs


// A check on the environment to identify environment health and validate VMware VCF licensing compliance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   checkProperty := &CheckProperty{
//   	Result: jsii.String("result"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ImpairedSince: jsii.String("impairedSince"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-check.html
//
type CfnEnvironment_CheckProperty struct {
	// The check result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-check.html#cfn-evs-environment-check-result
	//
	Result *string `field:"required" json:"result" yaml:"result"`
	// The check type. Amazon EVS performs the following checks.
	//
	// - `KEY_REUSE` : checks that the VCF license key is not used by another Amazon EVS environment. This check fails if a used license is added to the environment.
	// - `KEY_COVERAGE` : checks that your VCF license key allocates sufficient vCPU cores for all deployed hosts. The check fails when any assigned hosts in the EVS environment are not covered by license keys, or when any unassigned hosts cannot be covered by available vCPU cores in keys.
	// - `REACHABILITY` : checks that the Amazon EVS control plane has a persistent connection to SDDC Manager. If Amazon EVS cannot reach the environment, this check fails.
	// - `HOST_COUNT` : Checks that your environment has a minimum of 4 hosts, which is a requirement for VCF 5.2.1.
	//
	// If this check fails, you will need to add hosts so that your environment meets this minimum requirement. Amazon EVS only supports environments with 4-16 hosts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-check.html#cfn-evs-environment-check-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The time when environment health began to be impaired.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-check.html#cfn-evs-environment-check-impairedsince
	//
	ImpairedSince *string `field:"optional" json:"impairedSince" yaml:"impairedSince"`
}

