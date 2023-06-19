package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
	// Specifies the Amazon Resource Name (ARN) of the DataSync agent that connects to and reads from your on-premises storage system's management interface.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// Specifies the server name and network port required to connect with the management interface of your on-premises storage system.
	ServerConfiguration interface{} `field:"required" json:"serverConfiguration" yaml:"serverConfiguration"`
	// Specifies the type of on-premises storage system that you want DataSync Discovery to collect information about.
	//
	// > DataSync Discovery currently supports NetApp Fabric-Attached Storage (FAS) and All Flash FAS (AFF) systems running ONTAP 9.7 or later.
	SystemType *string `field:"required" json:"systemType" yaml:"systemType"`
	// Specifies the ARN of the Amazon CloudWatch log group for monitoring and logging discovery job events.
	CloudWatchLogGroupArn *string `field:"optional" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// Specifies a familiar name for your on-premises storage system.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the user name and password for accessing your on-premises storage system's management interface.
	ServerCredentials interface{} `field:"optional" json:"serverCredentials" yaml:"serverCredentials"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least a name tag for your on-premises storage system.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

