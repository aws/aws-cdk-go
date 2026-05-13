package awsdatazone


// The member of the project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   memberProperty := &MemberProperty{
//   	GroupIdentifier: jsii.String("groupIdentifier"),
//   	UserIdentifier: jsii.String("userIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-member.html
//
type CfnProjectPropsMixin_MemberProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-member.html#cfn-datazone-project-member-groupidentifier
	//
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-member.html#cfn-datazone-project-member-useridentifier
	//
	UserIdentifier *string `field:"optional" json:"userIdentifier" yaml:"userIdentifier"`
}

