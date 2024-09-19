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
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-serverlesscluster-clientauthentication.html
//
type CfnServerlessCluster_ClientAuthenticationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-serverlesscluster-clientauthentication.html#cfn-msk-serverlesscluster-clientauthentication-sasl
	//
	Sasl interface{} `field:"required" json:"sasl" yaml:"sasl"`
}

