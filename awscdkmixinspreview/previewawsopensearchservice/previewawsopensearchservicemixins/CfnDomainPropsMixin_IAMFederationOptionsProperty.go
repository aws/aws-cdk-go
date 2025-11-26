package previewawsopensearchservicemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iAMFederationOptionsProperty := map[string]interface{}{
//   	"enabled": jsii.Boolean(false),
//   	"rolesKey": jsii.String("rolesKey"),
//   	"subjectKey": jsii.String("subjectKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-iamfederationoptions.html
//
type CfnDomainPropsMixin_IAMFederationOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-iamfederationoptions.html#cfn-opensearchservice-domain-iamfederationoptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-iamfederationoptions.html#cfn-opensearchservice-domain-iamfederationoptions-roleskey
	//
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-iamfederationoptions.html#cfn-opensearchservice-domain-iamfederationoptions-subjectkey
	//
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

