package awsmsk


// Details for client authentication using TLS for vpcConnectivity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectivityTlsProperty := &VpcConnectivityTlsProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivitytls.html
//
type CfnCluster_VpcConnectivityTlsProperty struct {
	// TLS authentication is enabled or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivitytls.html#cfn-msk-cluster-vpcconnectivitytls-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

