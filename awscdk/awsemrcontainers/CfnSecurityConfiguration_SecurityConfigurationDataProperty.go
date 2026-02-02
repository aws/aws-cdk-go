package awsemrcontainers


// Security configuration data containing encryption and authorization settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityConfigurationDataProperty := &SecurityConfigurationDataProperty{
//   	AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   		IamConfiguration: map[string]*string{
//   			"systemRole": jsii.String("systemRole"),
//   		},
//   		IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   			EnableIdentityCenter: jsii.Boolean(false),
//   			IdentityCenterApplicationAssignmentRequired: jsii.Boolean(false),
//   			IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   		},
//   	},
//   	AuthorizationConfiguration: &AuthorizationConfigurationProperty{
//   		LakeFormationConfiguration: &LakeFormationConfigurationProperty{
//   			AuthorizedSessionTagValue: jsii.String("authorizedSessionTagValue"),
//   			QueryAccessControlEnabled: jsii.Boolean(false),
//   			QueryEngineRoleArn: jsii.String("queryEngineRoleArn"),
//   			SecureNamespaceInfo: &SecureNamespaceInfoProperty{
//   				ClusterId: jsii.String("clusterId"),
//   				Namespace: jsii.String("namespace"),
//   			},
//   		},
//   	},
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		AtRestEncryptionConfiguration: &AtRestEncryptionConfigurationProperty{
//   			LocalDiskEncryptionConfiguration: &LocalDiskEncryptionConfigurationProperty{
//   				AwsKmsKeyId: jsii.String("awsKmsKeyId"),
//   				EncryptionKeyProviderType: jsii.String("encryptionKeyProviderType"),
//   			},
//   			S3EncryptionConfiguration: &S3EncryptionConfigurationProperty{
//   				EncryptionOption: jsii.String("encryptionOption"),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   			},
//   		},
//   		InTransitEncryptionConfiguration: &InTransitEncryptionConfigurationProperty{
//   			TlsCertificateConfiguration: &TLSCertificateConfigurationProperty{
//   				CertificateProviderType: jsii.String("certificateProviderType"),
//   				PrivateKeySecretArn: jsii.String("privateKeySecretArn"),
//   				PublicKeySecretArn: jsii.String("publicKeySecretArn"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-securityconfigurationdata.html
//
type CfnSecurityConfiguration_SecurityConfigurationDataProperty struct {
	// Authentication configuration for the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-securityconfigurationdata.html#cfn-emrcontainers-securityconfiguration-securityconfigurationdata-authenticationconfiguration
	//
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Authorization configuration for the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-securityconfigurationdata.html#cfn-emrcontainers-securityconfiguration-securityconfigurationdata-authorizationconfiguration
	//
	AuthorizationConfiguration interface{} `field:"optional" json:"authorizationConfiguration" yaml:"authorizationConfiguration"`
	// Encryption configuration for the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-securityconfigurationdata.html#cfn-emrcontainers-securityconfiguration-securityconfigurationdata-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

