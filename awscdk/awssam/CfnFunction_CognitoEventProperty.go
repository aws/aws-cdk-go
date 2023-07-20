package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoEventProperty := &CognitoEventProperty{
//   	Trigger: jsii.String("trigger"),
//   	UserPool: jsii.String("userPool"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cognitoevent.html
//
type CfnFunction_CognitoEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cognitoevent.html#cfn-serverless-function-cognitoevent-trigger
	//
	Trigger *string `field:"required" json:"trigger" yaml:"trigger"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cognitoevent.html#cfn-serverless-function-cognitoevent-userpool
	//
	UserPool *string `field:"required" json:"userPool" yaml:"userPool"`
}

