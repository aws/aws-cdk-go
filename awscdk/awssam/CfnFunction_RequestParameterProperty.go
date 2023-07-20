package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestParameterProperty := &RequestParameterProperty{
//   	Caching: jsii.Boolean(false),
//   	Required: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-requestparameter.html
//
type CfnFunction_RequestParameterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-requestparameter.html#cfn-serverless-function-requestparameter-caching
	//
	Caching interface{} `field:"optional" json:"caching" yaml:"caching"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-requestparameter.html#cfn-serverless-function-requestparameter-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

