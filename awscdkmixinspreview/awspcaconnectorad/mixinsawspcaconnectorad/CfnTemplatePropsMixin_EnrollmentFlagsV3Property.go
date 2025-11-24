package mixinsawspcaconnectorad


// Template configurations for v3 template schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   enrollmentFlagsV3Property := &EnrollmentFlagsV3Property{
//   	EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   	IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   	NoSecurityExtension: jsii.Boolean(false),
//   	RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   	UserInteractionRequired: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html
//
type CfnTemplatePropsMixin_EnrollmentFlagsV3Property struct {
	// Allow renewal using the same key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html#cfn-pcaconnectorad-template-enrollmentflagsv3-enablekeyreuseonnttokenkeysetstoragefull
	//
	EnableKeyReuseOnNtTokenKeysetStorageFull interface{} `field:"optional" json:"enableKeyReuseOnNtTokenKeysetStorageFull" yaml:"enableKeyReuseOnNtTokenKeysetStorageFull"`
	// Include symmetric algorithms allowed by the subject.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html#cfn-pcaconnectorad-template-enrollmentflagsv3-includesymmetricalgorithms
	//
	IncludeSymmetricAlgorithms interface{} `field:"optional" json:"includeSymmetricAlgorithms" yaml:"includeSymmetricAlgorithms"`
	// This flag instructs the CA to not include the security extension szOID_NTDS_CA_SECURITY_EXT (OID:1.3.6.1.4.1.311.25.2), as specified in [MS-WCCE] sections 2.2.2.7.7.4 and 3.2.2.6.2.1.4.5.9, in the issued certificate. This addresses a Windows Kerberos elevation-of-privilege vulnerability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html#cfn-pcaconnectorad-template-enrollmentflagsv3-nosecurityextension
	//
	NoSecurityExtension interface{} `field:"optional" json:"noSecurityExtension" yaml:"noSecurityExtension"`
	// Delete expired or revoked certificates instead of archiving them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html#cfn-pcaconnectorad-template-enrollmentflagsv3-removeinvalidcertificatefrompersonalstore
	//
	RemoveInvalidCertificateFromPersonalStore interface{} `field:"optional" json:"removeInvalidCertificateFromPersonalStore" yaml:"removeInvalidCertificateFromPersonalStore"`
	// Require user interaction when the subject is enrolled and the private key associated with the certificate is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-enrollmentflagsv3.html#cfn-pcaconnectorad-template-enrollmentflagsv3-userinteractionrequired
	//
	UserInteractionRequired interface{} `field:"optional" json:"userInteractionRequired" yaml:"userInteractionRequired"`
}

