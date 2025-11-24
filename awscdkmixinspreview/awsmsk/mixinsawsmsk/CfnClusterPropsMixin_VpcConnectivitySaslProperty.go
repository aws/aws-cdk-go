package mixinsawsmsk


// Details for client authentication using SASL for VpcConnectivity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivitysasl.html
//
type CfnClusterPropsMixin_VpcConnectivitySaslProperty struct {
	// Details for ClientAuthentication using IAM for VpcConnectivity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivitysasl.html#cfn-msk-cluster-vpcconnectivitysasl-iam
	//
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// Details for SASL/SCRAM client authentication for VpcConnectivity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivitysasl.html#cfn-msk-cluster-vpcconnectivitysasl-scram
	//
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

