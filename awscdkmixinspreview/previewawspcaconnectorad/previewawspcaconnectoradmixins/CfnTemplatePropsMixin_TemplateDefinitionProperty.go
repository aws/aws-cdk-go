package previewawspcaconnectoradmixins


// Template configuration to define the information included in certificates.
//
// Define certificate validity and renewal periods, certificate request handling and enrollment options, key usage extensions, application policies, and cryptography settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   templateDefinitionProperty := &TemplateDefinitionProperty{
//   	TemplateV2: &TemplateV2Property{
//   		CertificateValidity: &CertificateValidityProperty{
//   			RenewalPeriod: &ValidityPeriodProperty{
//   				Period: jsii.Number(123),
//   				PeriodType: jsii.String("periodType"),
//   			},
//   			ValidityPeriod: &ValidityPeriodProperty{
//   				Period: jsii.Number(123),
//   				PeriodType: jsii.String("periodType"),
//   			},
//   		},
//   		EnrollmentFlags: &EnrollmentFlagsV2Property{
//   			EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   			IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   			NoSecurityExtension: jsii.Boolean(false),
//   			RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   			UserInteractionRequired: jsii.Boolean(false),
//   		},
//   		Extensions: &ExtensionsV2Property{
//   			ApplicationPolicies: &ApplicationPoliciesProperty{
//   				Critical: jsii.Boolean(false),
//   				Policies: []interface{}{
//   					&ApplicationPolicyProperty{
//   						PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   						PolicyType: jsii.String("policyType"),
//   					},
//   				},
//   			},
//   			KeyUsage: &KeyUsageProperty{
//   				Critical: jsii.Boolean(false),
//   				UsageFlags: &KeyUsageFlagsProperty{
//   					DataEncipherment: jsii.Boolean(false),
//   					DigitalSignature: jsii.Boolean(false),
//   					KeyAgreement: jsii.Boolean(false),
//   					KeyEncipherment: jsii.Boolean(false),
//   					NonRepudiation: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		GeneralFlags: &GeneralFlagsV2Property{
//   			AutoEnrollment: jsii.Boolean(false),
//   			MachineType: jsii.Boolean(false),
//   		},
//   		PrivateKeyAttributes: &PrivateKeyAttributesV2Property{
//   			CryptoProviders: []*string{
//   				jsii.String("cryptoProviders"),
//   			},
//   			KeySpec: jsii.String("keySpec"),
//   			MinimalKeyLength: jsii.Number(123),
//   		},
//   		PrivateKeyFlags: &PrivateKeyFlagsV2Property{
//   			ClientVersion: jsii.String("clientVersion"),
//   			ExportableKey: jsii.Boolean(false),
//   			StrongKeyProtectionRequired: jsii.Boolean(false),
//   		},
//   		SubjectNameFlags: &SubjectNameFlagsV2Property{
//   			RequireCommonName: jsii.Boolean(false),
//   			RequireDirectoryPath: jsii.Boolean(false),
//   			RequireDnsAsCn: jsii.Boolean(false),
//   			RequireEmail: jsii.Boolean(false),
//   			SanRequireDirectoryGuid: jsii.Boolean(false),
//   			SanRequireDns: jsii.Boolean(false),
//   			SanRequireDomainDns: jsii.Boolean(false),
//   			SanRequireEmail: jsii.Boolean(false),
//   			SanRequireSpn: jsii.Boolean(false),
//   			SanRequireUpn: jsii.Boolean(false),
//   		},
//   		SupersededTemplates: []*string{
//   			jsii.String("supersededTemplates"),
//   		},
//   	},
//   	TemplateV3: &TemplateV3Property{
//   		CertificateValidity: &CertificateValidityProperty{
//   			RenewalPeriod: &ValidityPeriodProperty{
//   				Period: jsii.Number(123),
//   				PeriodType: jsii.String("periodType"),
//   			},
//   			ValidityPeriod: &ValidityPeriodProperty{
//   				Period: jsii.Number(123),
//   				PeriodType: jsii.String("periodType"),
//   			},
//   		},
//   		EnrollmentFlags: &EnrollmentFlagsV3Property{
//   			EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   			IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   			NoSecurityExtension: jsii.Boolean(false),
//   			RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   			UserInteractionRequired: jsii.Boolean(false),
//   		},
//   		Extensions: &ExtensionsV3Property{
//   			ApplicationPolicies: &ApplicationPoliciesProperty{
//   				Critical: jsii.Boolean(false),
//   				Policies: []interface{}{
//   					&ApplicationPolicyProperty{
//   						PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   						PolicyType: jsii.String("policyType"),
//   					},
//   				},
//   			},
//   			KeyUsage: &KeyUsageProperty{
//   				Critical: jsii.Boolean(false),
//   				UsageFlags: &KeyUsageFlagsProperty{
//   					DataEncipherment: jsii.Boolean(false),
//   					DigitalSignature: jsii.Boolean(false),
//   					KeyAgreement: jsii.Boolean(false),
//   					KeyEncipherment: jsii.Boolean(false),
//   					NonRepudiation: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		GeneralFlags: &GeneralFlagsV3Property{
//   			AutoEnrollment: jsii.Boolean(false),
//   			MachineType: jsii.Boolean(false),
//   		},
//   		HashAlgorithm: jsii.String("hashAlgorithm"),
//   		PrivateKeyAttributes: &PrivateKeyAttributesV3Property{
//   			Algorithm: jsii.String("algorithm"),
//   			CryptoProviders: []*string{
//   				jsii.String("cryptoProviders"),
//   			},
//   			KeySpec: jsii.String("keySpec"),
//   			KeyUsageProperty: &KeyUsagePropertyProperty{
//   				PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   					Decrypt: jsii.Boolean(false),
//   					KeyAgreement: jsii.Boolean(false),
//   					Sign: jsii.Boolean(false),
//   				},
//   				PropertyType: jsii.String("propertyType"),
//   			},
//   			MinimalKeyLength: jsii.Number(123),
//   		},
//   		PrivateKeyFlags: &PrivateKeyFlagsV3Property{
//   			ClientVersion: jsii.String("clientVersion"),
//   			ExportableKey: jsii.Boolean(false),
//   			RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   			StrongKeyProtectionRequired: jsii.Boolean(false),
//   		},
//   		SubjectNameFlags: &SubjectNameFlagsV3Property{
//   			RequireCommonName: jsii.Boolean(false),
//   			RequireDirectoryPath: jsii.Boolean(false),
//   			RequireDnsAsCn: jsii.Boolean(false),
//   			RequireEmail: jsii.Boolean(false),
//   			SanRequireDirectoryGuid: jsii.Boolean(false),
//   			SanRequireDns: jsii.Boolean(false),
//   			SanRequireDomainDns: jsii.Boolean(false),
//   			SanRequireEmail: jsii.Boolean(false),
//   			SanRequireSpn: jsii.Boolean(false),
//   			SanRequireUpn: jsii.Boolean(false),
//   		},
//   		SupersededTemplates: []*string{
//   			jsii.String("supersededTemplates"),
//   		},
//   	},
//   	TemplateV4: &TemplateV4Property{
//   		CertificateValidity: &CertificateValidityProperty{
//   			RenewalPeriod: &ValidityPeriodProperty{
//   				Period: jsii.Number(123),
//   				PeriodType: jsii.String("periodType"),
//   			},
//   			ValidityPeriod: &ValidityPeriodProperty{
//   				Period: jsii.Number(123),
//   				PeriodType: jsii.String("periodType"),
//   			},
//   		},
//   		EnrollmentFlags: &EnrollmentFlagsV4Property{
//   			EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   			IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   			NoSecurityExtension: jsii.Boolean(false),
//   			RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   			UserInteractionRequired: jsii.Boolean(false),
//   		},
//   		Extensions: &ExtensionsV4Property{
//   			ApplicationPolicies: &ApplicationPoliciesProperty{
//   				Critical: jsii.Boolean(false),
//   				Policies: []interface{}{
//   					&ApplicationPolicyProperty{
//   						PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   						PolicyType: jsii.String("policyType"),
//   					},
//   				},
//   			},
//   			KeyUsage: &KeyUsageProperty{
//   				Critical: jsii.Boolean(false),
//   				UsageFlags: &KeyUsageFlagsProperty{
//   					DataEncipherment: jsii.Boolean(false),
//   					DigitalSignature: jsii.Boolean(false),
//   					KeyAgreement: jsii.Boolean(false),
//   					KeyEncipherment: jsii.Boolean(false),
//   					NonRepudiation: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		GeneralFlags: &GeneralFlagsV4Property{
//   			AutoEnrollment: jsii.Boolean(false),
//   			MachineType: jsii.Boolean(false),
//   		},
//   		HashAlgorithm: jsii.String("hashAlgorithm"),
//   		PrivateKeyAttributes: &PrivateKeyAttributesV4Property{
//   			Algorithm: jsii.String("algorithm"),
//   			CryptoProviders: []*string{
//   				jsii.String("cryptoProviders"),
//   			},
//   			KeySpec: jsii.String("keySpec"),
//   			KeyUsageProperty: &KeyUsagePropertyProperty{
//   				PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   					Decrypt: jsii.Boolean(false),
//   					KeyAgreement: jsii.Boolean(false),
//   					Sign: jsii.Boolean(false),
//   				},
//   				PropertyType: jsii.String("propertyType"),
//   			},
//   			MinimalKeyLength: jsii.Number(123),
//   		},
//   		PrivateKeyFlags: &PrivateKeyFlagsV4Property{
//   			ClientVersion: jsii.String("clientVersion"),
//   			ExportableKey: jsii.Boolean(false),
//   			RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   			RequireSameKeyRenewal: jsii.Boolean(false),
//   			StrongKeyProtectionRequired: jsii.Boolean(false),
//   			UseLegacyProvider: jsii.Boolean(false),
//   		},
//   		SubjectNameFlags: &SubjectNameFlagsV4Property{
//   			RequireCommonName: jsii.Boolean(false),
//   			RequireDirectoryPath: jsii.Boolean(false),
//   			RequireDnsAsCn: jsii.Boolean(false),
//   			RequireEmail: jsii.Boolean(false),
//   			SanRequireDirectoryGuid: jsii.Boolean(false),
//   			SanRequireDns: jsii.Boolean(false),
//   			SanRequireDomainDns: jsii.Boolean(false),
//   			SanRequireEmail: jsii.Boolean(false),
//   			SanRequireSpn: jsii.Boolean(false),
//   			SanRequireUpn: jsii.Boolean(false),
//   		},
//   		SupersededTemplates: []*string{
//   			jsii.String("supersededTemplates"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatedefinition.html
//
type CfnTemplatePropsMixin_TemplateDefinitionProperty struct {
	// Template configuration to define the information included in certificates.
	//
	// Define certificate validity and renewal periods, certificate request handling and enrollment options, key usage extensions, application policies, and cryptography settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatedefinition.html#cfn-pcaconnectorad-template-templatedefinition-templatev2
	//
	TemplateV2 interface{} `field:"optional" json:"templateV2" yaml:"templateV2"`
	// Template configuration to define the information included in certificates.
	//
	// Define certificate validity and renewal periods, certificate request handling and enrollment options, key usage extensions, application policies, and cryptography settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatedefinition.html#cfn-pcaconnectorad-template-templatedefinition-templatev3
	//
	TemplateV3 interface{} `field:"optional" json:"templateV3" yaml:"templateV3"`
	// Template configuration to define the information included in certificates.
	//
	// Define certificate validity and renewal periods, certificate request handling and enrollment options, key usage extensions, application policies, and cryptography settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatedefinition.html#cfn-pcaconnectorad-template-templatedefinition-templatev4
	//
	TemplateV4 interface{} `field:"optional" json:"templateV4" yaml:"templateV4"`
}

