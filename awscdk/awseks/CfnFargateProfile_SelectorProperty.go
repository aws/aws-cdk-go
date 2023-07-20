package awseks


// An object representing an AWS Fargate profile selector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectorProperty := &SelectorProperty{
//   	Namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	Labels: []interface{}{
//   		&LabelProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-fargateprofile-selector.html
//
type CfnFargateProfile_SelectorProperty struct {
	// The Kubernetes namespace that the selector should match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-fargateprofile-selector.html#cfn-eks-fargateprofile-selector-namespace
	//
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The Kubernetes labels that the selector should match.
	//
	// A pod must contain all of the labels that are specified in the selector for it to be considered a match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-fargateprofile-selector.html#cfn-eks-fargateprofile-selector-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
}

