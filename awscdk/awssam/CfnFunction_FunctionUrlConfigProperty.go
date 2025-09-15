package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionUrlConfigProperty := &FunctionUrlConfigProperty{
//   	AuthType: jsii.String("authType"),
//
//   	// the properties below are optional
//   	Cors: jsii.String("cors"),
//   	InvokeMode: jsii.String("invokeMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-functionurlconfig.html
//
type CfnFunction_FunctionUrlConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-functionurlconfig.html#cfn-serverless-function-functionurlconfig-authtype
	//
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-functionurlconfig.html#cfn-serverless-function-functionurlconfig-cors
	//
	Cors interface{} `field:"optional" json:"cors" yaml:"cors"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-functionurlconfig.html#cfn-serverless-function-functionurlconfig-invokemode
	//
	InvokeMode *string `field:"optional" json:"invokeMode" yaml:"invokeMode"`
}

