package previewawscasesmixins


// Object for a group of fields and associated properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fieldGroupProperty := &FieldGroupProperty{
//   	Fields: []interface{}{
//   		&FieldItemProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-fieldgroup.html
//
type CfnLayoutPropsMixin_FieldGroupProperty struct {
	// Represents an ordered list containing field related information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-fieldgroup.html#cfn-cases-layout-fieldgroup-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// Name of the field group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-fieldgroup.html#cfn-cases-layout-fieldgroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

