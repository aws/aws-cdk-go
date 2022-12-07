package awsconnect


// Describes a single task template field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldProperty := &fieldProperty{
//   	id: &fieldIdentifierProperty{
//   		name: jsii.String("name"),
//   	},
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	singleSelectOptions: []*string{
//   		jsii.String("singleSelectOptions"),
//   	},
//   }
//
type CfnTaskTemplate_FieldProperty struct {
	// The unique identifier for the field.
	Id interface{} `field:"required" json:"id" yaml:"id"`
	// Indicates the type of field.
	//
	// Following are the valid field types: `NAME` `DESCRIPTION` | `SCHEDULED_TIME` | `QUICK_CONNECT` | `URL` | `NUMBER` | `TEXT` | `TEXT_AREA` | `DATE_TIME` | `BOOLEAN` | `SINGLE_SELECT` | `EMAIL`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the field.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of options for a single select field.
	SingleSelectOptions *[]*string `field:"optional" json:"singleSelectOptions" yaml:"singleSelectOptions"`
}

