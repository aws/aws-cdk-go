package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sipMediaApplicationEndpointProperty := &SipMediaApplicationEndpointProperty{
//   	LambdaArn: jsii.String("lambdaArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-sipmediaapplication-sipmediaapplicationendpoint.html
//
type CfnSipMediaApplication_SipMediaApplicationEndpointProperty struct {
	// Valid Amazon Resource Name (ARN) of the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-sipmediaapplication-sipmediaapplicationendpoint.html#cfn-chime-sipmediaapplication-sipmediaapplicationendpoint-lambdaarn
	//
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
}

