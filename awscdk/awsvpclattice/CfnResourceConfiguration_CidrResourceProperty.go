package awsvpclattice


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-cidrresource.html#cfn-vpclattice-resourceconfiguration-cidrresource-cidrranges
	//
	CidrRanges *[]*string `field:"required" json:"cidrRanges" yaml:"cidrRanges"`
}

