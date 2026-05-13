package awsdatazone


// The project membership assignment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   projectMembershipAssignmentProperty := &ProjectMembershipAssignmentProperty{
//   	Designation: jsii.String("designation"),
//   	Member: &MemberProperty{
//   		GroupIdentifier: jsii.String("groupIdentifier"),
//   		UserIdentifier: jsii.String("userIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-projectmembershipassignment.html
//
type CfnProjectPropsMixin_ProjectMembershipAssignmentProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-projectmembershipassignment.html#cfn-datazone-project-projectmembershipassignment-designation
	//
	Designation *string `field:"optional" json:"designation" yaml:"designation"`
	// The member of the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-projectmembershipassignment.html#cfn-datazone-project-projectmembershipassignment-member
	//
	Member interface{} `field:"optional" json:"member" yaml:"member"`
}

