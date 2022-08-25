package awslex


// Defines a list of values that Amazon Lex should use as the default value for a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotDefaultValueSpecificationProperty := &slotDefaultValueSpecificationProperty{
//   	defaultValueList: []interface{}{
//   		&slotDefaultValueProperty{
//   			defaultValue: jsii.String("defaultValue"),
//   		},
//   	},
//   }
//
type CfnBot_SlotDefaultValueSpecificationProperty struct {
	// A list of default values.
	//
	// Amazon Lex chooses the default value to use in the order that they are presented in the list.
	DefaultValueList interface{} `field:"required" json:"defaultValueList" yaml:"defaultValueList"`
}

