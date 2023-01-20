package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultFieldValueProperty := &defaultFieldValueProperty{
//   	defaultValue: jsii.String("defaultValue"),
//   	id: &fieldIdentifierProperty{
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnTaskTemplate_DefaultFieldValueProperty struct {
	// `CfnTaskTemplate.DefaultFieldValueProperty.DefaultValue`.
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
	// `CfnTaskTemplate.DefaultFieldValueProperty.Id`.
	Id interface{} `field:"required" json:"id" yaml:"id"`
}

