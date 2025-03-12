package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   removalPolicyOptions := &RemovalPolicyOptions{
//   	ApplyToUpdateReplacePolicy: jsii.Boolean(false),
//   	Default: cdk.RemovalPolicy_DESTROY,
//   }
//
type RemovalPolicyOptions struct {
	// Apply the same deletion policy to the resource's "UpdateReplacePolicy".
	// Default: true.
	//
	ApplyToUpdateReplacePolicy *bool `field:"optional" json:"applyToUpdateReplacePolicy" yaml:"applyToUpdateReplacePolicy"`
	// The default policy to apply in case the removal policy is not defined.
	// Default: - Default value is resource specific. To determine the default value for a resource,
	// please consult that specific resource's documentation.
	//
	Default RemovalPolicy `field:"optional" json:"default" yaml:"default"`
}

