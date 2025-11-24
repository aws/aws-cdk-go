package mixinsawsrefactorspaces


// The input for the AWS Lambda endpoint type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lambdaEndpointInputProperty := &LambdaEndpointInputProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-service-lambdaendpointinput.html
//
type CfnServicePropsMixin_LambdaEndpointInputProperty struct {
	// The Amazon Resource Name (ARN) of the Lambda function or alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-service-lambdaendpointinput.html#cfn-refactorspaces-service-lambdaendpointinput-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

