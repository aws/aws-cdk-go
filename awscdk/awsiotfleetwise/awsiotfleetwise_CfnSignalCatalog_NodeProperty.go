package awsiotfleetwise


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
	// `CfnSignalCatalog.NodeProperty.Actuator`.
	Actuator interface{} `field:"optional" json:"actuator" yaml:"actuator"`
	// `CfnSignalCatalog.NodeProperty.Attribute`.
	Attribute interface{} `field:"optional" json:"attribute" yaml:"attribute"`
	// `CfnSignalCatalog.NodeProperty.Branch`.
	Branch interface{} `field:"optional" json:"branch" yaml:"branch"`
	// `CfnSignalCatalog.NodeProperty.Sensor`.
	Sensor interface{} `field:"optional" json:"sensor" yaml:"sensor"`
}

