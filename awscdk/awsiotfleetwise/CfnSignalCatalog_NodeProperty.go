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
//   nodeProperty := &NodeProperty{
//   	Actuator: &ActuatorProperty{
//   		DataType: jsii.String("dataType"),
//   		FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   		// the properties below are optional
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		AssignedValue: jsii.String("assignedValue"),
//   		Description: jsii.String("description"),
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   		Unit: jsii.String("unit"),
//   	},
//   	Attribute: &AttributeProperty{
//   		DataType: jsii.String("dataType"),
//   		FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   		// the properties below are optional
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		AssignedValue: jsii.String("assignedValue"),
//   		DefaultValue: jsii.String("defaultValue"),
//   		Description: jsii.String("description"),
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   		Unit: jsii.String("unit"),
//   	},
//   	Branch: &BranchProperty{
//   		FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   	},
//   	Sensor: &SensorProperty{
//   		DataType: jsii.String("dataType"),
//   		FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   		// the properties below are optional
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		Description: jsii.String("description"),
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   		Unit: jsii.String("unit"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-node.html
//
type CfnSignalCatalog_NodeProperty struct {
	// Information about a node specified as an actuator.
	//
	// > An actuator is a digital representation of a vehicle device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-node.html#cfn-iotfleetwise-signalcatalog-node-actuator
	//
	Actuator interface{} `field:"optional" json:"actuator" yaml:"actuator"`
	// Information about a node specified as an attribute.
	//
	// > An attribute represents static information about a vehicle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-node.html#cfn-iotfleetwise-signalcatalog-node-attribute
	//
	Attribute interface{} `field:"optional" json:"attribute" yaml:"attribute"`
	// Information about a node specified as a branch.
	//
	// > A group of signals that are defined in a hierarchical structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-node.html#cfn-iotfleetwise-signalcatalog-node-branch
	//
	Branch interface{} `field:"optional" json:"branch" yaml:"branch"`
	// An input component that reports the environmental condition of a vehicle.
	//
	// > You can collect data about fluid levels, temperatures, vibrations, or battery voltage from sensors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-node.html#cfn-iotfleetwise-signalcatalog-node-sensor
	//
	Sensor interface{} `field:"optional" json:"sensor" yaml:"sensor"`
}

