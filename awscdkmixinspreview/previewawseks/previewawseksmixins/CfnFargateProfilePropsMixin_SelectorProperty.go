package previewawseksmixins


// An object representing an AWS Fargate profile selector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   selectorProperty := &SelectorProperty{
//   	Labels: []interface{}{
//   		&LabelProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-fargateprofile-selector.html
//
type CfnFargateProfilePropsMixin_SelectorProperty struct {
	// The Kubernetes labels that the selector should match.
	//
	// A pod must contain all of the labels that are specified in the selector for it to be considered a match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-fargateprofile-selector.html#cfn-eks-fargateprofile-selector-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// The Kubernetes `namespace` that the selector should match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-fargateprofile-selector.html#cfn-eks-fargateprofile-selector-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

