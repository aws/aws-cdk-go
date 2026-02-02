package awsemrcontainers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSecurityConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityConfigurationProps := &CfnSecurityConfigurationProps{
//   	SecurityConfigurationData: &SecurityConfigurationDataProperty{
//   		AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   			IamConfiguration: map[string]*string{
//   				"systemRole": jsii.String("systemRole"),
//   			},
//   			IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   				EnableIdentityCenter: jsii.Boolean(false),
//   				IdentityCenterApplicationAssignmentRequired: jsii.Boolean(false),
//   				IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   			},
//   		},
//   		AuthorizationConfiguration: &AuthorizationConfigurationProperty{
//   			LakeFormationConfiguration: &LakeFormationConfigurationProperty{
//   				AuthorizedSessionTagValue: jsii.String("authorizedSessionTagValue"),
//   				QueryAccessControlEnabled: jsii.Boolean(false),
//   				QueryEngineRoleArn: jsii.String("queryEngineRoleArn"),
//   				SecureNamespaceInfo: &SecureNamespaceInfoProperty{
//   					ClusterId: jsii.String("clusterId"),
//   					Namespace: jsii.String("namespace"),
//   				},
//   			},
//   		},
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			AtRestEncryptionConfiguration: &AtRestEncryptionConfigurationProperty{
//   				LocalDiskEncryptionConfiguration: &LocalDiskEncryptionConfigurationProperty{
//   					AwsKmsKeyId: jsii.String("awsKmsKeyId"),
//   					EncryptionKeyProviderType: jsii.String("encryptionKeyProviderType"),
//   				},
//   				S3EncryptionConfiguration: &S3EncryptionConfigurationProperty{
//   					EncryptionOption: jsii.String("encryptionOption"),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   				},
//   			},
//   			InTransitEncryptionConfiguration: &InTransitEncryptionConfigurationProperty{
//   				TlsCertificateConfiguration: &TLSCertificateConfigurationProperty{
//   					CertificateProviderType: jsii.String("certificateProviderType"),
//   					PrivateKeySecretArn: jsii.String("privateKeySecretArn"),
//   					PublicKeySecretArn: jsii.String("publicKeySecretArn"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ContainerProvider: &ContainerProviderProperty{
//   		Id: jsii.String("id"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Info: &ContainerInfoProperty{
//   			EksInfo: &EksInfoProperty{
//   				Namespace: jsii.String("namespace"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-securityconfiguration.html
//
type CfnSecurityConfigurationProps struct {
	// Security configuration data containing encryption and authorization settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-securityconfiguration.html#cfn-emrcontainers-securityconfiguration-securityconfigurationdata
	//
	SecurityConfigurationData interface{} `field:"required" json:"securityConfigurationData" yaml:"securityConfigurationData"`
	// Container provider information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-securityconfiguration.html#cfn-emrcontainers-securityconfiguration-containerprovider
	//
	ContainerProvider interface{} `field:"optional" json:"containerProvider" yaml:"containerProvider"`
	// The name of the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-securityconfiguration.html#cfn-emrcontainers-securityconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-securityconfiguration.html#cfn-emrcontainers-securityconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

