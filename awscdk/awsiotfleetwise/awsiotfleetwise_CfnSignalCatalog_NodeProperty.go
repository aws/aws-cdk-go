package awsiotfleetwise


// A general abstraction of a signal.
//
// A node can be specified as an actuator, attribute, branch, or sensor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeProperty := &nodeProperty{
//   	actuator: &actuatorProperty{
//   		dataType: jsii.String("dataType"),
//   		fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   		// the properties below are optional
//   		allowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		assignedValue: jsii.String("assignedValue"),
//   		description: jsii.String("description"),
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   		unit: jsii.String("unit"),
//   	},
//   	attribute: &attributeProperty{
//   		dataType: jsii.String("dataType"),
//   		fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   		// the properties below are optional
//   		allowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		assignedValue: jsii.String("assignedValue"),
//   		defaultValue: jsii.String("defaultValue"),
//   		description: jsii.String("description"),
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   		unit: jsii.String("unit"),
//   	},
//   	branch: &branchProperty{
//   		fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   	},
//   	sensor: &sensorProperty{
//   		dataType: jsii.String("dataType"),
//   		fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   		// the properties below are optional
//   		allowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		description: jsii.String("description"),
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   		unit: jsii.String("unit"),
//   	},
//   }
//
type CfnSignalCatalog_NodeProperty struct {
	// Information about a node specified as an actuator.
	//
	// > An actuator is a digital representation of a vehicle device.
	Actuator interface{} `field:"optional" json:"actuator" yaml:"actuator"`
	// Information about a node specified as an attribute.
	//
	// > An attribute represents static information about a vehicle.
	Attribute interface{} `field:"optional" json:"attribute" yaml:"attribute"`
	// Information about a node specified as a branch.
	//
	// > A group of signals that are defined in a hierarchical structure.
	Branch interface{} `field:"optional" json:"branch" yaml:"branch"`
	// `CfnSignalCatalog.NodeProperty.Sensor`.
	Sensor interface{} `field:"optional" json:"sensor" yaml:"sensor"`
}

