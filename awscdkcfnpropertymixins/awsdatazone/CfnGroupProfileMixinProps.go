package awsdatazone


// Properties for CfnGroupProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnGroupProfileMixinProps := &CfnGroupProfileMixinProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	GroupIdentifier: jsii.String("groupIdentifier"),
//   	GroupType: jsii.String("groupType"),
//   	RolePrincipalArn: jsii.String("rolePrincipalArn"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-groupprofile.html
//
type CfnGroupProfileMixinProps struct {
	// The identifier of the Amazon DataZone domain in which a group profile exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-groupprofile.html#cfn-datazone-groupprofile-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the group of a project member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-groupprofile.html#cfn-datazone-groupprofile-groupidentifier
	//
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
	// The type of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-groupprofile.html#cfn-datazone-groupprofile-grouptype
	//
	GroupType *string `field:"optional" json:"groupType" yaml:"groupType"`
	// The ARN of the role principal for the group profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-groupprofile.html#cfn-datazone-groupprofile-roleprincipalarn
	//
	RolePrincipalArn *string `field:"optional" json:"rolePrincipalArn" yaml:"rolePrincipalArn"`
	// The status of a group profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-groupprofile.html#cfn-datazone-groupprofile-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

