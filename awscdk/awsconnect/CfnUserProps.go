package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnUser`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserProps := &CfnUserProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	PhoneConfig: &UserPhoneConfigProperty{
//   		PhoneType: jsii.String("phoneType"),
//
//   		// the properties below are optional
//   		AfterContactWorkTimeLimit: jsii.Number(123),
//   		AutoAccept: jsii.Boolean(false),
//   		DeskPhoneNumber: jsii.String("deskPhoneNumber"),
//   	},
//   	RoutingProfileArn: jsii.String("routingProfileArn"),
//   	SecurityProfileArns: []*string{
//   		jsii.String("securityProfileArns"),
//   	},
//   	Username: jsii.String("username"),
//
//   	// the properties below are optional
//   	DirectoryUserId: jsii.String("directoryUserId"),
//   	HierarchyGroupArn: jsii.String("hierarchyGroupArn"),
//   	IdentityInfo: &UserIdentityInfoProperty{
//   		Email: jsii.String("email"),
//   		FirstName: jsii.String("firstName"),
//   		LastName: jsii.String("lastName"),
//   		Mobile: jsii.String("mobile"),
//   		SecondaryEmail: jsii.String("secondaryEmail"),
//   	},
//   	Password: jsii.String("password"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

