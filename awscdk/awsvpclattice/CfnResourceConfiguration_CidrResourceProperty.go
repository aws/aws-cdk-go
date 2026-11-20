package awsvpclattice


// The CIDR ranges that make up a CIDR resource configuration.
//
// Only valid when ResourceConfigurationType is CIDR, and the resource gateway must have ResourceConfigDnsResolution set to IN_VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cidrResourceProperty := &CidrResourceProperty{
//   	CidrRanges: []*string{
//   		jsii.String("cidrRanges"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-cidrresource.html
//
type CfnResourceConfiguration_CidrResourceProperty struct {
	// The IPv4 CIDR ranges reachable through this resource configuration, in CIDR notation (for example, 10.0.0.0/24).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-cidrresource.html#cfn-vpclattice-resourceconfiguration-cidrresource-cidrranges
	//
	CidrRanges *[]*string `field:"required" json:"cidrRanges" yaml:"cidrRanges"`
}

