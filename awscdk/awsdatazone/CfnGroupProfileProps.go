package awsdatazone


// Properties for defining a `CfnGroupProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupProfileProps := &CfnGroupProfileProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	GroupIdentifier: jsii.String("groupIdentifier"),
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-groupprofile.html
//
type CfnGroupProfileProps struct {
	// The identifier of the Amazon DataZone domain in which the group profile would be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-groupprofile.html#cfn-datazone-groupprofile-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-groupprofile.html#cfn-datazone-groupprofile-groupidentifier
	//
	GroupIdentifier *string `field:"required" json:"groupIdentifier" yaml:"groupIdentifier"`
	// The status of the group profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-groupprofile.html#cfn-datazone-groupprofile-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

