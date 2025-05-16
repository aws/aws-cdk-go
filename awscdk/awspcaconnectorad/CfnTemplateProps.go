package awspcaconnectorad


// Properties for defining a `CfnTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTemplateProps := &CfnTemplateProps{
//   	ConnectorArn: jsii.String("connectorArn"),
//   	Definition: &TemplateDefinitionProperty{
//   		TemplateV2: &TemplateV2Property{
//   			CertificateValidity: &CertificateValidityProperty{
//   				RenewalPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   				ValidityPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   			},
//   			EnrollmentFlags: &EnrollmentFlagsV2Property{
//   				EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   				IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   				NoSecurityExtension: jsii.Boolean(false),
//   				RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   				UserInteractionRequired: jsii.Boolean(false),
//   			},
//   			Extensions: &ExtensionsV2Property{
//   				KeyUsage: &KeyUsageProperty{
//   					UsageFlags: &KeyUsageFlagsProperty{
//   						DataEncipherment: jsii.Boolean(false),
//   						DigitalSignature: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						KeyEncipherment: jsii.Boolean(false),
//   						NonRepudiation: jsii.Boolean(false),
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//
//   				// the properties below are optional
//   				ApplicationPolicies: &ApplicationPoliciesProperty{
//   					Policies: []interface{}{
//   						&ApplicationPolicyProperty{
//   							PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   							PolicyType: jsii.String("policyType"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//   			},
//   			GeneralFlags: &GeneralFlagsV2Property{
//   				AutoEnrollment: jsii.Boolean(false),
//   				MachineType: jsii.Boolean(false),
//   			},
//   			PrivateKeyAttributes: &PrivateKeyAttributesV2Property{
//   				KeySpec: jsii.String("keySpec"),
//   				MinimalKeyLength: jsii.Number(123),
//
//   				// the properties below are optional
//   				CryptoProviders: []*string{
//   					jsii.String("cryptoProviders"),
//   				},
//   			},
//   			PrivateKeyFlags: &PrivateKeyFlagsV2Property{
//   				ClientVersion: jsii.String("clientVersion"),
//
//   				// the properties below are optional
//   				ExportableKey: jsii.Boolean(false),
//   				StrongKeyProtectionRequired: jsii.Boolean(false),
//   			},
//   			SubjectNameFlags: &SubjectNameFlagsV2Property{
//   				RequireCommonName: jsii.Boolean(false),
//   				RequireDirectoryPath: jsii.Boolean(false),
//   				RequireDnsAsCn: jsii.Boolean(false),
//   				RequireEmail: jsii.Boolean(false),
//   				SanRequireDirectoryGuid: jsii.Boolean(false),
//   				SanRequireDns: jsii.Boolean(false),
//   				SanRequireDomainDns: jsii.Boolean(false),
//   				SanRequireEmail: jsii.Boolean(false),
//   				SanRequireSpn: jsii.Boolean(false),
//   				SanRequireUpn: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			SupersededTemplates: []*string{
//   				jsii.String("supersededTemplates"),
//   			},
//   		},
//   		TemplateV3: &TemplateV3Property{
//   			CertificateValidity: &CertificateValidityProperty{
//   				RenewalPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   				ValidityPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   			},
//   			EnrollmentFlags: &EnrollmentFlagsV3Property{
//   				EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   				IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   				NoSecurityExtension: jsii.Boolean(false),
//   				RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   				UserInteractionRequired: jsii.Boolean(false),
//   			},
//   			Extensions: &ExtensionsV3Property{
//   				KeyUsage: &KeyUsageProperty{
//   					UsageFlags: &KeyUsageFlagsProperty{
//   						DataEncipherment: jsii.Boolean(false),
//   						DigitalSignature: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						KeyEncipherment: jsii.Boolean(false),
//   						NonRepudiation: jsii.Boolean(false),
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//
//   				// the properties below are optional
//   				ApplicationPolicies: &ApplicationPoliciesProperty{
//   					Policies: []interface{}{
//   						&ApplicationPolicyProperty{
//   							PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   							PolicyType: jsii.String("policyType"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//   			},
//   			GeneralFlags: &GeneralFlagsV3Property{
//   				AutoEnrollment: jsii.Boolean(false),
//   				MachineType: jsii.Boolean(false),
//   			},
//   			HashAlgorithm: jsii.String("hashAlgorithm"),
//   			PrivateKeyAttributes: &PrivateKeyAttributesV3Property{
//   				Algorithm: jsii.String("algorithm"),
//   				KeySpec: jsii.String("keySpec"),
//   				KeyUsageProperty: &KeyUsagePropertyProperty{
//   					PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   						Decrypt: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						Sign: jsii.Boolean(false),
//   					},
//   					PropertyType: jsii.String("propertyType"),
//   				},
//   				MinimalKeyLength: jsii.Number(123),
//
//   				// the properties below are optional
//   				CryptoProviders: []*string{
//   					jsii.String("cryptoProviders"),
//   				},
//   			},
//   			PrivateKeyFlags: &PrivateKeyFlagsV3Property{
//   				ClientVersion: jsii.String("clientVersion"),
//
//   				// the properties below are optional
//   				ExportableKey: jsii.Boolean(false),
//   				RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   				StrongKeyProtectionRequired: jsii.Boolean(false),
//   			},
//   			SubjectNameFlags: &SubjectNameFlagsV3Property{
//   				RequireCommonName: jsii.Boolean(false),
//   				RequireDirectoryPath: jsii.Boolean(false),
//   				RequireDnsAsCn: jsii.Boolean(false),
//   				RequireEmail: jsii.Boolean(false),
//   				SanRequireDirectoryGuid: jsii.Boolean(false),
//   				SanRequireDns: jsii.Boolean(false),
//   				SanRequireDomainDns: jsii.Boolean(false),
//   				SanRequireEmail: jsii.Boolean(false),
//   				SanRequireSpn: jsii.Boolean(false),
//   				SanRequireUpn: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			SupersededTemplates: []*string{
//   				jsii.String("supersededTemplates"),
//   			},
//   		},
//   		TemplateV4: &TemplateV4Property{
//   			CertificateValidity: &CertificateValidityProperty{
//   				RenewalPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   				ValidityPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   			},
//   			EnrollmentFlags: &EnrollmentFlagsV4Property{
//   				EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   				IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   				NoSecurityExtension: jsii.Boolean(false),
//   				RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   				UserInteractionRequired: jsii.Boolean(false),
//   			},
//   			Extensions: &ExtensionsV4Property{
//   				KeyUsage: &KeyUsageProperty{
//   					UsageFlags: &KeyUsageFlagsProperty{
//   						DataEncipherment: jsii.Boolean(false),
//   						DigitalSignature: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						KeyEncipherment: jsii.Boolean(false),
//   						NonRepudiation: jsii.Boolean(false),
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//
//   				// the properties below are optional
//   				ApplicationPolicies: &ApplicationPoliciesProperty{
//   					Policies: []interface{}{
//   						&ApplicationPolicyProperty{
//   							PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   							PolicyType: jsii.String("policyType"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//   			},
//   			GeneralFlags: &GeneralFlagsV4Property{
//   				AutoEnrollment: jsii.Boolean(false),
//   				MachineType: jsii.Boolean(false),
//   			},
//   			PrivateKeyAttributes: &PrivateKeyAttributesV4Property{
//   				KeySpec: jsii.String("keySpec"),
//   				MinimalKeyLength: jsii.Number(123),
//
//   				// the properties below are optional
//   				Algorithm: jsii.String("algorithm"),
//   				CryptoProviders: []*string{
//   					jsii.String("cryptoProviders"),
//   				},
//   				KeyUsageProperty: &KeyUsagePropertyProperty{
//   					PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   						Decrypt: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						Sign: jsii.Boolean(false),
//   					},
//   					PropertyType: jsii.String("propertyType"),
//   				},
//   			},
//   			PrivateKeyFlags: &PrivateKeyFlagsV4Property{
//   				ClientVersion: jsii.String("clientVersion"),
//
//   				// the properties below are optional
//   				ExportableKey: jsii.Boolean(false),
//   				RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   				RequireSameKeyRenewal: jsii.Boolean(false),
//   				StrongKeyProtectionRequired: jsii.Boolean(false),
//   				UseLegacyProvider: jsii.Boolean(false),
//   			},
//   			SubjectNameFlags: &SubjectNameFlagsV4Property{
//   				RequireCommonName: jsii.Boolean(false),
//   				RequireDirectoryPath: jsii.Boolean(false),
//   				RequireDnsAsCn: jsii.Boolean(false),
//   				RequireEmail: jsii.Boolean(false),
//   				SanRequireDirectoryGuid: jsii.Boolean(false),
//   				SanRequireDns: jsii.Boolean(false),
//   				SanRequireDomainDns: jsii.Boolean(false),
//   				SanRequireEmail: jsii.Boolean(false),
//   				SanRequireSpn: jsii.Boolean(false),
//   				SanRequireUpn: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			HashAlgorithm: jsii.String("hashAlgorithm"),
//   			SupersededTemplates: []*string{
//   				jsii.String("supersededTemplates"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ReenrollAllCertificateHolders: jsii.Boolean(false),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-template.html
//
type CfnTemplateProps struct {
	// The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-template.html#cfn-pcaconnectorad-template-connectorarn
	//
	ConnectorArn *string `field:"required" json:"connectorArn" yaml:"connectorArn"`
	// Template configuration to define the information included in certificates.
	//
	// Define certificate validity and renewal periods, certificate request handling and enrollment options, key usage extensions, application policies, and cryptography settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-template.html#cfn-pcaconnectorad-template-definition
	//
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// Name of the templates.
	//
	// Template names must be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-template.html#cfn-pcaconnectorad-template-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// This setting allows the major version of a template to be increased automatically.
	//
	// All members of Active Directory groups that are allowed to enroll with a template will receive a new certificate issued using that template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-template.html#cfn-pcaconnectorad-template-reenrollallcertificateholders
	//
	ReenrollAllCertificateHolders interface{} `field:"optional" json:"reenrollAllCertificateHolders" yaml:"reenrollAllCertificateHolders"`
	// Metadata assigned to a template consisting of a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-template.html#cfn-pcaconnectorad-template-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

