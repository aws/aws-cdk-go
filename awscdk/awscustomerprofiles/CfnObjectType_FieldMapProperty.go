package awscustomerprofiles


// A map of the name and ObjectType field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldMapProperty := &FieldMapProperty{
//   	Name: jsii.String("name"),
//   	ObjectTypeField: &ObjectTypeFieldProperty{
//   		ContentType: jsii.String("contentType"),
//   		Source: jsii.String("source"),
//   		Target: jsii.String("target"),
//   	},
//   }
//
type CfnObjectType_FieldMapProperty struct {
	// Name of the field.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Represents a field in a ProfileObjectType.
	ObjectTypeField interface{} `field:"optional" json:"objectTypeField" yaml:"objectTypeField"`
}

