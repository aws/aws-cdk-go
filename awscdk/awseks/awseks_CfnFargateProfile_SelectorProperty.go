package awseks


// An object representing an AWS Fargate profile selector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectorProperty := &selectorProperty{
//   	namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	labels: []interface{}{
//   		&labelProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFargateProfile_SelectorProperty struct {
	// The Kubernetes namespace that the selector should match.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The Kubernetes labels that the selector should match.
	//
	// A pod must contain all of the labels that are specified in the selector for it to be considered a match.
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
}

