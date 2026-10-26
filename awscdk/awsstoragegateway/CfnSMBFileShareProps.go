package awsstoragegateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSMBFileShare`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSMBFileShareProps := &CfnSMBFileShareProps{
//   	GatewayArn: jsii.String("gatewayArn"),
//   	LocationArn: jsii.String("locationArn"),
//   	Role: jsii.String("role"),
//
//   	// the properties below are optional
//   	AccessBasedEnumeration: jsii.Boolean(false),
//   	AdminUserList: []*string{
//   		jsii.String("adminUserList"),
//   	},
//   	AuditDestinationArn: jsii.String("auditDestinationArn"),
//   	Authentication: jsii.String("authentication"),
//   	BucketRegion: jsii.String("bucketRegion"),
//   	CacheAttributes: &CacheAttributesProperty{
//   		CacheStaleTimeoutInSeconds: jsii.Number(123),
//   	},
//   	CaseSensitivity: jsii.String("caseSensitivity"),
//   	DefaultStorageClass: jsii.String("defaultStorageClass"),
//   	EncryptionType: jsii.String("encryptionType"),
//   	FileShareName: jsii.String("fileShareName"),
//   	GuessMimeTypeEnabled: jsii.Boolean(false),
//   	InvalidUserList: []*string{
//   		jsii.String("invalidUserList"),
//   	},
//   	KmsEncrypted: jsii.Boolean(false),
//   	KmsKey: jsii.String("kmsKey"),
//   	NotificationPolicy: jsii.String("notificationPolicy"),
//   	ObjectAcl: jsii.String("objectAcl"),
//   	OplocksEnabled: jsii.Boolean(false),
//   	ReadOnly: jsii.Boolean(false),
//   	RequesterPays: jsii.Boolean(false),
//   	SmbaclEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ValidUserList: []*string{
//   		jsii.String("validUserList"),
//   	},
//   	VpcEndpointDnsName: jsii.String("vpcEndpointDnsName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html
//
type CfnSMBFileShareProps struct {
	// The ARN of the S3 File Gateway that hosts the share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-gatewayarn
	//
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// The ARN of the S3 bucket, or bucket prefix, the share exposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-locationarn
	//
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
	// The ARN of the IAM role the gateway assumes when it accesses the S3 location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-role
	//
	Role *string `field:"required" json:"role" yaml:"role"`
	// Whether clients see only the files and folders they may access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-accessbasedenumeration
	//
	AccessBasedEnumeration interface{} `field:"optional" json:"accessBasedEnumeration" yaml:"accessBasedEnumeration"`
	// Users and groups granted administrative privileges on the SMB share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-adminuserlist
	//
	AdminUserList *[]*string `field:"optional" json:"adminUserList" yaml:"adminUserList"`
	// The ARN of the CloudWatch log group that receives the share's audit logs.
	//
	// Removing this property from a template detaches audit logging: the update sends the empty string, which the service treats as the reset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-auditdestinationarn
	//
	AuditDestinationArn *string `field:"optional" json:"auditDestinationArn" yaml:"auditDestinationArn"`
	// The authentication the SMB share uses.
	//
	// The service accepts ActiveDirectory and GuestAccess; it is a length-bounded string in the API model rather than an enumeration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-authentication
	//
	Authentication *string `field:"optional" json:"authentication" yaml:"authentication"`
	// The Region of the S3 bucket, required when the bucket is reached through a VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-bucketregion
	//
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// Refresh cache information for the file share.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-cacheattributes
	//
	CacheAttributes interface{} `field:"optional" json:"cacheAttributes" yaml:"cacheAttributes"`
	// Whether the SMB share is case sensitive or defers to the client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-casesensitivity
	//
	CaseSensitivity *string `field:"optional" json:"caseSensitivity" yaml:"caseSensitivity"`
	// The S3 storage class for new objects.
	//
	// The service accepts S3_STANDARD, S3_INTELLIGENT_TIERING, S3_STANDARD_IA and S3_ONEZONE_IA; it is a length-bounded string in the API model rather than an enumeration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-defaultstorageclass
	//
	DefaultStorageClass *string `field:"optional" json:"defaultStorageClass" yaml:"defaultStorageClass"`
	// The server-side encryption the gateway applies to objects it writes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The name of the file share, which defaults to the bucket name when omitted.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-filesharename
	//
	FileShareName *string `field:"optional" json:"fileShareName" yaml:"fileShareName"`
	// Whether the gateway infers the MIME type of uploaded objects from their extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-guessmimetypeenabled
	//
	GuessMimeTypeEnabled interface{} `field:"optional" json:"guessMimeTypeEnabled" yaml:"guessMimeTypeEnabled"`
	// Users and groups denied access to the SMB share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-invaliduserlist
	//
	InvalidUserList *[]*string `field:"optional" json:"invalidUserList" yaml:"invalidUserList"`
	// Deprecated by the service in favour of EncryptionType;
	//
	// true selects KMS encryption. Retained because it remains readable and settable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-kmsencrypted
	//
	KmsEncrypted interface{} `field:"optional" json:"kmsEncrypted" yaml:"kmsEncrypted"`
	// The KMS key used for encryption, when EncryptionType selects KMS.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The upload-notification settling period, as a JSON document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-notificationpolicy
	//
	NotificationPolicy *string `field:"optional" json:"notificationPolicy" yaml:"notificationPolicy"`
	// The canned S3 ACL applied to objects the gateway writes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-objectacl
	//
	ObjectAcl *string `field:"optional" json:"objectAcl" yaml:"objectAcl"`
	// Whether the SMB share grants opportunistic locks.
	//
	// Removing this property from a template leaves the live value unchanged rather than restoring the created default, because that default depends on the gateway's software version -- opportunistic locking is enabled by default only on appliances at 1.19.0.0 or above -- and the version is not part of this resource. Set the value explicitly rather than relying on removal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-oplocksenabled
	//
	OplocksEnabled interface{} `field:"optional" json:"oplocksEnabled" yaml:"oplocksEnabled"`
	// Whether clients may write to the share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Whether the bucket is a Requester Pays bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-requesterpays
	//
	RequesterPays interface{} `field:"optional" json:"requesterPays" yaml:"requesterPays"`
	// Whether the SMB share uses Windows ACLs rather than POSIX permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-smbaclenabled
	//
	SmbaclEnabled interface{} `field:"optional" json:"smbaclEnabled" yaml:"smbaclEnabled"`
	// Tags assigned to the file share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Users and groups permitted to access the SMB share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-validuserlist
	//
	ValidUserList *[]*string `field:"optional" json:"validUserList" yaml:"validUserList"`
	// The DNS name of the S3 VPC endpoint the gateway uses to reach the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html#cfn-storagegateway-smbfileshare-vpcendpointdnsname
	//
	VpcEndpointDnsName *string `field:"optional" json:"vpcEndpointDnsName" yaml:"vpcEndpointDnsName"`
}

