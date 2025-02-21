package awsmsk


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivitysasl.html
//
type CfnCluster_VpcConnectivitySaslProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivitysasl.html#cfn-msk-cluster-vpcconnectivitysasl-iam
	//
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivitysasl.html#cfn-msk-cluster-vpcconnectivitysasl-scram
	//
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

