package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   saslProperty := &SaslProperty{
//   	Iam: &IamProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Scram: &ScramProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-sasl.html
//
type CfnCluster_SaslProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-sasl.html#cfn-msk-cluster-sasl-iam
	//
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-sasl.html#cfn-msk-cluster-sasl-scram
	//
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

