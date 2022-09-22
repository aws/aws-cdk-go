package awstransfer

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
//   	role: jsii.String("role"),
//   	serverId: jsii.String("serverId"),
//   	userName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	homeDirectory: jsii.String("homeDirectory"),
//   	homeDirectoryMappings: []interface{}{
//   		&homeDirectoryMapEntryProperty{
//   			entry: jsii.String("entry"),
//   			target: jsii.String("target"),
//   		},
//   	},
//   	homeDirectoryType: jsii.String("homeDirectoryType"),
//   	policy: jsii.String("policy"),
//   	posixProfile: &posixProfileProperty{
//   		gid: jsii.Number(123),
//   		uid: jsii.Number(123),
//
//   		// the properties below are optional
//   		secondaryGids: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	sshPublicKeys: []*string{
//   		jsii.String("sshPublicKeys"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnUserProps struct {
	// Specifies the Amazon Resource Name (ARN) of the IAM role that controls your users' access to your Amazon S3 bucket or EFS file system.
	//
	// The policies attached to this role determine the level of access that you want to provide your users when transferring files into and out of your Amazon S3 bucket or EFS file system. The IAM role should also contain a trust relationship that allows the server to access your resources when servicing your users' transfer requests.
	Role *string `field:"required" json:"role" yaml:"role"`
	// A system-assigned unique identifier for a server instance.
	//
	// This is the specific server that you added your user to.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// A unique string that identifies a user and is associated with a `ServerId` .
	//
	// This user name must be a minimum of 3 and a maximum of 100 characters long. The following are valid characters: a-z, A-Z, 0-9, underscore '_', hyphen '-', period '.', and at sign '@'. The user name can't start with a hyphen, period, or at sign.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// The landing directory (folder) for a user when they log in to the server using the client.
	//
	// A `HomeDirectory` example is `/bucket_name/home/mydirectory` .
	HomeDirectory *string `field:"optional" json:"homeDirectory" yaml:"homeDirectory"`
	// Logical directory mappings that specify what Amazon S3 paths and keys should be visible to your user and how you want to make them visible.
	//
	// You will need to specify the " `Entry` " and " `Target` " pair, where `Entry` shows how the path is made visible and `Target` is the actual Amazon S3 path. If you only specify a target, it will be displayed as is. You will need to also make sure that your IAM role provides access to paths in `Target` . The following is an example.
	//
	// `'[ { "Entry": "/", "Target": "/bucket3/customized-reports/" } ]'`
	//
	// In most cases, you can use this value instead of the session policy to lock your user down to the designated home directory ("chroot"). To do this, you can set `Entry` to '/' and set `Target` to the HomeDirectory parameter value.
	//
	// > If the target of a logical directory entry does not exist in Amazon S3, the entry will be ignored. As a workaround, you can use the Amazon S3 API to create 0 byte objects as place holders for your directory. If using the CLI, use the `s3api` call instead of `s3` so you can use the put-object operation. For example, you use the following: `AWS s3api put-object --bucket bucketname --key path/to/folder/` . Make sure that the end of the key name ends in a '/' for it to be considered a folder.
	HomeDirectoryMappings interface{} `field:"optional" json:"homeDirectoryMappings" yaml:"homeDirectoryMappings"`
	// The type of landing directory (folder) you want your users' home directory to be when they log into the server.
	//
	// If you set it to `PATH` , the user will see the absolute Amazon S3 bucket or EFS paths as is in their file transfer protocol clients. If you set it `LOGICAL` , you need to provide mappings in the `HomeDirectoryMappings` for how you want to make Amazon S3 or EFS paths visible to your users.
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
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Specifies the full POSIX identity, including user ID ( `Uid` ), group ID ( `Gid` ), and any secondary groups IDs ( `SecondaryGids` ), that controls your users' access to your Amazon Elastic File System (Amazon EFS) file systems.
	//
	// The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
	PosixProfile interface{} `field:"optional" json:"posixProfile" yaml:"posixProfile"`
	// Specifies the public key portion of the Secure Shell (SSH) keys stored for the described user.
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// Key-value pairs that can be used to group and search for users.
	//
	// Tags are metadata attached to users for any purpose.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

