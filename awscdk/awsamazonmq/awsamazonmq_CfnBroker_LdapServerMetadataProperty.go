package awsamazonmq


// Optional. The metadata of the LDAP server used to authenticate and authorize connections to the broker.
//
// > Does not apply to RabbitMQ brokers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ldapServerMetadataProperty := &ldapServerMetadataProperty{
//   	hosts: []*string{
//   		jsii.String("hosts"),
//   	},
//   	roleBase: jsii.String("roleBase"),
//   	roleSearchMatching: jsii.String("roleSearchMatching"),
//   	serviceAccountPassword: jsii.String("serviceAccountPassword"),
//   	serviceAccountUsername: jsii.String("serviceAccountUsername"),
//   	userBase: jsii.String("userBase"),
//   	userSearchMatching: jsii.String("userSearchMatching"),
//
//   	// the properties below are optional
//   	roleName: jsii.String("roleName"),
//   	roleSearchSubtree: jsii.Boolean(false),
//   	userRoleName: jsii.String("userRoleName"),
//   	userSearchSubtree: jsii.Boolean(false),
//   }
//
type CfnBroker_LdapServerMetadataProperty struct {
	// Specifies the location of the LDAP server such as AWS Directory Service for Microsoft Active Directory .
	//
	// Optional failover server.
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// The distinguished name of the node in the directory information tree (DIT) to search for roles or groups.
	//
	// For example, `ou=group` , `ou=corp` , `dc=corp` , `dc=example` , `dc=com` .
	RoleBase *string `field:"required" json:"roleBase" yaml:"roleBase"`
	// The LDAP search filter used to find roles within the roleBase.
	//
	// The distinguished name of the user matched by userSearchMatching is substituted into the `{0}` placeholder in the search filter. The client's username is substituted into the `{1}` placeholder. For example, if you set this option to `(member=uid={1})` for the user janedoe, the search filter becomes `(member=uid=janedoe)` after string substitution. It matches all role entries that have a member attribute equal to `uid=janedoe` under the subtree selected by the `RoleBases` .
	RoleSearchMatching *string `field:"required" json:"roleSearchMatching" yaml:"roleSearchMatching"`
	// Service account password.
	//
	// A service account is an account in your LDAP server that has access to initiate a connection. For example, `cn=admin` , `dc=corp` , `dc=example` , `dc=com` .
	ServiceAccountPassword *string `field:"required" json:"serviceAccountPassword" yaml:"serviceAccountPassword"`
	// Service account username.
	//
	// A service account is an account in your LDAP server that has access to initiate a connection. For example, `cn=admin` , `ou=corp` , `dc=corp` , `dc=example` , `dc=com` .
	ServiceAccountUsername *string `field:"required" json:"serviceAccountUsername" yaml:"serviceAccountUsername"`
	// Select a particular subtree of the directory information tree (DIT) to search for user entries.
	//
	// The subtree is specified by a DN, which specifies the base node of the subtree. For example, by setting this option to `ou=Users` , `ou=corp` , `dc=corp` , `dc=example` , `dc=com` , the search for user entries is restricted to the subtree beneath `ou=Users` , `ou=corp` , `dc=corp` , `dc=example` , `dc=com` .
	UserBase *string `field:"required" json:"userBase" yaml:"userBase"`
	// The LDAP search filter used to find users within the `userBase` .
	//
	// The client's username is substituted into the `{0}` placeholder in the search filter. For example, if this option is set to `(uid={0})` and the received username is `janedoe` , the search filter becomes `(uid=janedoe)` after string substitution. It will result in matching an entry like `uid=janedoe` , `ou=Users` , `ou=corp` , `dc=corp` , `dc=example` , `dc=com` .
	UserSearchMatching *string `field:"required" json:"userSearchMatching" yaml:"userSearchMatching"`
	// The group name attribute in a role entry whose value is the name of that role.
	//
	// For example, you can specify `cn` for a group entry's common name. If authentication succeeds, then the user is assigned the the value of the `cn` attribute for each role entry that they are a member of.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// The directory search scope for the role.
	//
	// If set to true, scope is to search the entire subtree.
	RoleSearchSubtree interface{} `field:"optional" json:"roleSearchSubtree" yaml:"roleSearchSubtree"`
	// The name of the LDAP attribute in the user's directory entry for the user's group membership.
	//
	// In some cases, user roles may be identified by the value of an attribute in the user's directory entry. The `UserRoleName` option allows you to provide the name of this attribute.
	UserRoleName *string `field:"optional" json:"userRoleName" yaml:"userRoleName"`
	// The directory search scope for the user.
	//
	// If set to true, scope is to search the entire subtree.
	UserSearchSubtree interface{} `field:"optional" json:"userSearchSubtree" yaml:"userSearchSubtree"`
}

