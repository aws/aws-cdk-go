package awsdatazone


// The details about a project member.
//
// Important - this data type is a UNION, so only one of the following members can be specified when used or returned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberProperty := &MemberProperty{
//   	GroupIdentifier: jsii.String("groupIdentifier"),
//   	UserIdentifier: jsii.String("userIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectmembership-member.html
//
type CfnProjectMembership_MemberProperty struct {
	// The ID of the group of a project member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectmembership-member.html#cfn-datazone-projectmembership-member-groupidentifier
	//
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
	// The user ID of a project member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectmembership-member.html#cfn-datazone-projectmembership-member-useridentifier
	//
	UserIdentifier *string `field:"optional" json:"userIdentifier" yaml:"userIdentifier"`
}

