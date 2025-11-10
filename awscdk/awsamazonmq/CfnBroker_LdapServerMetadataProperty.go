package awsamazonmq


// Optional.
//
// The metadata of the LDAP server used to authenticate and authorize connections to the broker. Does not apply to RabbitMQ brokers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ldapServerMetadataProperty := &LdapServerMetadataProperty{
//   	Hosts: []*string{
//   		jsii.String("hosts"),
//   	},
//   	RoleBase: jsii.String("roleBase"),
//   	RoleSearchMatching: jsii.String("roleSearchMatching"),
//   	ServiceAccountUsername: jsii.String("serviceAccountUsername"),
//   	UserBase: jsii.String("userBase"),
//   	UserSearchMatching: jsii.String("userSearchMatching"),
//
//   	// the properties below are optional
//   	RoleName: jsii.String("roleName"),
//   	RoleSearchSubtree: jsii.Boolean(false),
//   	ServiceAccountPassword: jsii.String("serviceAccountPassword"),
//   	UserRoleName: jsii.String("userRoleName"),
//   	UserSearchSubtree: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html
//
type CfnBroker_LdapServerMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-hosts
	//
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-rolebase
	//
	RoleBase *string `field:"required" json:"roleBase" yaml:"roleBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-rolesearchmatching
	//
	RoleSearchMatching *string `field:"required" json:"roleSearchMatching" yaml:"roleSearchMatching"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-serviceaccountusername
	//
	ServiceAccountUsername *string `field:"required" json:"serviceAccountUsername" yaml:"serviceAccountUsername"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-userbase
	//
	UserBase *string `field:"required" json:"userBase" yaml:"userBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-usersearchmatching
	//
	UserSearchMatching *string `field:"required" json:"userSearchMatching" yaml:"userSearchMatching"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-rolename
	//
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-rolesearchsubtree
	//
	RoleSearchSubtree interface{} `field:"optional" json:"roleSearchSubtree" yaml:"roleSearchSubtree"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-serviceaccountpassword
	//
	ServiceAccountPassword *string `field:"optional" json:"serviceAccountPassword" yaml:"serviceAccountPassword"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-userrolename
	//
	UserRoleName *string `field:"optional" json:"userRoleName" yaml:"userRoleName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-ldapservermetadata.html#cfn-amazonmq-broker-ldapservermetadata-usersearchsubtree
	//
	UserSearchSubtree interface{} `field:"optional" json:"userSearchSubtree" yaml:"userSearchSubtree"`
}

