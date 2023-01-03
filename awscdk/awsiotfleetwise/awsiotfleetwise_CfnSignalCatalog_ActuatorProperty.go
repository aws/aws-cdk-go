package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actuatorProperty := &actuatorProperty{
//   	dataType: jsii.String("dataType"),
//   	fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   	// the properties below are optional
//   	allowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	assignedValue: jsii.String("assignedValue"),
//   	description: jsii.String("description"),
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   	unit: jsii.String("unit"),
//   }
//
type CfnSignalCatalog_ActuatorProperty struct {
	// `CfnSignalCatalog.ActuatorProperty.DataType`.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// `CfnSignalCatalog.ActuatorProperty.FullyQualifiedName`.
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// `CfnSignalCatalog.ActuatorProperty.AllowedValues`.
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// `CfnSignalCatalog.ActuatorProperty.AssignedValue`.
	AssignedValue *string `field:"optional" json:"assignedValue" yaml:"assignedValue"`
	// `CfnSignalCatalog.ActuatorProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnSignalCatalog.ActuatorProperty.Max`.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// `CfnSignalCatalog.ActuatorProperty.Min`.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// `CfnSignalCatalog.ActuatorProperty.Unit`.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

