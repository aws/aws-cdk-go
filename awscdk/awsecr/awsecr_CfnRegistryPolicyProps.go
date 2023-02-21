package awsecr


// Properties for defining a `CfnRegistryPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyText interface{}
//
//   cfnRegistryPolicyProps := &cfnRegistryPolicyProps{
//   	policyText: policyText,
//   }
//
type CfnRegistryPolicyProps struct {
	// The JSON policy text for your registry.
	PolicyText interface{} `field:"required" json:"policyText" yaml:"policyText"`
}

