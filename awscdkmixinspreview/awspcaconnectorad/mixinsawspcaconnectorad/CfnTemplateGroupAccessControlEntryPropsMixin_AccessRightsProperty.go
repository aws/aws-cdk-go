package mixinsawspcaconnectorad


// Allow or deny permissions for an Active Directory group to enroll or autoenroll certificates for a template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessRightsProperty := &AccessRightsProperty{
//   	AutoEnroll: jsii.String("autoEnroll"),
//   	Enroll: jsii.String("enroll"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-templategroupaccesscontrolentry-accessrights.html
//
type CfnTemplateGroupAccessControlEntryPropsMixin_AccessRightsProperty struct {
	// Allow or deny an Active Directory group from autoenrolling certificates issued against a template.
	//
	// The Active Directory group must be allowed to enroll to allow autoenrollment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-templategroupaccesscontrolentry-accessrights.html#cfn-pcaconnectorad-templategroupaccesscontrolentry-accessrights-autoenroll
	//
	AutoEnroll *string `field:"optional" json:"autoEnroll" yaml:"autoEnroll"`
	// Allow or deny an Active Directory group from enrolling certificates issued against a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-templategroupaccesscontrolentry-accessrights.html#cfn-pcaconnectorad-templategroupaccesscontrolentry-accessrights-enroll
	//
	Enroll *string `field:"optional" json:"enroll" yaml:"enroll"`
}

