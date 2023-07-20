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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-defaultfieldvalue.html
//
type CfnTaskTemplate_DefaultFieldValueProperty struct {
	// Default value for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-defaultfieldvalue.html#cfn-connect-tasktemplate-defaultfieldvalue-defaultvalue
	//
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
	// Identifier of a field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-defaultfieldvalue.html#cfn-connect-tasktemplate-defaultfieldvalue-id
	//
	Id interface{} `field:"required" json:"id" yaml:"id"`
}

