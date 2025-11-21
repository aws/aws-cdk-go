package awsamplify


// Describes the backend associated with an Amplify `Branch` .
//
// This property is available to Amplify Gen 2 apps only. When you deploy an application with Amplify Gen 2, you provision the app's backend infrastructure using Typescript code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backendProperty := &BackendProperty{
//   	StackArn: jsii.String("stackArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-branch-backend.html
//
type CfnBranch_BackendProperty struct {
	// The Amazon Resource Name (ARN) for the CloudFormation stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-branch-backend.html#cfn-amplify-branch-backend-stackarn
	//
	StackArn *string `field:"optional" json:"stackArn" yaml:"stackArn"`
}

