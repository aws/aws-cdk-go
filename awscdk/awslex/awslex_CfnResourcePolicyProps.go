package awslex


// Properties for defining a `CfnResourcePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnResourcePolicyProps := &cfnResourcePolicyProps{
//   	policy: policy,
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnResourcePolicyProps struct {
	// A resource policy to add to the resource.
	//
	// The policy is a JSON structure that contains one or more statements that define the policy. The policy must follow IAM syntax. If the policy isn't valid, Amazon Lex returns a validation exception.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

