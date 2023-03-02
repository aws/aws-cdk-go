package awscustomerprofiles


// Represents a field in a ProfileObjectType.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectTypeFieldProperty := &objectTypeFieldProperty{
//   	contentType: jsii.String("contentType"),
//   	source: jsii.String("source"),
//   	target: jsii.String("target"),
//   }
//
type CfnObjectType_ObjectTypeFieldProperty struct {
	// The content type of the field.
	//
	// Used for determining equality when searching.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A field of a ProfileObject.
	//
	// For example: _source.FirstName, where “_source” is a ProfileObjectType of a Zendesk user and “FirstName” is a field in that ObjectType.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The location of the data in the standard ProfileObject model.
	//
	// For example: _profile.Address.PostalCode.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

