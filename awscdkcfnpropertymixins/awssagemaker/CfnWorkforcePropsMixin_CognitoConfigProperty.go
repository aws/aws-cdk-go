package awssagemaker


// The configuration of an Amazon Cognito workforce.
//
// A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cognitoConfigProperty := &CognitoConfigProperty{
//   	ClientId: jsii.String("clientId"),
//   	UserPool: jsii.String("userPool"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-cognitoconfig.html
//
type CfnWorkforcePropsMixin_CognitoConfigProperty struct {
	// The client ID for your Amazon Cognito user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-cognitoconfig.html#cfn-sagemaker-workforce-cognitoconfig-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The ID for your Amazon Cognito user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-cognitoconfig.html#cfn-sagemaker-workforce-cognitoconfig-userpool
	//
	UserPool *string `field:"optional" json:"userPool" yaml:"userPool"`
}

