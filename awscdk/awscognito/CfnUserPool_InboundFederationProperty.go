package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inboundFederationProperty := &InboundFederationProperty{
//   	LambdaArn: jsii.String("lambdaArn"),
//   	LambdaVersion: jsii.String("lambdaVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-inboundfederation.html
//
type CfnUserPool_InboundFederationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-inboundfederation.html#cfn-cognito-userpool-inboundfederation-lambdaarn
	//
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-inboundfederation.html#cfn-cognito-userpool-inboundfederation-lambdaversion
	//
	LambdaVersion *string `field:"optional" json:"lambdaVersion" yaml:"lambdaVersion"`
}

