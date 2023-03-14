package awsiot


// Properties for defining a `CfnPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnPolicyProps := &cfnPolicyProps{
//   	policyDocument: policyDocument,
//
//   	// the properties below are optional
//   	policyName: jsii.String("policyName"),
//   }
//
type CfnPolicyProps struct {
	// The JSON document that describes the policy.
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The policy name.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

