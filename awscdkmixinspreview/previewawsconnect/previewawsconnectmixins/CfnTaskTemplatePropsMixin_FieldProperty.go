package previewawsconnectmixins


// Describes a single task template field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fieldProperty := &FieldProperty{
//   	Description: jsii.String("description"),
//   	Id: &FieldIdentifierProperty{
//   		Name: jsii.String("name"),
//   	},
//   	SingleSelectOptions: []*string{
//   		jsii.String("singleSelectOptions"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-field.html
//
type CfnTaskTemplatePropsMixin_FieldProperty struct {
	// The description of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-field.html#cfn-connect-tasktemplate-field-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-field.html#cfn-connect-tasktemplate-field-id
	//
	Id interface{} `field:"optional" json:"id" yaml:"id"`
	// A list of options for a single select field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-field.html#cfn-connect-tasktemplate-field-singleselectoptions
	//
	SingleSelectOptions *[]*string `field:"optional" json:"singleSelectOptions" yaml:"singleSelectOptions"`
	// Indicates the type of field.
	//
	// Following are the valid field types: `NAME` `DESCRIPTION` | `SCHEDULED_TIME` | `QUICK_CONNECT` | `URL` | `NUMBER` | `TEXT` | `TEXT_AREA` | `DATE_TIME` | `BOOLEAN` | `SINGLE_SELECT` | `EMAIL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-field.html#cfn-connect-tasktemplate-field-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

