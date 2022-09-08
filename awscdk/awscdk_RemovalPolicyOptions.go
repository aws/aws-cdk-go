// An experiment to bundle the entire CDK into a single module
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   removalPolicyOptions := &removalPolicyOptions{
//   	applyToUpdateReplacePolicy: jsii.Boolean(false),
//   	default: monocdk.removalPolicy_DESTROY,
//   }
//
// Experimental.
type RemovalPolicyOptions struct {
	// Apply the same deletion policy to the resource's "UpdateReplacePolicy".
	// Experimental.
	ApplyToUpdateReplacePolicy *bool `field:"optional" json:"applyToUpdateReplacePolicy" yaml:"applyToUpdateReplacePolicy"`
	// The default policy to apply in case the removal policy is not defined.
	// Experimental.
	Default RemovalPolicy `field:"optional" json:"default" yaml:"default"`
}

