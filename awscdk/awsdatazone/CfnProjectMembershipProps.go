package awsdatazone


// Properties for defining a `CfnProjectMembership`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectMembershipProps := &CfnProjectMembershipProps{
//   	Designation: jsii.String("designation"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Member: &MemberProperty{
//   		GroupIdentifier: jsii.String("groupIdentifier"),
//   		UserIdentifier: jsii.String("userIdentifier"),
//   	},
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectmembership.html
//
type CfnProjectMembershipProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectmembership.html#cfn-datazone-projectmembership-designation
	//
	Designation *string `field:"required" json:"designation" yaml:"designation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectmembership.html#cfn-datazone-projectmembership-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectmembership.html#cfn-datazone-projectmembership-member
	//
	Member interface{} `field:"required" json:"member" yaml:"member"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectmembership.html#cfn-datazone-projectmembership-projectidentifier
	//
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
}

