package awsiotfleetwise


// An input component that reports the environmental condition of a vehicle.
//
// > You can collect data about fluid levels, temperatures, vibrations, or battery voltage from sensors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sensorProperty := &SensorProperty{
//   	DataType: jsii.String("dataType"),
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   	// the properties below are optional
//   	AllowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	Description: jsii.String("description"),
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   }
//
type CfnSignalCatalog_SensorProperty struct {
	// The specified data type of the sensor.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The fully qualified name of the sensor.
	//
	// For example, the fully qualified name of a sensor might be `Vehicle.Body.Engine.Battery` .
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// A list of possible values a sensor can take.
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// A brief description of a sensor.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The specified possible maximum value of the sensor.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The specified possible minimum value of the sensor.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// The scientific unit of measurement for data collected by the sensor.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

