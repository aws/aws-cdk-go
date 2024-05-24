package awsdatazone


// Properties for defining a `CfnUserProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserProfileProps := &CfnUserProfileProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	UserIdentifier: jsii.String("userIdentifier"),
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   	UserType: jsii.String("userType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html
//
type CfnUserProfileProps struct {
	// The identifier of the Amazon DataZone domain in which the user profile would be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-useridentifier
	//
	UserIdentifier *string `field:"required" json:"userIdentifier" yaml:"userIdentifier"`
	// The status of the user profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The type of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-usertype
	//
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
}

