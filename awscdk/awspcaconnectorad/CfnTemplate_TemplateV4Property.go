package awspcaconnectorad


// v4 template schema that can use either Legacy Cryptographic Providers or Key Storage Providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   		KeyUsage: &KeyUsageProperty{
//   			UsageFlags: &KeyUsageFlagsProperty{
//   				DataEncipherment: jsii.Boolean(false),
//   				DigitalSignature: jsii.Boolean(false),
//   				KeyAgreement: jsii.Boolean(false),
//   				KeyEncipherment: jsii.Boolean(false),
//   				NonRepudiation: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			Critical: jsii.Boolean(false),
//   		},
//
//   		// the properties below are optional
//   		ApplicationPolicies: &ApplicationPoliciesProperty{
//   			Policies: []interface{}{
//   				&ApplicationPolicyProperty{
//   					PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   					PolicyType: jsii.String("policyType"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Critical: jsii.Boolean(false),
//   		},
//   	},
//   	GeneralFlags: &GeneralFlagsV4Property{
//   		AutoEnrollment: jsii.Boolean(false),
//   		MachineType: jsii.Boolean(false),
//   	},
//   	PrivateKeyAttributes: &PrivateKeyAttributesV4Property{
//   		KeySpec: jsii.String("keySpec"),
//   		MinimalKeyLength: jsii.Number(123),
//
//   		// the properties below are optional
//   		Algorithm: jsii.String("algorithm"),
//   		CryptoProviders: []*string{
//   			jsii.String("cryptoProviders"),
//   		},
//   		KeyUsageProperty: &KeyUsagePropertyProperty{
//   			PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   				Decrypt: jsii.Boolean(false),
//   				KeyAgreement: jsii.Boolean(false),
//   				Sign: jsii.Boolean(false),
//   			},
//   			PropertyType: jsii.String("propertyType"),
//   		},
//   	},
//   	PrivateKeyFlags: &PrivateKeyFlagsV4Property{
//   		ClientVersion: jsii.String("clientVersion"),
//
//   		// the properties below are optional
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
//
//   	// the properties below are optional
//   	HashAlgorithm: jsii.String("hashAlgorithm"),
//   	SupersededTemplates: []*string{
//   		jsii.String("supersededTemplates"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html
//
type CfnTemplate_TemplateV4Property struct {
	// Certificate validity describes the validity and renewal periods of a certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-certificatevalidity
	//
	CertificateValidity interface{} `field:"required" json:"certificateValidity" yaml:"certificateValidity"`
	// Enrollment flags describe the enrollment settings for certificates using the existing private key and deleting expired or revoked certificates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-enrollmentflags
	//
	EnrollmentFlags interface{} `field:"required" json:"enrollmentFlags" yaml:"enrollmentFlags"`
	// Extensions describe the key usage extensions and application policies for a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-extensions
	//
	Extensions interface{} `field:"required" json:"extensions" yaml:"extensions"`
	// General flags describe whether the template is used for computers or users and if the template can be used with autoenrollment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-generalflags
	//
	GeneralFlags interface{} `field:"required" json:"generalFlags" yaml:"generalFlags"`
	// Private key attributes allow you to specify the minimal key length, key spec, key usage, and cryptographic providers for the private key of a certificate for v4 templates.
	//
	// V4 templates allow you to use either Key Storage Providers or Legacy Cryptographic Service Providers. You specify the cryptography provider category in private key flags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-privatekeyattributes
	//
	PrivateKeyAttributes interface{} `field:"required" json:"privateKeyAttributes" yaml:"privateKeyAttributes"`
	// Private key flags for v4 templates specify the client compatibility, if the private key can be exported, if user input is required when using a private key, if an alternate signature algorithm should be used, and if certificates are renewed using the same private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-privatekeyflags
	//
	PrivateKeyFlags interface{} `field:"required" json:"privateKeyFlags" yaml:"privateKeyFlags"`
	// Subject name flags describe the subject name and subject alternate name that is included in a certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-subjectnameflags
	//
	SubjectNameFlags interface{} `field:"required" json:"subjectNameFlags" yaml:"subjectNameFlags"`
	// Specifies the hash algorithm used to hash the private key.
	//
	// Hash algorithm can only be specified when using Key Storage Providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-hashalgorithm
	//
	HashAlgorithm *string `field:"optional" json:"hashAlgorithm" yaml:"hashAlgorithm"`
	// List of templates in Active Directory that are superseded by this template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev4.html#cfn-pcaconnectorad-template-templatev4-supersededtemplates
	//
	SupersededTemplates *[]*string `field:"optional" json:"supersededTemplates" yaml:"supersededTemplates"`
}

