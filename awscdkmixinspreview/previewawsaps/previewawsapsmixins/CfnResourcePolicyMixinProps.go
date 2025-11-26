package previewawsapsmixins


// Properties for CfnResourcePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourcePolicyMixinProps := &CfnResourcePolicyMixinProps{
//   	PolicyDocument: jsii.String("policyDocument"),
//   	WorkspaceArn: jsii.String("workspaceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-resourcepolicy.html
//
type CfnResourcePolicyMixinProps struct {
	// The JSON to use as the Resource-based Policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-resourcepolicy.html#cfn-aps-resourcepolicy-policydocument
	//
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// An ARN identifying a Workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-resourcepolicy.html#cfn-aps-resourcepolicy-workspacearn
	//
	WorkspaceArn *string `field:"optional" json:"workspaceArn" yaml:"workspaceArn"`
}

