package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery"
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
// Experimental.
type CloudMapOptions struct {
	// The service discovery namespace for the Cloud Map service to attach to the ECS service.
	// Experimental.
	CloudMapNamespace awsservicediscovery.INamespace `field:"optional" json:"cloudMapNamespace" yaml:"cloudMapNamespace"`
	// The container to point to for a SRV record.
	// Experimental.
	Container ContainerDefinition `field:"optional" json:"container" yaml:"container"`
	// The port to point to for a SRV record.
	// Experimental.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The DNS record type that you want AWS Cloud Map to create.
	//
	// The supported record types are A or SRV.
	// Experimental.
	DnsRecordType awsservicediscovery.DnsRecordType `field:"optional" json:"dnsRecordType" yaml:"dnsRecordType"`
	// The amount of time that you want DNS resolvers to cache the settings for this record.
	// Experimental.
	DnsTtl awscdk.Duration `field:"optional" json:"dnsTtl" yaml:"dnsTtl"`
	// The number of 30-second intervals that you want Cloud Map to wait after receiving an UpdateInstanceCustomHealthStatus request before it changes the health status of a service instance.
	//
	// NOTE: This is used for HealthCheckCustomConfig.
	// Experimental.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// The name of the Cloud Map service to attach to the ECS service.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

