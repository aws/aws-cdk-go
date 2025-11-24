package mixinsawssynthetics


// A structure that contains information about a dependency for a canary.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dependencyProperty := &DependencyProperty{
//   	Reference: jsii.String("reference"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-dependency.html
//
type CfnCanaryPropsMixin_DependencyProperty struct {
	// The dependency reference.
	//
	// For Lambda layers, this is the ARN of the Lambda layer. For more information about Lambda ARN format, see [Lambda](https://docs.aws.amazon.com/lambda/latest/api/API_Layer.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-dependency.html#cfn-synthetics-canary-dependency-reference
	//
	Reference *string `field:"optional" json:"reference" yaml:"reference"`
	// The type of dependency.
	//
	// Valid value is `LambdaLayer` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-dependency.html#cfn-synthetics-canary-dependency-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

