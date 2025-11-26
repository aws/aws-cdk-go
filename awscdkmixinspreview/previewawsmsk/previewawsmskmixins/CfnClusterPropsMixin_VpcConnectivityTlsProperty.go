package previewawsmskmixins


// Details for client authentication using TLS for VpcConnectivity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcConnectivityTlsProperty := &VpcConnectivityTlsProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivitytls.html
//
type CfnClusterPropsMixin_VpcConnectivityTlsProperty struct {
	// TLS authentication is enabled or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivitytls.html#cfn-msk-cluster-vpcconnectivitytls-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

