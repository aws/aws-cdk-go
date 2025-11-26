package previewawsmskmixins


// Details for SASL/SCRAM client authentication for VpcConnectivity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcConnectivityScramProperty := &VpcConnectivityScramProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivityscram.html
//
type CfnClusterPropsMixin_VpcConnectivityScramProperty struct {
	// SASL/SCRAM authentication is enabled or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivityscram.html#cfn-msk-cluster-vpcconnectivityscram-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

