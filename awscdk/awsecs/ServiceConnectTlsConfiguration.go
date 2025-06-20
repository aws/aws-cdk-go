package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// TLS configuration for Service Connect service.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var kmsKey iKey
//   var role iRole
//
//
//   service := ecs.NewFargateService(this, jsii.String("FargateService"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	ServiceConnectConfiguration: &ServiceConnectProps{
//   		Services: []serviceConnectService{
//   			&serviceConnectService{
//   				Tls: &ServiceConnectTlsConfiguration{
//   					Role: *Role,
//   					KmsKey: *KmsKey,
//   					AwsPcaAuthorityArn: jsii.String("arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/123456789012"),
//   				},
//   				PortMappingName: jsii.String("api"),
//   			},
//   		},
//   		Namespace: jsii.String("sample namespace"),
//   	},
//   })
//
type ServiceConnectTlsConfiguration struct {
	// The ARN of the certificate root authority that secures your service.
	// Default: - none.
	//
	AwsPcaAuthorityArn *string `field:"optional" json:"awsPcaAuthorityArn" yaml:"awsPcaAuthorityArn"`
	// The KMS key used for encryption and decryption.
	// Default: - none.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The IAM role that's associated with the Service Connect TLS.
	// Default: - none.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

