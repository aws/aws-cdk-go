package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnUser`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserProps := &cfnUserProps{
//   	instanceArn: jsii.String("instanceArn"),
//   	phoneConfig: &userPhoneConfigProperty{
//   		phoneType: jsii.String("phoneType"),
//
//   		// the properties below are optional
//   		afterContactWorkTimeLimit: jsii.Number(123),
//   		autoAccept: jsii.Boolean(false),
//   		deskPhoneNumber: jsii.String("deskPhoneNumber"),
//   	},
//   	routingProfileArn: jsii.String("routingProfileArn"),
//   	securityProfileArns: []*string{
//   		jsii.String("securityProfileArns"),
//   	},
//   	username: jsii.String("username"),
//
//   	// the properties below are optional
//   	directoryUserId: jsii.String("directoryUserId"),
//   	hierarchyGroupArn: jsii.String("hierarchyGroupArn"),
//   	identityInfo: &userIdentityInfoProperty{
//   		email: jsii.String("email"),
//   		firstName: jsii.String("firstName"),
//   		lastName: jsii.String("lastName"),
//   		mobile: jsii.String("mobile"),
//   		secondaryEmail: jsii.String("secondaryEmail"),
//   	},
//   	password: jsii.String("password"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnUserProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Information about the phone configuration for the user.
	PhoneConfig interface{} `field:"required" json:"phoneConfig" yaml:"phoneConfig"`
	// The Amazon Resource Name (ARN) of the user's routing profile.
	RoutingProfileArn *string `field:"required" json:"routingProfileArn" yaml:"routingProfileArn"`
	// The Amazon Resource Name (ARN) of the user's security profile.
	SecurityProfileArns *[]*string `field:"required" json:"securityProfileArns" yaml:"securityProfileArns"`
	// The user name assigned to the user account.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The identifier of the user account in the directory used for identity management.
	DirectoryUserId *string `field:"optional" json:"directoryUserId" yaml:"directoryUserId"`
	// The Amazon Resource Name (ARN) of the user's hierarchy group.
	HierarchyGroupArn *string `field:"optional" json:"hierarchyGroupArn" yaml:"hierarchyGroupArn"`
	// Information about the user identity.
	IdentityInfo interface{} `field:"optional" json:"identityInfo" yaml:"identityInfo"`
	// The user's password.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

