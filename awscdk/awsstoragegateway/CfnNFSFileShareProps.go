package awsstoragegateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnNFSFileShare`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNFSFileShareProps := &CfnNFSFileShareProps{
//   	GatewayArn: jsii.String("gatewayArn"),
//   	LocationArn: jsii.String("locationArn"),
//   	Role: jsii.String("role"),
//
//   	// the properties below are optional
//   	AuditDestinationArn: jsii.String("auditDestinationArn"),
//   	BucketRegion: jsii.String("bucketRegion"),
//   	CacheAttributes: &CacheAttributesProperty{
//   		CacheStaleTimeoutInSeconds: jsii.Number(123),
//   	},
//   	ClientList: []*string{
//   		jsii.String("clientList"),
//   	},
//   	DefaultStorageClass: jsii.String("defaultStorageClass"),
//   	EncryptionType: jsii.String("encryptionType"),
//   	FileShareName: jsii.String("fileShareName"),
//   	GuessMimeTypeEnabled: jsii.Boolean(false),
//   	KmsEncrypted: jsii.Boolean(false),
//   	KmsKey: jsii.String("kmsKey"),
//   	NfsFileShareDefaults: &NFSFileShareDefaultsProperty{
//   		DirectoryMode: jsii.String("directoryMode"),
//   		FileMode: jsii.String("fileMode"),
//   		GroupId: jsii.Number(123),
//   		OwnerId: jsii.Number(123),
//   	},
//   	NotificationPolicy: jsii.String("notificationPolicy"),
//   	ObjectAcl: jsii.String("objectAcl"),
//   	ReadOnly: jsii.Boolean(false),
//   	RequesterPays: jsii.Boolean(false),
//   	Squash: jsii.String("squash"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcEndpointDnsName: jsii.String("vpcEndpointDnsName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html
//
type CfnNFSFileShareProps struct {
	// The ARN of the S3 File Gateway that hosts the share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-gatewayarn
	//
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// The ARN of the S3 bucket, or bucket prefix, the share exposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-locationarn
	//
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
	// The ARN of the IAM role the gateway assumes when it accesses the S3 location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-role
	//
	Role *string `field:"required" json:"role" yaml:"role"`
	// The ARN of the CloudWatch log group that receives the share's audit logs.
	//
	// Removing this property from a template detaches audit logging: the update sends the empty string, which the service treats as the reset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-auditdestinationarn
	//
	AuditDestinationArn *string `field:"optional" json:"auditDestinationArn" yaml:"auditDestinationArn"`
	// The Region of the S3 bucket, required when the bucket is reached through a VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-bucketregion
	//
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// Refresh cache information for the file share.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-cacheattributes
	//
	CacheAttributes interface{} `field:"optional" json:"cacheAttributes" yaml:"cacheAttributes"`
	// The IPv4 or IPv6 addresses and CIDR blocks permitted to mount the NFS share.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-clientlist
	//
	ClientList *[]*string `field:"optional" json:"clientList" yaml:"clientList"`
	// The S3 storage class for new objects.
	//
	// The service accepts S3_STANDARD, S3_INTELLIGENT_TIERING, S3_STANDARD_IA and S3_ONEZONE_IA; it is a length-bounded string in the API model rather than an enumeration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-defaultstorageclass
	//
	DefaultStorageClass *string `field:"optional" json:"defaultStorageClass" yaml:"defaultStorageClass"`
	// The server-side encryption the gateway applies to objects it writes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The name of the file share, which defaults to the bucket name when omitted.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-filesharename
	//
	FileShareName *string `field:"optional" json:"fileShareName" yaml:"fileShareName"`
	// Whether the gateway infers the MIME type of uploaded objects from their extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-guessmimetypeenabled
	//
	GuessMimeTypeEnabled interface{} `field:"optional" json:"guessMimeTypeEnabled" yaml:"guessMimeTypeEnabled"`
	// Deprecated by the service in favour of EncryptionType;
	//
	// true selects KMS encryption. Retained because it remains readable and settable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-kmsencrypted
	//
	KmsEncrypted interface{} `field:"optional" json:"kmsEncrypted" yaml:"kmsEncrypted"`
	// The KMS key used for encryption, when EncryptionType selects KMS.
	//
	// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Default values for the file and directory permissions of an NFS share.
	//
	// Removing this property from a template leaves the live values unchanged rather than clearing them, because the service has no reset for it: a create that omits the object leaves it absent, but the update path merges member by member keyed on the same internal sentinel that reports absence, so no request can restore that state. Set the members you want rather than relying on removal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-nfsfilesharedefaults
	//
	NfsFileShareDefaults interface{} `field:"optional" json:"nfsFileShareDefaults" yaml:"nfsFileShareDefaults"`
	// The upload-notification settling period, as a JSON document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-notificationpolicy
	//
	NotificationPolicy *string `field:"optional" json:"notificationPolicy" yaml:"notificationPolicy"`
	// The canned S3 ACL applied to objects the gateway writes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-objectacl
	//
	ObjectAcl *string `field:"optional" json:"objectAcl" yaml:"objectAcl"`
	// Whether clients may write to the share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Whether the bucket is a Requester Pays bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-requesterpays
	//
	RequesterPays interface{} `field:"optional" json:"requesterPays" yaml:"requesterPays"`
	// How the NFS share maps client user IDs.
	//
	// The service accepts RootSquash, NoSquash and AllSquash; it is a length-bounded string in the API model rather than an enumeration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-squash
	//
	Squash *string `field:"optional" json:"squash" yaml:"squash"`
	// Tags assigned to the file share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The DNS name of the S3 VPC endpoint the gateway uses to reach the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html#cfn-storagegateway-nfsfileshare-vpcendpointdnsname
	//
	VpcEndpointDnsName *string `field:"optional" json:"vpcEndpointDnsName" yaml:"vpcEndpointDnsName"`
}

