package previewawsmskmixins


// Details for SASL/IAM client authentication for VpcConnectivity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcConnectivityIamProperty := &VpcConnectivityIamProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivityiam.html
//
type CfnClusterPropsMixin_VpcConnectivityIamProperty struct {
	// SASL/IAM authentication is enabled or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivityiam.html#cfn-msk-cluster-vpcconnectivityiam-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

