package awscodebuild


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pullRequestBuildPolicyProperty := &PullRequestBuildPolicyProperty{
//   	RequiresCommentApproval: jsii.String("requiresCommentApproval"),
//
//   	// the properties below are optional
//   	ApproverRoles: []*string{
//   		jsii.String("approverRoles"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-pullrequestbuildpolicy.html
//
type CfnProject_PullRequestBuildPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-pullrequestbuildpolicy.html#cfn-codebuild-project-pullrequestbuildpolicy-requirescommentapproval
	//
	RequiresCommentApproval *string `field:"required" json:"requiresCommentApproval" yaml:"requiresCommentApproval"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-pullrequestbuildpolicy.html#cfn-codebuild-project-pullrequestbuildpolicy-approverroles
	//
	ApproverRoles *[]*string `field:"optional" json:"approverRoles" yaml:"approverRoles"`
}

