package mixinsawsmsk


// Includes all client authentication information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clientAuthenticationProperty := &ClientAuthenticationProperty{
//   	Sasl: &SaslProperty{
//   		Iam: &IamProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-serverlesscluster-clientauthentication.html
//
type CfnServerlessClusterPropsMixin_ClientAuthenticationProperty struct {
	// Details for client authentication using SASL.
	//
	// To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT` . If you choose `TLS_PLAINTEXT` , then you must also set `unauthenticated` to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-serverlesscluster-clientauthentication.html#cfn-msk-serverlesscluster-clientauthentication-sasl
	//
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
}

