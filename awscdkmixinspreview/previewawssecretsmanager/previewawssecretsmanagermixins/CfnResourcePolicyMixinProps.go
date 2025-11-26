package previewawssecretsmanagermixins


// Properties for CfnResourcePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var resourcePolicy interface{}
//
//   cfnResourcePolicyMixinProps := &CfnResourcePolicyMixinProps{
//   	BlockPublicPolicy: jsii.Boolean(false),
//   	ResourcePolicy: resourcePolicy,
//   	SecretId: jsii.String("secretId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-resourcepolicy.html
//
type CfnResourcePolicyMixinProps struct {
	// Specifies whether to block resource-based policies that allow broad access to the secret.
	//
	// By default, Secrets Manager blocks policies that allow broad access, for example those that use a wildcard for the principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-resourcepolicy.html#cfn-secretsmanager-resourcepolicy-blockpublicpolicy
	//
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// A JSON-formatted string for an AWS resource-based policy.
	//
	// For example policies, see [Permissions policy examples](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-resourcepolicy.html#cfn-secretsmanager-resourcepolicy-resourcepolicy
	//
	ResourcePolicy interface{} `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// The ARN or name of the secret to attach the resource-based policy.
	//
	// For an ARN, we recommend that you specify a complete ARN rather than a partial ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-resourcepolicy.html#cfn-secretsmanager-resourcepolicy-secretid
	//
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
}

