package previewawscustomerprofilesmixins


// A map of the name and ObjectType field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-fieldmap.html
//
type CfnObjectTypePropsMixin_FieldMapProperty struct {
	// Name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-fieldmap.html#cfn-customerprofiles-objecttype-fieldmap-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Represents a field in a ProfileObjectType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-fieldmap.html#cfn-customerprofiles-objecttype-fieldmap-objecttypefield
	//
	ObjectTypeField interface{} `field:"optional" json:"objectTypeField" yaml:"objectTypeField"`
}

