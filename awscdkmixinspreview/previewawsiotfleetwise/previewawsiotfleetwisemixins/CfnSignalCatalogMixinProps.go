package previewawsiotfleetwisemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSignalCatalogPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSignalCatalogMixinProps := &CfnSignalCatalogMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	NodeCounts: &NodeCountsProperty{
//   		TotalActuators: jsii.Number(123),
//   		TotalAttributes: jsii.Number(123),
//   		TotalBranches: jsii.Number(123),
//   		TotalNodes: jsii.Number(123),
//   		TotalSensors: jsii.Number(123),
//   	},
//   	Nodes: []interface{}{
//   		&NodeProperty{
//   			Actuator: &ActuatorProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				AssignedValue: jsii.String("assignedValue"),
//   				DataType: jsii.String("dataType"),
//   				Description: jsii.String("description"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   			Attribute: &AttributeProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				AssignedValue: jsii.String("assignedValue"),
//   				DataType: jsii.String("dataType"),
//   				DefaultValue: jsii.String("defaultValue"),
//   				Description: jsii.String("description"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   			Branch: &BranchProperty{
//   				Description: jsii.String("description"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   			},
//   			Sensor: &SensorProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DataType: jsii.String("dataType"),
//   				Description: jsii.String("description"),
//   				FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-signalcatalog.html
//
type CfnSignalCatalogMixinProps struct {
	// A brief description of the signal catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-signalcatalog.html#cfn-iotfleetwise-signalcatalog-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the signal catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-signalcatalog.html#cfn-iotfleetwise-signalcatalog-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the number of nodes and node types in a vehicle network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-signalcatalog.html#cfn-iotfleetwise-signalcatalog-nodecounts
	//
	NodeCounts interface{} `field:"optional" json:"nodeCounts" yaml:"nodeCounts"`
	// A list of information about nodes, which are a general abstraction of signals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-signalcatalog.html#cfn-iotfleetwise-signalcatalog-nodes
	//
	Nodes interface{} `field:"optional" json:"nodes" yaml:"nodes"`
	// Metadata that can be used to manage the signal catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-signalcatalog.html#cfn-iotfleetwise-signalcatalog-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

