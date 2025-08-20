package awsservicecatalog


// Properties for defining a `CfnTagOption`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTagOptionProps := &CfnTagOptionProps{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//
//   	// the properties below are optional
//   	Active: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html
//
type CfnTagOptionProps struct {
	// The TagOption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html#cfn-servicecatalog-tagoption-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The TagOption value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html#cfn-servicecatalog-tagoption-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
	// The TagOption active state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html#cfn-servicecatalog-tagoption-active
	//
	Active interface{} `field:"optional" json:"active" yaml:"active"`
}

