package awscustomerprofiles


// Represents a field in a ProfileObjectType.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectTypeFieldProperty := &ObjectTypeFieldProperty{
//   	ContentType: jsii.String("contentType"),
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-objecttypefield.html
//
type CfnObjectType_ObjectTypeFieldProperty struct {
	// The content type of the field.
	//
	// Used for determining equality when searching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-objecttypefield.html#cfn-customerprofiles-objecttype-objecttypefield-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A field of a ProfileObject.
	//
	// For example: _source.FirstName, where “_source” is a ProfileObjectType of a Zendesk user and “FirstName” is a field in that ObjectType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-objecttypefield.html#cfn-customerprofiles-objecttype-objecttypefield-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The location of the data in the standard ProfileObject model.
	//
	// For example: _profile.Address.PostalCode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-objecttypefield.html#cfn-customerprofiles-objecttype-objecttypefield-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
}

