package awsconnect


// Describes constraints that apply to the template fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnTaskTemplate_ConstraintsProperty struct {
	// Lists the fields that are invisible to agents.
	InvisibleFields interface{} `field:"optional" json:"invisibleFields" yaml:"invisibleFields"`
	// Lists the fields that are read-only to agents, and cannot be edited.
	ReadOnlyFields interface{} `field:"optional" json:"readOnlyFields" yaml:"readOnlyFields"`
	// Lists the fields that are required to be filled by agents.
	RequiredFields interface{} `field:"optional" json:"requiredFields" yaml:"requiredFields"`
}

