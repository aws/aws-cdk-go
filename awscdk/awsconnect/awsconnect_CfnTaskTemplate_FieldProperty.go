package awsconnect


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
	// `CfnTaskTemplate.FieldProperty.Id`.
	Id interface{} `field:"required" json:"id" yaml:"id"`
	// `CfnTaskTemplate.FieldProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnTaskTemplate.FieldProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnTaskTemplate.FieldProperty.SingleSelectOptions`.
	SingleSelectOptions *[]*string `field:"optional" json:"singleSelectOptions" yaml:"singleSelectOptions"`
}

