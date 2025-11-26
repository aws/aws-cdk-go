package previewawsrolesanywheremixins


// A mapping applied to the authenticating end-entity certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attributeMappingProperty := &AttributeMappingProperty{
//   	CertificateField: jsii.String("certificateField"),
//   	MappingRules: []interface{}{
//   		&MappingRuleProperty{
//   			Specifier: jsii.String("specifier"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-profile-attributemapping.html
//
type CfnProfilePropsMixin_AttributeMappingProperty struct {
	// Fields (x509Subject, x509Issuer and x509SAN) within X.509 certificates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-profile-attributemapping.html#cfn-rolesanywhere-profile-attributemapping-certificatefield
	//
	CertificateField *string `field:"optional" json:"certificateField" yaml:"certificateField"`
	// A list of mapping entries for every supported specifier or sub-field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-profile-attributemapping.html#cfn-rolesanywhere-profile-attributemapping-mappingrules
	//
	MappingRules interface{} `field:"optional" json:"mappingRules" yaml:"mappingRules"`
}

