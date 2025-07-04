package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
)

// The options to enabling AWS Cloud Map for an Amazon ECS service.
//
// Example:
//   var taskDefinition taskDefinition
//   var cluster cluster
//
//
//   service := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	CloudMapOptions: &CloudMapOptions{
//   		// Create A records - useful for AWSVPC network mode.
//   		DnsRecordType: cloudmap.DnsRecordType_A,
//   	},
//   })
//
type CloudMapOptions struct {
	// The service discovery namespace for the Cloud Map service to attach to the ECS service.
	// Default: - the defaultCloudMapNamespace associated to the cluster.
	//
	CloudMapNamespace awsservicediscovery.INamespace `field:"optional" json:"cloudMapNamespace" yaml:"cloudMapNamespace"`
	// The container to point to for a SRV record.
	// Default: - the task definition's default container.
	//
	Container ContainerDefinition `field:"optional" json:"container" yaml:"container"`
	// The port to point to for a SRV record.
	// Default: - the default port of the task definition's default container.
	//
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The DNS record type that you want AWS Cloud Map to create.
	//
	// The supported record types are A or SRV.
	// Default: - DnsRecordType.A if TaskDefinition.networkMode = AWS_VPC, otherwise DnsRecordType.SRV
	//
	DnsRecordType awsservicediscovery.DnsRecordType `field:"optional" json:"dnsRecordType" yaml:"dnsRecordType"`
	// The amount of time that you want DNS resolvers to cache the settings for this record.
	// Default: Duration.minutes(1)
	//
	DnsTtl awscdk.Duration `field:"optional" json:"dnsTtl" yaml:"dnsTtl"`
	// The number of 30-second intervals that you want Cloud Map to wait after receiving an UpdateInstanceCustomHealthStatus request before it changes the health status of a service instance.
	//
	// NOTE: This is used for HealthCheckCustomConfig.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// The name of the Cloud Map service to attach to the ECS service.
	// Default: CloudFormation-generated name.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

