package awscustomerprofiles


// A map of the name and ObjectType field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldMapProperty := &fieldMapProperty{
//   	name: jsii.String("name"),
//   	objectTypeField: &objectTypeFieldProperty{
//   		contentType: jsii.String("contentType"),
//   		source: jsii.String("source"),
//   		target: jsii.String("target"),
//   	},
//   }
//
type CfnObjectType_FieldMapProperty struct {
	// Name of the field.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Represents a field in a ProfileObjectType.
	ObjectTypeField interface{} `field:"optional" json:"objectTypeField" yaml:"objectTypeField"`
}

