package awsquicksight


// Properties for defining a `CfnAssignment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssignmentProps := &CfnAssignmentProps{
//   	AssignmentName: jsii.String("assignmentName"),
//   	AssignmentStatus: jsii.String("assignmentStatus"),
//   	Namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	Identities: map[string][]*string{
//   		"identitiesKey": []*string{
//   			jsii.String("identities"),
//   		},
//   	},
//   	PolicyArn: jsii.String("policyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html
//
type CfnAssignmentProps struct {
	// The name of the assignment, also called a rule.
	//
	// The name must be unique within the AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html#cfn-quicksight-assignment-assignmentname
	//
	AssignmentName *string `field:"required" json:"assignmentName" yaml:"assignmentName"`
	// The status of the assignment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html#cfn-quicksight-assignment-assignmentstatus
	//
	AssignmentStatus *string `field:"required" json:"assignmentStatus" yaml:"assignmentStatus"`
	// The namespace that contains the assignment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html#cfn-quicksight-assignment-namespace
	//
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The QuickSight users, groups, or both that you want to assign the policy to.
	//
	// This maps identity type to a list of identity names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html#cfn-quicksight-assignment-identities
	//
	Identities interface{} `field:"optional" json:"identities" yaml:"identities"`
	// The ARN for the IAM policy to apply to the QuickSight users and groups specified in this assignment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html#cfn-quicksight-assignment-policyarn
	//
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
}

