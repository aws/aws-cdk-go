package awsecs


// The details for the service registry.
//
// Each service may be associated with one service registry. Multiple service registries for each service are not supported.
//
// When you add, update, or remove the service registries configuration, Amazon ECS starts a new deployment. New tasks are registered and deregistered to the updated service registry configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceRegistryProperty := &serviceRegistryProperty{
//   	containerName: jsii.String("containerName"),
//   	containerPort: jsii.Number(123),
//   	port: jsii.Number(123),
//   	registryArn: jsii.String("registryArn"),
//   }
//
type CfnTaskSet_ServiceRegistryProperty struct {
	// The container name value to be used for your service discovery service.
	//
	// It's already specified in the task definition. If the task definition that your service task specifies uses the `bridge` or `host` network mode, you must specify a `containerName` and `containerPort` combination from the task definition. If the task definition that your service task specifies uses the `awsvpc` network mode and a type SRV DNS record is used, you must specify either a `containerName` and `containerPort` combination or a `port` value. However, you can't specify both.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port value to be used for your service discovery service.
	//
	// It's already specified in the task definition. If the task definition your service task specifies uses the `bridge` or `host` network mode, you must specify a `containerName` and `containerPort` combination from the task definition. If the task definition your service task specifies uses the `awsvpc` network mode and a type SRV DNS record is used, you must specify either a `containerName` and `containerPort` combination or a `port` value. However, you can't specify both.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The port value used if your service discovery service specified an SRV record.
	//
	// This field might be used if both the `awsvpc` network mode and SRV records are used.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The Amazon Resource Name (ARN) of the service registry.
	//
	// The currently supported service registry is AWS Cloud Map . For more information, see [CreateService](https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html) .
	RegistryArn *string `field:"optional" json:"registryArn" yaml:"registryArn"`
}

