package awsconnect


// Describes a single task template field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldProperty := &FieldProperty{
//   	Id: &FieldIdentifierProperty{
//   		Name: jsii.String("name"),
//   	},
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	SingleSelectOptions: []*string{
//   		jsii.String("singleSelectOptions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-field.html
//
type CfnTaskTemplate_FieldProperty struct {
	// The unique identifier for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-field.html#cfn-connect-tasktemplate-field-id
	//
	Id interface{} `field:"required" json:"id" yaml:"id"`
	// Indicates the type of field.
	//
	// Following are the valid field types: `NAME` `DESCRIPTION` | `SCHEDULED_TIME` | `QUICK_CONNECT` | `URL` | `NUMBER` | `TEXT` | `TEXT_AREA` | `DATE_TIME` | `BOOLEAN` | `SINGLE_SELECT` | `EMAIL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-field.html#cfn-connect-tasktemplate-field-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-field.html#cfn-connect-tasktemplate-field-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of options for a single select field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-field.html#cfn-connect-tasktemplate-field-singleselectoptions
	//
	SingleSelectOptions *[]*string `field:"optional" json:"singleSelectOptions" yaml:"singleSelectOptions"`
}

