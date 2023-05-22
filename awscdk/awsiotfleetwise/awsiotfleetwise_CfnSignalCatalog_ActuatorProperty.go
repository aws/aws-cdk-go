package awsiotfleetwise


// A signal that represents a vehicle device such as the engine, heater, and door locks.
//
// Data from an actuator reports the state of a certain vehicle device.
//
// > Updating actuator data can change the state of a device. For example, you can turn on or off the heater by updating its actuator data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actuatorProperty := &ActuatorProperty{
//   	DataType: jsii.String("dataType"),
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   	// the properties below are optional
//   	AllowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	AssignedValue: jsii.String("assignedValue"),
//   	Description: jsii.String("description"),
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   }
//
type CfnSignalCatalog_ActuatorProperty struct {
	// The specified data type of the actuator.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The fully qualified name of the actuator.
	//
	// For example, the fully qualified name of an actuator might be `Vehicle.Front.Left.Door.Lock` .
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// A list of possible values an actuator can take.
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// A specified value for the actuator.
	AssignedValue *string `field:"optional" json:"assignedValue" yaml:"assignedValue"`
	// A brief description of the actuator.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The specified possible maximum value of an actuator.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The specified possible minimum value of an actuator.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// The scientific unit for the actuator.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

