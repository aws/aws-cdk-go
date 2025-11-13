package interfacesawsecs


// A reference to a TaskSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskSetReference := &TaskSetReference{
//   	Cluster: jsii.String("cluster"),
//   	Service: jsii.String("service"),
//   	TaskSetId: jsii.String("taskSetId"),
//   }
//
type TaskSetReference struct {
	// The Cluster of the TaskSet resource.
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The Service of the TaskSet resource.
	Service *string `field:"required" json:"service" yaml:"service"`
	// The Id of the TaskSet resource.
	TaskSetId *string `field:"required" json:"taskSetId" yaml:"taskSetId"`
}

