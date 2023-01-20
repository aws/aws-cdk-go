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
//   cfnLocationObjectStorageProps := &cfnLocationObjectStorageProps{
//   	agentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//
//   	// the properties below are optional
//   	accessKey: jsii.String("accessKey"),
//   	bucketName: jsii.String("bucketName"),
//   	secretKey: jsii.String("secretKey"),
//   	serverHostname: jsii.String("serverHostname"),
//   	serverPort: jsii.Number(123),
//   	serverProtocol: jsii.String("serverProtocol"),
//   	subdirectory: jsii.String("subdirectory"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationObjectStorageProps struct {
	// Specifies the Amazon Resource Names (ARNs) of the agents associated with the location.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// Specifies the access key (or user name) if credentials are required to access the object storage server.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Specifies the name of the bucket that DataSync reads from or writes to.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Specifies the secret key (or password) if credentials are required to access the object storage server.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Specifies the domain name or IP address of the object storage server.
	//
	// A DataSync agent uses this hostname to mount the object storage server.
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// Specifies the port that your object storage server accepts inbound network traffic on.
	//
	// Set to port 80 (HTTP), 443 (HTTPS), or a custom port if needed.
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
	// Specifies the protocol that your object storage server uses to communicate.
	ServerProtocol *string `field:"optional" json:"serverProtocol" yaml:"serverProtocol"`
	// Specifies the object prefix that DataSync reads from or writes to.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies the key-value pair that represents the tag to help you manage, filter, and search for your location.
	//
	// We recommend using tags for naming your locations.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

