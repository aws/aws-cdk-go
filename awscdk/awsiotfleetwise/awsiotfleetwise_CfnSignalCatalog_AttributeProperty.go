package awsiotfleetwise


// A signal that represents static information about the vehicle, such as engine type or manufacturing date.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeProperty := &AttributeProperty{
//   	DataType: jsii.String("dataType"),
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//
//   	// the properties below are optional
//   	AllowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	AssignedValue: jsii.String("assignedValue"),
//   	DefaultValue: jsii.String("defaultValue"),
//   	Description: jsii.String("description"),
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   }
//
type CfnSignalCatalog_AttributeProperty struct {
	// The specified data type of the attribute.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The fully qualified name of the attribute.
	//
	// For example, the fully qualified name of an attribute might be `Vehicle.Body.Engine.Type` .
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// A list of possible values an attribute can be assigned.
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// A specified value for the attribute.
	AssignedValue *string `field:"optional" json:"assignedValue" yaml:"assignedValue"`
	// The default value of the attribute.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// A brief description of the attribute.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The specified possible maximum value of the attribute.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The specified possible minimum value of the attribute.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// The scientific unit for the attribute.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

