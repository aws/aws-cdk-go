package awssynthetics


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dependencyProperty := &DependencyProperty{
//   	Reference: jsii.String("reference"),
//
//   	// the properties below are optional
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-dependency.html
//
type CfnCanary_DependencyProperty struct {
	// ARN of the Lambda layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-dependency.html#cfn-synthetics-canary-dependency-reference
	//
	Reference *string `field:"required" json:"reference" yaml:"reference"`
	// Type of dependency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-dependency.html#cfn-synthetics-canary-dependency-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

