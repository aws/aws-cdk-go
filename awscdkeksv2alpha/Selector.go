package awscdkeksv2alpha


// Fargate profile selector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//
//   selector := &Selector{
//   	Namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   }
//
// Experimental.
type Selector struct {
	// The Kubernetes namespace that the selector should match.
	//
	// You must specify a namespace for a selector. The selector only matches pods
	// that are created in this namespace, but you can create multiple selectors
	// to target multiple namespaces.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The Kubernetes labels that the selector should match.
	//
	// A pod must contain
	// all of the labels that are specified in the selector for it to be
	// considered a match.
	// Default: - all pods within the namespace will be selected.
	//
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

