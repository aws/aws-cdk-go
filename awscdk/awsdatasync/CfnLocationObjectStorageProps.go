package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocationObjectStorage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationObjectStorageProps := &CfnLocationObjectStorageProps{
//   	AccessKey: jsii.String("accessKey"),
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	BucketName: jsii.String("bucketName"),
//   	CmkSecretConfig: &CmkSecretConfigProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	CustomSecretConfig: &CustomSecretConfigProperty{
//   		SecretAccessRoleArn: jsii.String("secretAccessRoleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	SecretKey: jsii.String("secretKey"),
//   	ServerCertificate: jsii.String("serverCertificate"),
//   	ServerHostname: jsii.String("serverHostname"),
//   	ServerPort: jsii.Number(123),
//   	ServerProtocol: jsii.String("serverProtocol"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html
//
type CfnLocationObjectStorageProps struct {
	// Specifies the access key (for example, a user name) if credentials are required to authenticate with the object storage server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-accesskey
	//
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// (Optional) Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can connect with your object storage system.
	//
	// If you are setting up an agentless cross-cloud transfer, you do not need to specify a value for this parameter.
	//
	// > Make sure you configure this parameter correctly when you first create your storage location. You cannot add or remove agents from a storage location after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-agentarns
	//
	AgentArns *[]*string `field:"optional" json:"agentArns" yaml:"agentArns"`
	// Specifies the name of the object storage bucket involved in the transfer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Specifies configuration information for a DataSync-managed secret, which includes the `SecretKey` that DataSync uses to access a specific object storage location, with a customer-managed AWS KMS key .
	//
	// When you include this paramater as part of a `CreateLocationObjectStorage` request, you provide only the KMS key ARN. DataSync uses this KMS key together with the value you specify for the `SecretKey` parameter to create a DataSync-managed secret to store the location access credentials.
	//
	// Make sure the DataSync has permission to access the KMS key that you specify.
	//
	// > You can use either `CmkSecretConfig` (with `SecretKey` ) or `CustomSecretConfig` (without `SecretKey` ) to provide credentials for a `CreateLocationObjectStorage` request. Do not provide both parameters for the same request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-cmksecretconfig
	//
	CmkSecretConfig interface{} `field:"optional" json:"cmkSecretConfig" yaml:"cmkSecretConfig"`
	// Specifies configuration information for a customer-managed Secrets Manager secret where the secret key for a specific object storage location is stored in plain text.
	//
	// This configuration includes the secret ARN, and the ARN for an IAM role that provides access to the secret.
	//
	// > You can use either `CmkSecretConfig` (with `SecretKey` ) or `CustomSecretConfig` (without `SecretKey` ) to provide credentials for a `CreateLocationObjectStorage` request. Do not provide both parameters for the same request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-customsecretconfig
	//
	CustomSecretConfig interface{} `field:"optional" json:"customSecretConfig" yaml:"customSecretConfig"`
	// Specifies the secret key (for example, a password) if credentials are required to authenticate with the object storage server.
	//
	// > If you provide a secret using `SecretKey` , but do not provide secret configuration details using `CmkSecretConfig` or `CustomSecretConfig` , then DataSync stores the token using your AWS account's Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-secretkey
	//
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Specifies a certificate chain for DataSync to authenticate with your object storage system if the system uses a private or self-signed certificate authority (CA).
	//
	// You must specify a single `.pem` file with a full certificate chain (for example, `file:///home/user/.ssh/object_storage_certificates.pem` ).
	//
	// The certificate chain might include:
	//
	// - The object storage system's certificate
	// - All intermediate certificates (if there are any)
	// - The root certificate of the signing CA
	//
	// You can concatenate your certificates into a `.pem` file (which can be up to 32768 bytes before base64 encoding). The following example `cat` command creates an `object_storage_certificates.pem` file that includes three certificates:
	//
	// `cat object_server_certificate.pem intermediate_certificate.pem ca_root_certificate.pem > object_storage_certificates.pem`
	//
	// To use this parameter, configure `ServerProtocol` to `HTTPS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-servercertificate
	//
	ServerCertificate *string `field:"optional" json:"serverCertificate" yaml:"serverCertificate"`
	// Specifies the domain name or IP address (IPv4 or IPv6) of the object storage server that your DataSync agent connects to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-serverhostname
	//
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// Specifies the port that your object storage server accepts inbound network traffic on (for example, port 443).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-serverport
	//
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
	// Specifies the protocol that your object storage server uses to communicate.
	//
	// If not specified, the default value is `HTTPS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-serverprotocol
	//
	ServerProtocol *string `field:"optional" json:"serverProtocol" yaml:"serverProtocol"`
	// Specifies the object prefix for your object storage server.
	//
	// If this is a source location, DataSync only copies objects with this prefix. If this is a destination location, DataSync writes all objects with this prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies the key-value pair that represents a tag that you want to add to the resource.
	//
	// Tags can help you manage, filter, and search for your resources. We recommend creating a name tag for your location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html#cfn-datasync-locationobjectstorage-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

