package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   constraintsProperty := &constraintsProperty{
//   	invisibleFields: []interface{}{
//   		&invisibleFieldInfoProperty{
//   			id: &fieldIdentifierProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	readOnlyFields: []interface{}{
//   		&readOnlyFieldInfoProperty{
//   			id: &fieldIdentifierProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	requiredFields: []interface{}{
//   		&requiredFieldInfoProperty{
//   			id: &fieldIdentifierProperty{
//   				name: jsii.String("name"),
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

