package previewawsmskmixins


// Details for client authentication using SASL.
//
// To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT` . If you choose `TLS_PLAINTEXT` , then you must also set `unauthenticated` to true.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnClusterPropsMixin_SaslProperty struct {
	// Details for ClientAuthentication using IAM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-sasl.html#cfn-msk-cluster-sasl-iam
	//
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// Details for SASL/SCRAM client authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-sasl.html#cfn-msk-cluster-sasl-scram
	//
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

