package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sensorProperty := &sensorProperty{
//   	dataType: jsii.String("dataType"),
//   	fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   	// the properties below are optional
//   	allowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	description: jsii.String("description"),
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   	unit: jsii.String("unit"),
//   }
//
type CfnSignalCatalog_SensorProperty struct {
	// `CfnSignalCatalog.SensorProperty.DataType`.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// `CfnSignalCatalog.SensorProperty.FullyQualifiedName`.
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// `CfnSignalCatalog.SensorProperty.AllowedValues`.
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// `CfnSignalCatalog.SensorProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnSignalCatalog.SensorProperty.Max`.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// `CfnSignalCatalog.SensorProperty.Min`.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// `CfnSignalCatalog.SensorProperty.Unit`.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

