package previewawscodebuildmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pullRequestBuildPolicyProperty := &PullRequestBuildPolicyProperty{
//   	ApproverRoles: []*string{
//   		jsii.String("approverRoles"),
//   	},
//   	RequiresCommentApproval: jsii.String("requiresCommentApproval"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-pullrequestbuildpolicy.html
//
type CfnProjectPropsMixin_PullRequestBuildPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-pullrequestbuildpolicy.html#cfn-codebuild-project-pullrequestbuildpolicy-approverroles
	//
	ApproverRoles *[]*string `field:"optional" json:"approverRoles" yaml:"approverRoles"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-pullrequestbuildpolicy.html#cfn-codebuild-project-pullrequestbuildpolicy-requirescommentapproval
	//
	RequiresCommentApproval *string `field:"optional" json:"requiresCommentApproval" yaml:"requiresCommentApproval"`
}

