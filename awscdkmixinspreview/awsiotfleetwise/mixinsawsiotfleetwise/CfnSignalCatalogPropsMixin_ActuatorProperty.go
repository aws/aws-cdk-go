package mixinsawsiotfleetwise


// A signal that represents a vehicle device such as the engine, heater, and door locks.
//
// Data from an actuator reports the state of a certain vehicle device.
//
// > Updating actuator data can change the state of a device. For example, you can turn on or off the heater by updating its actuator data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actuatorProperty := &ActuatorProperty{
//   	AllowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	AssignedValue: jsii.String("assignedValue"),
//   	DataType: jsii.String("dataType"),
//   	Description: jsii.String("description"),
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html
//
type CfnSignalCatalogPropsMixin_ActuatorProperty struct {
	// A list of possible values an actuator can take.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-allowedvalues
	//
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// A specified value for the actuator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-assignedvalue
	//
	AssignedValue *string `field:"optional" json:"assignedValue" yaml:"assignedValue"`
	// The specified data type of the actuator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-datatype
	//
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// A brief description of the actuator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The fully qualified name of the actuator.
	//
	// For example, the fully qualified name of an actuator might be `Vehicle.Front.Left.Door.Lock` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-fullyqualifiedname
	//
	FullyQualifiedName *string `field:"optional" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// The specified possible maximum value of an actuator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The specified possible minimum value of an actuator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// The scientific unit for the actuator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

