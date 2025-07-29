package awssagemaker


// A key-value pair representing a parameter used in the CloudFormation stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStackParameterProperty := &CfnStackParameterProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-cfnstackparameter.html
//
type CfnProject_CfnStackParameterProperty struct {
	// The name of the CloudFormation parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-cfnstackparameter.html#cfn-sagemaker-project-cfnstackparameter-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the CloudFormation parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-cfnstackparameter.html#cfn-sagemaker-project-cfnstackparameter-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

