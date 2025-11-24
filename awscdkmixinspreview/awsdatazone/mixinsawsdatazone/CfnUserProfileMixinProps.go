package mixinsawsdatazone


// Properties for CfnUserProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserProfileMixinProps := &CfnUserProfileMixinProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Status: jsii.String("status"),
//   	UserIdentifier: jsii.String("userIdentifier"),
//   	UserType: jsii.String("userType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html
//
type CfnUserProfileMixinProps struct {
	// The identifier of a Amazon DataZone domain in which a user profile exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The status of the user profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The identifier of the user for which the user profile is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-useridentifier
	//
	UserIdentifier *string `field:"optional" json:"userIdentifier" yaml:"userIdentifier"`
	// The user type of the user for which the user profile is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-usertype
	//
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
}

