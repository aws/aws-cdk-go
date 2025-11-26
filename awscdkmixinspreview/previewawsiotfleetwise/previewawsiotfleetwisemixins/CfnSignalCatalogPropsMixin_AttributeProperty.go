package previewawsiotfleetwisemixins


// A signal that represents static information about the vehicle, such as engine type or manufacturing date.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attributeProperty := &AttributeProperty{
//   	AllowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	AssignedValue: jsii.String("assignedValue"),
//   	DataType: jsii.String("dataType"),
//   	DefaultValue: jsii.String("defaultValue"),
//   	Description: jsii.String("description"),
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-attribute.html
//
type CfnSignalCatalogPropsMixin_AttributeProperty struct {
	// A list of possible values an attribute can be assigned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-attribute.html#cfn-iotfleetwise-signalcatalog-attribute-allowedvalues
	//
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// A specified value for the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-attribute.html#cfn-iotfleetwise-signalcatalog-attribute-assignedvalue
	//
	AssignedValue *string `field:"optional" json:"assignedValue" yaml:"assignedValue"`
	// The specified data type of the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-attribute.html#cfn-iotfleetwise-signalcatalog-attribute-datatype
	//
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The default value of the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-attribute.html#cfn-iotfleetwise-signalcatalog-attribute-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// A brief description of the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-attribute.html#cfn-iotfleetwise-signalcatalog-attribute-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The fully qualified name of the attribute.
	//
	// For example, the fully qualified name of an attribute might be `Vehicle.Body.Engine.Type` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-attribute.html#cfn-iotfleetwise-signalcatalog-attribute-fullyqualifiedname
	//
	FullyQualifiedName *string `field:"optional" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// The specified possible maximum value of the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-attribute.html#cfn-iotfleetwise-signalcatalog-attribute-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The specified possible minimum value of the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-attribute.html#cfn-iotfleetwise-signalcatalog-attribute-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// The scientific unit for the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-attribute.html#cfn-iotfleetwise-signalcatalog-attribute-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

