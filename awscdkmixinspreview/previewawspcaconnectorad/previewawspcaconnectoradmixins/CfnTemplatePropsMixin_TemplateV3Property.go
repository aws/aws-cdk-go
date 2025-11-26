package previewawspcaconnectoradmixins


// v3 template schema that uses Key Storage Providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   templateV3Property := &TemplateV3Property{
//   	CertificateValidity: &CertificateValidityProperty{
//   		RenewalPeriod: &ValidityPeriodProperty{
//   			Period: jsii.Number(123),
//   			PeriodType: jsii.String("periodType"),
//   		},
//   		ValidityPeriod: &ValidityPeriodProperty{
//   			Period: jsii.Number(123),
//   			PeriodType: jsii.String("periodType"),
//   		},
//   	},
//   	EnrollmentFlags: &EnrollmentFlagsV3Property{
//   		EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   		IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   		NoSecurityExtension: jsii.Boolean(false),
//   		RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   		UserInteractionRequired: jsii.Boolean(false),
//   	},
//   	Extensions: &ExtensionsV3Property{
//   		ApplicationPolicies: &ApplicationPoliciesProperty{
//   			Critical: jsii.Boolean(false),
//   			Policies: []interface{}{
//   				&ApplicationPolicyProperty{
//   					PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   					PolicyType: jsii.String("policyType"),
//   				},
//   			},
//   		},
//   		KeyUsage: &KeyUsageProperty{
//   			Critical: jsii.Boolean(false),
//   			UsageFlags: &KeyUsageFlagsProperty{
//   				DataEncipherment: jsii.Boolean(false),
//   				DigitalSignature: jsii.Boolean(false),
//   				KeyAgreement: jsii.Boolean(false),
//   				KeyEncipherment: jsii.Boolean(false),
//   				NonRepudiation: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	GeneralFlags: &GeneralFlagsV3Property{
//   		AutoEnrollment: jsii.Boolean(false),
//   		MachineType: jsii.Boolean(false),
//   	},
//   	HashAlgorithm: jsii.String("hashAlgorithm"),
//   	PrivateKeyAttributes: &PrivateKeyAttributesV3Property{
//   		Algorithm: jsii.String("algorithm"),
//   		CryptoProviders: []*string{
//   			jsii.String("cryptoProviders"),
//   		},
//   		KeySpec: jsii.String("keySpec"),
//   		KeyUsageProperty: &KeyUsagePropertyProperty{
//   			PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   				Decrypt: jsii.Boolean(false),
//   				KeyAgreement: jsii.Boolean(false),
//   				Sign: jsii.Boolean(false),
//   			},
//   			PropertyType: jsii.String("propertyType"),
//   		},
//   		MinimalKeyLength: jsii.Number(123),
//   	},
//   	PrivateKeyFlags: &PrivateKeyFlagsV3Property{
//   		ClientVersion: jsii.String("clientVersion"),
//   		ExportableKey: jsii.Boolean(false),
//   		RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   		StrongKeyProtectionRequired: jsii.Boolean(false),
//   	},
//   	SubjectNameFlags: &SubjectNameFlagsV3Property{
//   		RequireCommonName: jsii.Boolean(false),
//   		RequireDirectoryPath: jsii.Boolean(false),
//   		RequireDnsAsCn: jsii.Boolean(false),
//   		RequireEmail: jsii.Boolean(false),
//   		SanRequireDirectoryGuid: jsii.Boolean(false),
//   		SanRequireDns: jsii.Boolean(false),
//   		SanRequireDomainDns: jsii.Boolean(false),
//   		SanRequireEmail: jsii.Boolean(false),
//   		SanRequireSpn: jsii.Boolean(false),
//   		SanRequireUpn: jsii.Boolean(false),
//   	},
//   	SupersededTemplates: []*string{
//   		jsii.String("supersededTemplates"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html
//
type CfnTemplatePropsMixin_TemplateV3Property struct {
	// Certificate validity describes the validity and renewal periods of a certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-certificatevalidity
	//
	CertificateValidity interface{} `field:"optional" json:"certificateValidity" yaml:"certificateValidity"`
	// Enrollment flags describe the enrollment settings for certificates such as using the existing private key and deleting expired or revoked certificates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-enrollmentflags
	//
	EnrollmentFlags interface{} `field:"optional" json:"enrollmentFlags" yaml:"enrollmentFlags"`
	// Extensions describe the key usage extensions and application policies for a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-extensions
	//
	Extensions interface{} `field:"optional" json:"extensions" yaml:"extensions"`
	// General flags describe whether the template is used for computers or users and if the template can be used with autoenrollment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-generalflags
	//
	GeneralFlags interface{} `field:"optional" json:"generalFlags" yaml:"generalFlags"`
	// Specifies the hash algorithm used to hash the private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-hashalgorithm
	//
	HashAlgorithm *string `field:"optional" json:"hashAlgorithm" yaml:"hashAlgorithm"`
	// Private key attributes allow you to specify the algorithm, minimal key length, key spec, key usage, and cryptographic providers for the private key of a certificate for v3 templates.
	//
	// V3 templates allow you to use Key Storage Providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-privatekeyattributes
	//
	PrivateKeyAttributes interface{} `field:"optional" json:"privateKeyAttributes" yaml:"privateKeyAttributes"`
	// Private key flags for v3 templates specify the client compatibility, if the private key can be exported, if user input is required when using a private key, and if an alternate signature algorithm should be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-privatekeyflags
	//
	PrivateKeyFlags interface{} `field:"optional" json:"privateKeyFlags" yaml:"privateKeyFlags"`
	// Subject name flags describe the subject name and subject alternate name that is included in a certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-subjectnameflags
	//
	SubjectNameFlags interface{} `field:"optional" json:"subjectNameFlags" yaml:"subjectNameFlags"`
	// List of templates in Active Directory that are superseded by this template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-supersededtemplates
	//
	SupersededTemplates *[]*string `field:"optional" json:"supersededTemplates" yaml:"supersededTemplates"`
}

