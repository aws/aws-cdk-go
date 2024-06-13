package awsrolesanywhere


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnProfile_AttributeMappingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-profile-attributemapping.html#cfn-rolesanywhere-profile-attributemapping-certificatefield
	//
	CertificateField *string `field:"required" json:"certificateField" yaml:"certificateField"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-profile-attributemapping.html#cfn-rolesanywhere-profile-attributemapping-mappingrules
	//
	MappingRules interface{} `field:"required" json:"mappingRules" yaml:"mappingRules"`
}

