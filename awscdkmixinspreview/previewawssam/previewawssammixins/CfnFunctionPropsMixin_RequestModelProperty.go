package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestModelProperty := &RequestModelProperty{
//   	Model: jsii.String("model"),
//   	Required: jsii.Boolean(false),
//   	ValidateBody: jsii.Boolean(false),
//   	ValidateParameters: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-requestmodel.html
//
type CfnFunctionPropsMixin_RequestModelProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-requestmodel.html#cfn-serverless-function-requestmodel-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-requestmodel.html#cfn-serverless-function-requestmodel-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-requestmodel.html#cfn-serverless-function-requestmodel-validatebody
	//
	ValidateBody interface{} `field:"optional" json:"validateBody" yaml:"validateBody"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-requestmodel.html#cfn-serverless-function-requestmodel-validateparameters
	//
	ValidateParameters interface{} `field:"optional" json:"validateParameters" yaml:"validateParameters"`
}

