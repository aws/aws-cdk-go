package interfacesawsbatch


// A reference to a JobDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobDefinitionReference := &JobDefinitionReference{
//   	JobDefinitionArn: jsii.String("jobDefinitionArn"),
//   }
//
type JobDefinitionReference struct {
	// The JobDefinitionArn of the JobDefinition resource.
	JobDefinitionArn *string `field:"required" json:"jobDefinitionArn" yaml:"jobDefinitionArn"`
}

