package awsconnect


// Describes a default field and its corresponding value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultFieldValueProperty := &DefaultFieldValueProperty{
//   	DefaultValue: jsii.String("defaultValue"),
//   	Id: &FieldIdentifierProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
type CfnTaskTemplate_DefaultFieldValueProperty struct {
	// Default value for the field.
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
	// Identifier of a field.
	Id interface{} `field:"required" json:"id" yaml:"id"`
}

