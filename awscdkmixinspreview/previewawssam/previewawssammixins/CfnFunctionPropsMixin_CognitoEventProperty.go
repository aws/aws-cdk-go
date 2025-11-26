package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cognitoEventProperty := &CognitoEventProperty{
//   	Trigger: jsii.String("trigger"),
//   	UserPool: jsii.String("userPool"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cognitoevent.html
//
type CfnFunctionPropsMixin_CognitoEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cognitoevent.html#cfn-serverless-function-cognitoevent-trigger
	//
	Trigger *string `field:"optional" json:"trigger" yaml:"trigger"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cognitoevent.html#cfn-serverless-function-cognitoevent-userpool
	//
	UserPool *string `field:"optional" json:"userPool" yaml:"userPool"`
}

