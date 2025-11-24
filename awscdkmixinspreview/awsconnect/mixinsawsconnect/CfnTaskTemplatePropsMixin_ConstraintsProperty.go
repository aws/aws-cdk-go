package mixinsawsconnect


// Describes constraints that apply to the template fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   constraintsProperty := &ConstraintsProperty{
//   	InvisibleFields: []interface{}{
//   		&InvisibleFieldInfoProperty{
//   			Id: &FieldIdentifierProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	ReadOnlyFields: []interface{}{
//   		&ReadOnlyFieldInfoProperty{
//   			Id: &FieldIdentifierProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	RequiredFields: []interface{}{
//   		&RequiredFieldInfoProperty{
//   			Id: &FieldIdentifierProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-constraints.html
//
type CfnTaskTemplatePropsMixin_ConstraintsProperty struct {
	// Lists the fields that are invisible to agents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-constraints.html#cfn-connect-tasktemplate-constraints-invisiblefields
	//
	InvisibleFields interface{} `field:"optional" json:"invisibleFields" yaml:"invisibleFields"`
	// Lists the fields that are read-only to agents, and cannot be edited.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-constraints.html#cfn-connect-tasktemplate-constraints-readonlyfields
	//
	ReadOnlyFields interface{} `field:"optional" json:"readOnlyFields" yaml:"readOnlyFields"`
	// Lists the fields that are required to be filled by agents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-constraints.html#cfn-connect-tasktemplate-constraints-requiredfields
	//
	RequiredFields interface{} `field:"optional" json:"requiredFields" yaml:"requiredFields"`
}

