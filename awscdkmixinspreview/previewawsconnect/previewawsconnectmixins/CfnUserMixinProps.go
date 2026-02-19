package previewawsconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnUserPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserMixinProps := &CfnUserMixinProps{
//   	AfterContactWorkConfigs: []interface{}{
//   		&AfterContactWorkConfigPerChannelProperty{
//   			AfterContactWorkConfig: &AfterContactWorkConfigProperty{
//   				AfterContactWorkTimeLimit: jsii.Number(123),
//   			},
//   			AgentFirstCallbackAfterContactWorkConfig: &AfterContactWorkConfigProperty{
//   				AfterContactWorkTimeLimit: jsii.Number(123),
//   			},
//   			Channel: jsii.String("channel"),
//   		},
//   	},
//   	AutoAcceptConfigs: []interface{}{
//   		&AutoAcceptConfigProperty{
//   			AgentFirstCallbackAutoAccept: jsii.Boolean(false),
//   			AutoAccept: jsii.Boolean(false),
//   			Channel: jsii.String("channel"),
//   		},
//   	},
//   	DirectoryUserId: jsii.String("directoryUserId"),
//   	HierarchyGroupArn: jsii.String("hierarchyGroupArn"),
//   	IdentityInfo: &UserIdentityInfoProperty{
//   		Email: jsii.String("email"),
//   		FirstName: jsii.String("firstName"),
//   		LastName: jsii.String("lastName"),
//   		Mobile: jsii.String("mobile"),
//   		SecondaryEmail: jsii.String("secondaryEmail"),
//   	},
//   	InstanceArn: jsii.String("instanceArn"),
//   	Password: jsii.String("password"),
//   	PersistentConnectionConfigs: []interface{}{
//   		&PersistentConnectionConfigProperty{
//   			Channel: jsii.String("channel"),
//   			PersistentConnection: jsii.Boolean(false),
//   		},
//   	},
//   	PhoneConfig: &UserPhoneConfigProperty{
//   		AfterContactWorkTimeLimit: jsii.Number(123),
//   		AutoAccept: jsii.Boolean(false),
//   		DeskPhoneNumber: jsii.String("deskPhoneNumber"),
//   		PersistentConnection: jsii.Boolean(false),
//   		PhoneType: jsii.String("phoneType"),
//   	},
//   	PhoneNumberConfigs: []interface{}{
//   		&PhoneNumberConfigProperty{
//   			Channel: jsii.String("channel"),
//   			PhoneNumber: jsii.String("phoneNumber"),
//   			PhoneType: jsii.String("phoneType"),
//   		},
//   	},
//   	RoutingProfileArn: jsii.String("routingProfileArn"),
//   	SecurityProfileArns: []*string{
//   		jsii.String("securityProfileArns"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Username: jsii.String("username"),
//   	UserProficiencies: []interface{}{
//   		&UserProficiencyProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			AttributeValue: jsii.String("attributeValue"),
//   			Level: jsii.Number(123),
//   		},
//   	},
//   	VoiceEnhancementConfigs: []interface{}{
//   		&VoiceEnhancementConfigProperty{
//   			Channel: jsii.String("channel"),
//   			VoiceEnhancementMode: jsii.String("voiceEnhancementMode"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html
//
type CfnUserMixinProps struct {
	// After Contact Work configurations of a user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-aftercontactworkconfigs
	//
	AfterContactWorkConfigs interface{} `field:"optional" json:"afterContactWorkConfigs" yaml:"afterContactWorkConfigs"`
	// Auto-accept configurations of a user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-autoacceptconfigs
	//
	AutoAcceptConfigs interface{} `field:"optional" json:"autoAcceptConfigs" yaml:"autoAcceptConfigs"`
	// The identifier of the user account in the directory used for identity management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-directoryuserid
	//
	DirectoryUserId *string `field:"optional" json:"directoryUserId" yaml:"directoryUserId"`
	// The Amazon Resource Name (ARN) of the user's hierarchy group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-hierarchygrouparn
	//
	HierarchyGroupArn *string `field:"optional" json:"hierarchyGroupArn" yaml:"hierarchyGroupArn"`
	// Information about the user identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-identityinfo
	//
	IdentityInfo interface{} `field:"optional" json:"identityInfo" yaml:"identityInfo"`
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The user's password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Persistent Connection configurations of a user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-persistentconnectionconfigs
	//
	PersistentConnectionConfigs interface{} `field:"optional" json:"persistentConnectionConfigs" yaml:"persistentConnectionConfigs"`
	// Information about the phone configuration for the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-phoneconfig
	//
	PhoneConfig interface{} `field:"optional" json:"phoneConfig" yaml:"phoneConfig"`
	// Phone Number configurations of a user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-phonenumberconfigs
	//
	PhoneNumberConfigs interface{} `field:"optional" json:"phoneNumberConfigs" yaml:"phoneNumberConfigs"`
	// The Amazon Resource Name (ARN) of the user's routing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-routingprofilearn
	//
	RoutingProfileArn *string `field:"optional" json:"routingProfileArn" yaml:"routingProfileArn"`
	// The Amazon Resource Name (ARN) of the user's security profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-securityprofilearns
	//
	SecurityProfileArns *[]*string `field:"optional" json:"securityProfileArns" yaml:"securityProfileArns"`
	// The tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The user name assigned to the user account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
	// One or more predefined attributes assigned to a user, with a numeric value that indicates how their level of skill in a specified area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-userproficiencies
	//
	UserProficiencies interface{} `field:"optional" json:"userProficiencies" yaml:"userProficiencies"`
	// Voice Enhancement configurations of a user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html#cfn-connect-user-voiceenhancementconfigs
	//
	VoiceEnhancementConfigs interface{} `field:"optional" json:"voiceEnhancementConfigs" yaml:"voiceEnhancementConfigs"`
}

