package awsrolesanywhere


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mappingRuleProperty := &MappingRuleProperty{
//   	Specifier: jsii.String("specifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-profile-mappingrule.html
//
type CfnProfile_MappingRuleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-profile-mappingrule.html#cfn-rolesanywhere-profile-mappingrule-specifier
	//
	Specifier *string `field:"required" json:"specifier" yaml:"specifier"`
}

