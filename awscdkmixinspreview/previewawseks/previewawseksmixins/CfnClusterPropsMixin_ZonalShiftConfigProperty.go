package previewawseksmixins


// The configuration for zonal shift for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   zonalShiftConfigProperty := &ZonalShiftConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-zonalshiftconfig.html
//
type CfnClusterPropsMixin_ZonalShiftConfigProperty struct {
	// If zonal shift is enabled, AWS configures zonal autoshift for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-zonalshiftconfig.html#cfn-eks-cluster-zonalshiftconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

