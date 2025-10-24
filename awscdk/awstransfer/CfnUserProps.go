package awstransfer

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
//   	Role: jsii.String("role"),
//   	ServerId: jsii.String("serverId"),
//   	UserName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	HomeDirectory: jsii.String("homeDirectory"),
//   	HomeDirectoryMappings: []interface{}{
//   		&HomeDirectoryMapEntryProperty{
//   			Entry: jsii.String("entry"),
//   			Target: jsii.String("target"),
//
//   			// the properties below are optional
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	HomeDirectoryType: jsii.String("homeDirectoryType"),
//   	Policy: jsii.String("policy"),
//   	PosixProfile: &PosixProfileProperty{
//   		Gid: jsii.Number(123),
//   		Uid: jsii.Number(123),
//
//   		// the properties below are optional
//   		SecondaryGids: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	SshPublicKeys: []*string{
//   		jsii.String("sshPublicKeys"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html
//
type CfnUserProps struct {
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that controls your users' access to your Amazon S3 bucket or Amazon EFS file system.
	//
	// The policies attached to this role determine the level of access that you want to provide your users when transferring files into and out of your Amazon S3 bucket or Amazon EFS file system. The IAM role should also contain a trust relationship that allows the server to access your resources when servicing your users' transfer requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-role
	//
	Role *string `field:"required" json:"role" yaml:"role"`
	// A system-assigned unique identifier for a server instance.
	//
	// This is the specific server that you added your user to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-serverid
	//
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// A unique string that identifies a user and is associated with a `ServerId` .
	//
	// This user name must be a minimum of 3 and a maximum of 100 characters long. The following are valid characters: a-z, A-Z, 0-9, underscore '_', hyphen '-', period '.', and at sign '@'. The user name can't start with a hyphen, period, or at sign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-username
	//
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// The landing directory (folder) for a user when they log in to the server using the client.
	//
	// A `HomeDirectory` example is `/bucket_name/home/mydirectory` .
	//
	// > You can use the `HomeDirectory` parameter for `HomeDirectoryType` when it is set to either `PATH` or `LOGICAL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-homedirectory
	//
	HomeDirectory *string `field:"optional" json:"homeDirectory" yaml:"homeDirectory"`
	// Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and keys should be visible to your user and how you want to make them visible.
	//
	// You must specify the `Entry` and `Target` pair, where `Entry` shows how the path is made visible and `Target` is the actual Amazon S3 or Amazon EFS path. If you only specify a target, it is displayed as is. You also must ensure that your AWS Identity and Access Management (IAM) role provides access to paths in `Target` . This value can be set only when `HomeDirectoryType` is set to *LOGICAL* .
	//
	// The following is an `Entry` and `Target` pair example.
	//
	// `[ { "Entry": "/directory1", "Target": "/bucket_name/home/mydirectory" } ]`
	//
	// In most cases, you can use this value instead of the session policy to lock your user down to the designated home directory (" `chroot` "). To do this, you can set `Entry` to `/` and set `Target` to the value the user should see for their home directory when they log in.
	//
	// The following is an `Entry` and `Target` pair example for `chroot` .
	//
	// `[ { "Entry": "/", "Target": "/bucket_name/home/mydirectory" } ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-homedirectorymappings
	//
	HomeDirectoryMappings interface{} `field:"optional" json:"homeDirectoryMappings" yaml:"homeDirectoryMappings"`
	// The type of landing directory (folder) that you want your users' home directory to be when they log in to the server.
	//
	// If you set it to `PATH` , the user will see the absolute Amazon S3 bucket or Amazon EFS path as is in their file transfer protocol clients. If you set it to `LOGICAL` , you need to provide mappings in the `HomeDirectoryMappings` for how you want to make Amazon S3 or Amazon EFS paths visible to your users.
	//
	// > If `HomeDirectoryType` is `LOGICAL` , you must provide mappings, using the `HomeDirectoryMappings` parameter. If, on the other hand, `HomeDirectoryType` is `PATH` , you provide an absolute path using the `HomeDirectory` parameter. You cannot have both `HomeDirectory` and `HomeDirectoryMappings` in your template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-homedirectorytype
	//
	HomeDirectoryType *string `field:"optional" json:"homeDirectoryType" yaml:"homeDirectoryType"`
	// A session policy for your user so you can use the same IAM role across multiple users.
	//
	// This policy restricts user access to portions of their Amazon S3 bucket. Variables that you can use inside this policy include `${Transfer:UserName}` , `${Transfer:HomeDirectory}` , and `${Transfer:HomeBucket}` .
	//
	// > For session policies, AWS Transfer Family stores the policy as a JSON blob, instead of the Amazon Resource Name (ARN) of the policy. You save the policy as a JSON blob and pass it in the `Policy` argument.
	// >
	// > For an example of a session policy, see [Example session policy](https://docs.aws.amazon.com/transfer/latest/userguide/session-policy.html) .
	// >
	// > For more information, see [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) in the *AWS Security Token Service API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-policy
	//
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Specifies the full POSIX identity, including user ID ( `Uid` ), group ID ( `Gid` ), and any secondary groups IDs ( `SecondaryGids` ), that controls your users' access to your Amazon Elastic File System (Amazon EFS) file systems.
	//
	// The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-posixprofile
	//
	PosixProfile interface{} `field:"optional" json:"posixProfile" yaml:"posixProfile"`
	// Specifies the public key portion of the Secure Shell (SSH) keys stored for the described user.
	//
	// > To delete the public key body, set its value to zero keys, as shown here:
	// >
	// > `SshPublicKeys: []`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-sshpublickeys
	//
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// Key-value pairs that can be used to group and search for users.
	//
	// Tags are metadata attached to users for any purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

