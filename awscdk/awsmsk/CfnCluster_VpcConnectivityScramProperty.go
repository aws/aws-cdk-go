package awsmsk


// Details for SASL/SCRAM client authentication for vpcConnectivity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectivityScramProperty := &VpcConnectivityScramProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
type CfnCluster_VpcConnectivityScramProperty struct {
	// SASL/SCRAM authentication is enabled or not.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

