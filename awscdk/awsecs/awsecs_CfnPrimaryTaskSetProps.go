package awsecs


// Properties for defining a `CfnPrimaryTaskSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPrimaryTaskSetProps := &cfnPrimaryTaskSetProps{
//   	cluster: jsii.String("cluster"),
//   	service: jsii.String("service"),
//   	taskSetId: jsii.String("taskSetId"),
//   }
//
type CfnPrimaryTaskSetProps struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service that the task set exists in.
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The short name or full Amazon Resource Name (ARN) of the service that the task set exists in.
	Service *string `field:"required" json:"service" yaml:"service"`
	// The short name or full Amazon Resource Name (ARN) of the task set to set as the primary task set in the deployment.
	TaskSetId *string `field:"required" json:"taskSetId" yaml:"taskSetId"`
}

