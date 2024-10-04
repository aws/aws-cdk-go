package awseks


// The current zonal shift configuration to use for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zonalShiftConfigProperty := &ZonalShiftConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-zonalshiftconfig.html
//
type CfnCluster_ZonalShiftConfigProperty struct {
	// Set this value to true to enable zonal shift for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-zonalshiftconfig.html#cfn-eks-cluster-zonalshiftconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

