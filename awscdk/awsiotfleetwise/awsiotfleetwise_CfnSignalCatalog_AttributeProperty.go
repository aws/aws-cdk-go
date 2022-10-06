package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeProperty := &attributeProperty{
//   	dataType: jsii.String("dataType"),
//   	fullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   	// the properties below are optional
//   	allowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	assignedValue: jsii.String("assignedValue"),
//   	defaultValue: jsii.String("defaultValue"),
//   	description: jsii.String("description"),
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   	unit: jsii.String("unit"),
//   }
//
type CfnSignalCatalog_AttributeProperty struct {
	// `CfnSignalCatalog.AttributeProperty.DataType`.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// `CfnSignalCatalog.AttributeProperty.FullyQualifiedName`.
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// `CfnSignalCatalog.AttributeProperty.AllowedValues`.
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// `CfnSignalCatalog.AttributeProperty.AssignedValue`.
	AssignedValue *string `field:"optional" json:"assignedValue" yaml:"assignedValue"`
	// `CfnSignalCatalog.AttributeProperty.DefaultValue`.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// `CfnSignalCatalog.AttributeProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnSignalCatalog.AttributeProperty.Max`.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// `CfnSignalCatalog.AttributeProperty.Min`.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// `CfnSignalCatalog.AttributeProperty.Unit`.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

