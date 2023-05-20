package awsmsk


// Details for SASL/IAM client authentication for VpcConnectivity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectivityIamProperty := &VpcConnectivityIamProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
type CfnCluster_VpcConnectivityIamProperty struct {
	// SASL/IAM authentication is enabled or not.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

