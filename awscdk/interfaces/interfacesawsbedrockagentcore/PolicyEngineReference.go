package interfacesawsbedrockagentcore


// A reference to a PolicyEngine resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyEngineReference := &PolicyEngineReference{
//   	PolicyEngineArn: jsii.String("policyEngineArn"),
//   }
//
type PolicyEngineReference struct {
	// The PolicyEngineArn of the PolicyEngine resource.
	PolicyEngineArn *string `field:"required" json:"policyEngineArn" yaml:"policyEngineArn"`
}

