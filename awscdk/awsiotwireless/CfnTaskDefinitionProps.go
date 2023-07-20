package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTaskDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskDefinitionProps := &CfnTaskDefinitionProps{
//   	AutoCreateTasks: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	LoRaWanUpdateGatewayTaskEntry: &LoRaWANUpdateGatewayTaskEntryProperty{
//   		CurrentVersion: &LoRaWANGatewayVersionProperty{
//   			Model: jsii.String("model"),
//   			PackageVersion: jsii.String("packageVersion"),
//   			Station: jsii.String("station"),
//   		},
//   		UpdateVersion: &LoRaWANGatewayVersionProperty{
//   			Model: jsii.String("model"),
//   			PackageVersion: jsii.String("packageVersion"),
//   			Station: jsii.String("station"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskDefinitionType: jsii.String("taskDefinitionType"),
//   	Update: &UpdateWirelessGatewayTaskCreateProperty{
//   		LoRaWan: &LoRaWANUpdateGatewayTaskCreateProperty{
//   			CurrentVersion: &LoRaWANGatewayVersionProperty{
//   				Model: jsii.String("model"),
//   				PackageVersion: jsii.String("packageVersion"),
//   				Station: jsii.String("station"),
//   			},
//   			SigKeyCrc: jsii.Number(123),
//   			UpdateSignature: jsii.String("updateSignature"),
//   			UpdateVersion: &LoRaWANGatewayVersionProperty{
//   				Model: jsii.String("model"),
//   				PackageVersion: jsii.String("packageVersion"),
//   				Station: jsii.String("station"),
//   			},
//   		},
//   		UpdateDataRole: jsii.String("updateDataRole"),
//   		UpdateDataSource: jsii.String("updateDataSource"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-taskdefinition.html
//
type CfnTaskDefinitionProps struct {
	// Whether to automatically create tasks using this task definition for all gateways with the specified current version.
	//
	// If `false` , the task must be created by calling `CreateWirelessGatewayTask` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-taskdefinition.html#cfn-iotwireless-taskdefinition-autocreatetasks
	//
	AutoCreateTasks interface{} `field:"required" json:"autoCreateTasks" yaml:"autoCreateTasks"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-taskdefinition.html#cfn-iotwireless-taskdefinition-lorawanupdategatewaytaskentry
	//
	LoRaWanUpdateGatewayTaskEntry interface{} `field:"optional" json:"loRaWanUpdateGatewayTaskEntry" yaml:"loRaWanUpdateGatewayTaskEntry"`
	// The name of the new resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-taskdefinition.html#cfn-iotwireless-taskdefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-taskdefinition.html#cfn-iotwireless-taskdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A filter to list only the wireless gateway task definitions that use this task definition type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-taskdefinition.html#cfn-iotwireless-taskdefinition-taskdefinitiontype
	//
	TaskDefinitionType *string `field:"optional" json:"taskDefinitionType" yaml:"taskDefinitionType"`
	// Information about the gateways to update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-taskdefinition.html#cfn-iotwireless-taskdefinition-update
	//
	Update interface{} `field:"optional" json:"update" yaml:"update"`
}

