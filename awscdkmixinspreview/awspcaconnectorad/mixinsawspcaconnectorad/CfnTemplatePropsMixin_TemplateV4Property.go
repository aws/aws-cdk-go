package mixinsawspcaconnectorad


// v4 template schema that can use either Legacy Cryptographic Providers or Key Storage Providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   templateV4Property := &TemplateV4Property{
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
//   	EnrollmentFlags: &EnrollmentFlagsV4Property{
//   		EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   		IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   		NoSecurityExtension: jsii.Boolean(false),
//   		RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   		UserInteractionRequired: jsii.Boolean(false),
//   	},
//   	Extensions: &ExtensionsV4Property{
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
//   	GeneralFlags: &GeneralFlagsV4Property{
//   		AutoEnrollment: jsii.Boolean(false),
//   		MachineType: jsii.Boolean(false),
//   	},
//   	HashAlgorithm: jsii.String("hashAlgorithm"),
//   	PrivateKeyAttributes: &PrivateKeyAttributesV4Property{
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
//   	PrivateKeyFlags: &PrivateKeyFlagsV4Property{
//   		ClientVersion: jsii.String("clientVersion"),
//   		ExportableKey: jsii.Boolean(false),
//   		RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   		RequireSameKeyRenewal: jsii.Boolean(false),
//   		StrongKeyProtectionRequired: jsii.Boolean(false),
//   		UseLegacyProvider: jsii.Boolean(false),
//   	},
//   	SubjectNameFlags: &SubjectNameFlagsV4Property{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html
//
type CfnTemplatePropsMixin_TemplateV4Property struct {
	// Certificate validity describes the validity and renewal periods of a certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-certificatevalidity
	//
	CertificateValidity interface{} `field:"optional" json:"certificateValidity" yaml:"certificateValidity"`
	// Enrollment flags describe the enrollment settings for certificates using the existing private key and deleting expired or revoked certificates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-enrollmentflags
	//
	EnrollmentFlags interface{} `field:"optional" json:"enrollmentFlags" yaml:"enrollmentFlags"`
	// Extensions describe the key usage extensions and application policies for a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-extensions
	//
	Extensions interface{} `field:"optional" json:"extensions" yaml:"extensions"`
	// General flags describe whether the template is used for computers or users and if the template can be used with autoenrollment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-generalflags
	//
	GeneralFlags interface{} `field:"optional" json:"generalFlags" yaml:"generalFlags"`
	// Specifies the hash algorithm used to hash the private key.
	//
	// Hash algorithm can only be specified when using Key Storage Providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-hashalgorithm
	//
	HashAlgorithm *string `field:"optional" json:"hashAlgorithm" yaml:"hashAlgorithm"`
	// Private key attributes allow you to specify the minimal key length, key spec, key usage, and cryptographic providers for the private key of a certificate for v4 templates.
	//
	// V4 templates allow you to use either Key Storage Providers or Legacy Cryptographic Service Providers. You specify the cryptography provider category in private key flags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-privatekeyattributes
	//
	PrivateKeyAttributes interface{} `field:"optional" json:"privateKeyAttributes" yaml:"privateKeyAttributes"`
	// Private key flags for v4 templates specify the client compatibility, if the private key can be exported, if user input is required when using a private key, if an alternate signature algorithm should be used, and if certificates are renewed using the same private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-privatekeyflags
	//
	PrivateKeyFlags interface{} `field:"optional" json:"privateKeyFlags" yaml:"privateKeyFlags"`
	// Subject name flags describe the subject name and subject alternate name that is included in a certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-subjectnameflags
	//
	SubjectNameFlags interface{} `field:"optional" json:"subjectNameFlags" yaml:"subjectNameFlags"`
	// List of templates in Active Directory that are superseded by this template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-supersededtemplates
	//
	SupersededTemplates *[]*string `field:"optional" json:"supersededTemplates" yaml:"supersededTemplates"`
}

