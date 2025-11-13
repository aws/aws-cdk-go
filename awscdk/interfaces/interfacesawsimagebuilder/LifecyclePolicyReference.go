package interfacesawsimagebuilder


// A reference to a LifecyclePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecyclePolicyReference := &LifecyclePolicyReference{
//   	LifecyclePolicyArn: jsii.String("lifecyclePolicyArn"),
//   }
//
type LifecyclePolicyReference struct {
	// The Arn of the LifecyclePolicy resource.
	LifecyclePolicyArn *string `field:"required" json:"lifecyclePolicyArn" yaml:"lifecyclePolicyArn"`
}

