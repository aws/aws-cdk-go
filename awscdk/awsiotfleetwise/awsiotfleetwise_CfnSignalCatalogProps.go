package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSignalCatalog`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSignalCatalogProps := &cfnSignalCatalogProps{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	nodes: []interface{}{
//   		&nodeProperty{
//   			actuator: &actuatorProperty{
//   				dataType: jsii.String("dataType"),
//   				fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   				// the properties below are optional
//   				allowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				assignedValue: jsii.String("assignedValue"),
//   				description: jsii.String("description"),
//   				max: jsii.Number(123),
//   				min: jsii.Number(123),
//   				unit: jsii.String("unit"),
//   			},
//   			attribute: &attributeProperty{
//   				dataType: jsii.String("dataType"),
//   				fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   				// the properties below are optional
//   				allowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				assignedValue: jsii.String("assignedValue"),
//   				defaultValue: jsii.String("defaultValue"),
//   				description: jsii.String("description"),
//   				max: jsii.Number(123),
//   				min: jsii.Number(123),
//   				unit: jsii.String("unit"),
//   			},
//   			branch: &branchProperty{
//   				fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   				// the properties below are optional
//   				description: jsii.String("description"),
//   			},
//   			sensor: &sensorProperty{
//   				dataType: jsii.String("dataType"),
//   				fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   				// the properties below are optional
//   				allowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				description: jsii.String("description"),
//   				max: jsii.Number(123),
//   				min: jsii.Number(123),
//   				unit: jsii.String("unit"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSignalCatalogProps struct {
	// `AWS::IoTFleetWise::SignalCatalog.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTFleetWise::SignalCatalog.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::IoTFleetWise::SignalCatalog.Nodes`.
	Nodes interface{} `field:"optional" json:"nodes" yaml:"nodes"`
	// `AWS::IoTFleetWise::SignalCatalog.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

