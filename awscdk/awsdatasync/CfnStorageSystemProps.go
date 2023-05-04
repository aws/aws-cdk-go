package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnStorageSystem`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStorageSystemProps := &CfnStorageSystemProps{
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	ServerConfiguration: &ServerConfigurationProperty{
//   		ServerHostname: jsii.String("serverHostname"),
//
//   		// the properties below are optional
//   		ServerPort: jsii.Number(123),
//   	},
//   	SystemType: jsii.String("systemType"),
//
//   	// the properties below are optional
//   	CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	Name: jsii.String("name"),
//   	ServerCredentials: &ServerCredentialsProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStorageSystemProps struct {
	// `AWS::DataSync::StorageSystem.AgentArns`.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// `AWS::DataSync::StorageSystem.ServerConfiguration`.
	ServerConfiguration interface{} `field:"required" json:"serverConfiguration" yaml:"serverConfiguration"`
	// `AWS::DataSync::StorageSystem.SystemType`.
	SystemType *string `field:"required" json:"systemType" yaml:"systemType"`
	// `AWS::DataSync::StorageSystem.CloudWatchLogGroupArn`.
	CloudWatchLogGroupArn *string `field:"optional" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// `AWS::DataSync::StorageSystem.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::DataSync::StorageSystem.ServerCredentials`.
	ServerCredentials interface{} `field:"optional" json:"serverCredentials" yaml:"serverCredentials"`
	// `AWS::DataSync::StorageSystem.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

