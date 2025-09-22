package awsaps


// Properties for defining a `CfnResourcePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourcePolicyProps := &CfnResourcePolicyProps{
//   	PolicyDocument: jsii.String("policyDocument"),
//   	WorkspaceArn: jsii.String("workspaceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-resourcepolicy.html
//
type CfnResourcePolicyProps struct {
	// The JSON to use as the Resource-based Policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-resourcepolicy.html#cfn-aps-resourcepolicy-policydocument
	//
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// An ARN identifying a Workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-resourcepolicy.html#cfn-aps-resourcepolicy-workspacearn
	//
	WorkspaceArn *string `field:"required" json:"workspaceArn" yaml:"workspaceArn"`
}

