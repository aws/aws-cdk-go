package awssecretsmanager


// Construction properties for a ResourcePolicy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//
//   resourcePolicyProps := &resourcePolicyProps{
//   	secret: secret,
//   }
//
type ResourcePolicyProps struct {
	// The secret to attach a resource-based permissions policy.
	Secret ISecret `field:"required" json:"secret" yaml:"secret"`
}

