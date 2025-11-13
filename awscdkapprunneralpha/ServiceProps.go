package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Properties of the AppRunner Service.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   service := apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   })
//
//   service.AddToRolePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Effect: iam.Effect_ALLOW,
//   	Actions: []*string{
//   		jsii.String("s3:GetObject"),
//   	},
//   	Resources: []*string{
//   		jsii.String("*"),
//   	},
//   }))
//
// Experimental.
type ServiceProps struct {
	// The source of the repository for the service.
	// Experimental.
	Source Source `field:"required" json:"source" yaml:"source"`
	// The IAM role that grants the App Runner service access to a source repository.
	//
	// It's required for ECR image repositories (but not for ECR Public repositories).
	//
	// The role must be assumable by the 'build.apprunner.amazonaws.com' service principal.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles-service.access
	//
	// Default: - generate a new access role.
	//
	// Experimental.
	AccessRole awsiam.IRole `field:"optional" json:"accessRole" yaml:"accessRole"`
	// Specifies whether to enable continuous integration from the source repository.
	//
	// If true, continuous integration from the source repository is enabled for the App Runner service.
	// Each repository change (including any source code commit or new image version) starts a deployment.
	// By default, App Runner sets to false for a source image that uses an ECR Public repository or an ECR repository that's in an AWS account other than the one that the service is in.
	// App Runner sets to true in all other cases (which currently include a source code repository or a source image using a same-account ECR repository).
	// Default: - no value will be passed.
	//
	// Experimental.
	AutoDeploymentsEnabled *bool `field:"optional" json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// Specifies an App Runner Auto Scaling Configuration.
	//
	// A default configuration is either the AWS recommended configuration,
	// or the configuration you set as the default.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/manage-autoscaling.html
	//
	// Default: - the latest revision of a default auto scaling configuration is used.
	//
	// Experimental.
	AutoScalingConfiguration IAutoScalingConfiguration `field:"optional" json:"autoScalingConfiguration" yaml:"autoScalingConfiguration"`
	// The number of CPU units reserved for each instance of your App Runner service.
	// Default: Cpu.ONE_VCPU
	//
	// Experimental.
	Cpu Cpu `field:"optional" json:"cpu" yaml:"cpu"`
	// Settings for the health check that AWS App Runner performs to monitor the health of a service.
	//
	// You can specify it by static methods `HealthCheck.http` or `HealthCheck.tcp`.
	// Default: - no health check configuration.
	//
	// Experimental.
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The IAM role that provides permissions to your App Runner service.
	//
	// These are permissions that your code needs when it calls any AWS APIs.
	//
	// The role must be assumable by the 'tasks.apprunner.amazonaws.com' service principal.
	// See: https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles-service.instance
	//
	// Default: - generate a new instance role.
	//
	// Experimental.
	InstanceRole awsiam.IRole `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// The IP address type for your incoming public network configuration.
	// Default: - IpAddressType.IPV4
	//
	// Experimental.
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Specifies whether your App Runner service is publicly accessible.
	//
	// If you use `VpcIngressConnection`, you must set this property to `false`.
	// Default: true.
	//
	// Experimental.
	IsPubliclyAccessible *bool `field:"optional" json:"isPubliclyAccessible" yaml:"isPubliclyAccessible"`
	// The customer managed key that AWS App Runner uses to encrypt copies of the source repository and service logs.
	// Default: - Use an AWS managed key.
	//
	// Experimental.
	KmsKey interfacesawskms.IKeyRef `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The amount of memory reserved for each instance of your App Runner service.
	// Default: Memory.TWO_GB
	//
	// Experimental.
	Memory Memory `field:"optional" json:"memory" yaml:"memory"`
	// Settings for an App Runner observability configuration.
	// Default: - no observability configuration resource is associated with the service.
	//
	// Experimental.
	ObservabilityConfiguration IObservabilityConfiguration `field:"optional" json:"observabilityConfiguration" yaml:"observabilityConfiguration"`
	// Name of the service.
	// Default: - auto-generated if undefined.
	//
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Settings for an App Runner VPC connector to associate with the service.
	// Default: - no VPC connector, uses the DEFAULT egress type instead.
	//
	// Experimental.
	VpcConnector IVpcConnector `field:"optional" json:"vpcConnector" yaml:"vpcConnector"`
}

