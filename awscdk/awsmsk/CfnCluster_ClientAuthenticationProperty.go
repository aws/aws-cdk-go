package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientAuthenticationProperty := &ClientAuthenticationProperty{
//   	Sasl: &SaslProperty{
//   		Iam: &IamProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		Scram: &ScramProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Tls: &TlsProperty{
//   		CertificateAuthorityArnList: []*string{
//   			jsii.String("certificateAuthorityArnList"),
//   		},
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Unauthenticated: &UnauthenticatedProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html
//
type CfnCluster_ClientAuthenticationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html#cfn-msk-cluster-clientauthentication-sasl
	//
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html#cfn-msk-cluster-clientauthentication-tls
	//
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html#cfn-msk-cluster-clientauthentication-unauthenticated
	//
	Unauthenticated interface{} `field:"optional" json:"unauthenticated" yaml:"unauthenticated"`
}

