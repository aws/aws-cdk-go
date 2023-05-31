package awsmsk


// Details for client authentication using SASL for VpcConnectivity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectivitySaslProperty := &VpcConnectivitySaslProperty{
//   	Iam: &VpcConnectivityIamProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Scram: &VpcConnectivityScramProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnCluster_VpcConnectivitySaslProperty struct {
	// Details for ClientAuthentication using IAM for VpcConnectivity.
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// Details for SASL/SCRAM client authentication for VpcConnectivity.
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

