package awsconnect


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
	// `CfnTaskTemplate.ConstraintsProperty.InvisibleFields`.
	InvisibleFields interface{} `field:"optional" json:"invisibleFields" yaml:"invisibleFields"`
	// `CfnTaskTemplate.ConstraintsProperty.ReadOnlyFields`.
	ReadOnlyFields interface{} `field:"optional" json:"readOnlyFields" yaml:"readOnlyFields"`
	// `CfnTaskTemplate.ConstraintsProperty.RequiredFields`.
	RequiredFields interface{} `field:"optional" json:"requiredFields" yaml:"requiredFields"`
}

