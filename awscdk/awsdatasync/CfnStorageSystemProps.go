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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-storagesystem.html
//
type CfnStorageSystemProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-storagesystem.html#cfn-datasync-storagesystem-agentarns
	//
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-storagesystem.html#cfn-datasync-storagesystem-serverconfiguration
	//
	ServerConfiguration interface{} `field:"required" json:"serverConfiguration" yaml:"serverConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-storagesystem.html#cfn-datasync-storagesystem-systemtype
	//
	SystemType *string `field:"required" json:"systemType" yaml:"systemType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-storagesystem.html#cfn-datasync-storagesystem-cloudwatchloggrouparn
	//
	CloudWatchLogGroupArn *string `field:"optional" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-storagesystem.html#cfn-datasync-storagesystem-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-storagesystem.html#cfn-datasync-storagesystem-servercredentials
	//
	ServerCredentials interface{} `field:"optional" json:"serverCredentials" yaml:"serverCredentials"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-storagesystem.html#cfn-datasync-storagesystem-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

