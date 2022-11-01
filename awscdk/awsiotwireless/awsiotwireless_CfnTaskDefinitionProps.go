package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTaskDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskDefinitionProps := &cfnTaskDefinitionProps{
//   	autoCreateTasks: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	loRaWanUpdateGatewayTaskEntry: &loRaWANUpdateGatewayTaskEntryProperty{
//   		currentVersion: &loRaWANGatewayVersionProperty{
//   			model: jsii.String("model"),
//   			packageVersion: jsii.String("packageVersion"),
//   			station: jsii.String("station"),
//   		},
//   		updateVersion: &loRaWANGatewayVersionProperty{
//   			model: jsii.String("model"),
//   			packageVersion: jsii.String("packageVersion"),
//   			station: jsii.String("station"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskDefinitionType: jsii.String("taskDefinitionType"),
//   	update: &updateWirelessGatewayTaskCreateProperty{
//   		loRaWan: &loRaWANUpdateGatewayTaskCreateProperty{
//   			currentVersion: &loRaWANGatewayVersionProperty{
//   				model: jsii.String("model"),
//   				packageVersion: jsii.String("packageVersion"),
//   				station: jsii.String("station"),
//   			},
//   			sigKeyCrc: jsii.Number(123),
//   			updateSignature: jsii.String("updateSignature"),
//   			updateVersion: &loRaWANGatewayVersionProperty{
//   				model: jsii.String("model"),
//   				packageVersion: jsii.String("packageVersion"),
//   				station: jsii.String("station"),
//   			},
//   		},
//   		updateDataRole: jsii.String("updateDataRole"),
//   		updateDataSource: jsii.String("updateDataSource"),
//   	},
//   }
//
type CfnTaskDefinitionProps struct {
	// Whether to automatically create tasks using this task definition for all gateways with the specified current version.
	//
	// If `false` , the task must me created by calling `CreateWirelessGatewayTask` .
	AutoCreateTasks interface{} `field:"required" json:"autoCreateTasks" yaml:"autoCreateTasks"`
	// `AWS::IoTWireless::TaskDefinition.LoRaWANUpdateGatewayTaskEntry`.
	LoRaWanUpdateGatewayTaskEntry interface{} `field:"optional" json:"loRaWanUpdateGatewayTaskEntry" yaml:"loRaWanUpdateGatewayTaskEntry"`
	// The name of the new resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::IoTWireless::TaskDefinition.TaskDefinitionType`.
	TaskDefinitionType *string `field:"optional" json:"taskDefinitionType" yaml:"taskDefinitionType"`
	// Information about the gateways to update.
	Update interface{} `field:"optional" json:"update" yaml:"update"`
}

