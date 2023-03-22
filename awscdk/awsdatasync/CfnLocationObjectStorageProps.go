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
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//
//   	// the properties below are optional
//   	AccessKey: jsii.String("accessKey"),
//   	BucketName: jsii.String("bucketName"),
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
type CfnLocationObjectStorageProps struct {
	// Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can securely connect with your location.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// Specifies the access key (for example, a user name) if credentials are required to authenticate with the object storage server.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Specifies the name of the object storage bucket involved in the transfer.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Specifies the secret key (for example, a password) if credentials are required to authenticate with the object storage server.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Specifies a certificate to authenticate with an object storage system that uses a private or self-signed certificate authority (CA).
	//
	// You must specify a Base64-encoded `.pem` file (for example, `file:///home/user/.ssh/storage_sys_certificate.pem` ). The certificate can be up to 32768 bytes (before Base64 encoding).
	//
	// To use this parameter, configure `ServerProtocol` to `HTTPS` .
	ServerCertificate *string `field:"optional" json:"serverCertificate" yaml:"serverCertificate"`
	// Specifies the domain name or IP address of the object storage server.
	//
	// A DataSync agent uses this hostname to mount the object storage server in a network.
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// Specifies the port that your object storage server accepts inbound network traffic on (for example, port 443).
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
	// Specifies the protocol that your object storage server uses to communicate.
	ServerProtocol *string `field:"optional" json:"serverProtocol" yaml:"serverProtocol"`
	// Specifies the object prefix for your object storage server.
	//
	// If this is a source location, DataSync only copies objects with this prefix. If this is a destination location, DataSync writes all objects with this prefix.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies the key-value pair that represents a tag that you want to add to the resource.
	//
	// Tags can help you manage, filter, and search for your resources. We recommend creating a name tag for your location.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

