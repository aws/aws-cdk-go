package awsbatch


// The network configuration for jobs that are running on Fargate resources.
//
// Jobs that are running on EC2 resources must not specify this parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	assignPublicIp: jsii.String("assignPublicIp"),
//   }
//
type CfnJobDefinition_NetworkConfigurationProperty struct {
	// Indicates whether the job should have a public IP address.
	//
	// For a job that is running on Fargate resources in a private subnet to send outbound traffic to the internet (for example, to pull container images), the private subnet requires a NAT gateway be attached to route requests to the internet. For more information, see [Amazon ECS task networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) . The default value is "DISABLED".
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
}

