package awsquicksight


// Properties for CfnAssignmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAssignmentMixinProps := &CfnAssignmentMixinProps{
//   	AssignmentName: jsii.String("assignmentName"),
//   	AssignmentStatus: jsii.String("assignmentStatus"),
//   	Identities: map[string][]*string{
//   		"identitiesKey": []*string{
//   			jsii.String("identities"),
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	PolicyArn: jsii.String("policyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html
//
type CfnAssignmentMixinProps struct {
	// The name of the assignment, also called a rule.
	//
	// The name must be unique within the AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html#cfn-quicksight-assignment-assignmentname
	//
	AssignmentName *string `field:"optional" json:"assignmentName" yaml:"assignmentName"`
	// The status of the assignment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html#cfn-quicksight-assignment-assignmentstatus
	//
	AssignmentStatus *string `field:"optional" json:"assignmentStatus" yaml:"assignmentStatus"`
	// The QuickSight users, groups, or both that you want to assign the policy to.
	//
	// This maps identity type to a list of identity names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html#cfn-quicksight-assignment-identities
	//
	Identities interface{} `field:"optional" json:"identities" yaml:"identities"`
	// The namespace that contains the assignment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html#cfn-quicksight-assignment-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The ARN for the IAM policy to apply to the QuickSight users and groups specified in this assignment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assignment.html#cfn-quicksight-assignment-policyarn
	//
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
}

