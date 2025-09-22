package awsopensearchserverless


// A reference to a LifecyclePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecyclePolicyReference := &LifecyclePolicyReference{
//   	LifecyclePolicyName: jsii.String("lifecyclePolicyName"),
//   	Type: jsii.String("type"),
//   }
//
type LifecyclePolicyReference struct {
	// The Name of the LifecyclePolicy resource.
	LifecyclePolicyName *string `field:"required" json:"lifecyclePolicyName" yaml:"lifecyclePolicyName"`
	// The Type of the LifecyclePolicy resource.
	Type *string `field:"required" json:"type" yaml:"type"`
}

