package previewawsrolesanywheremixins


// A single mapping entry for each supported specifier or sub-field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mappingRuleProperty := &MappingRuleProperty{
//   	Specifier: jsii.String("specifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-profile-mappingrule.html
//
type CfnProfilePropsMixin_MappingRuleProperty struct {
	// Specifier within a certificate field, such as CN, OU, or UID from the Subject field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-profile-mappingrule.html#cfn-rolesanywhere-profile-mappingrule-specifier
	//
	Specifier *string `field:"optional" json:"specifier" yaml:"specifier"`
}

