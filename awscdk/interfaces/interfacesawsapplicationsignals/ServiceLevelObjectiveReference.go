package interfacesawsapplicationsignals


// A reference to a ServiceLevelObjective resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceLevelObjectiveReference := &ServiceLevelObjectiveReference{
//   	ServiceLevelObjectiveArn: jsii.String("serviceLevelObjectiveArn"),
//   }
//
type ServiceLevelObjectiveReference struct {
	// The Arn of the ServiceLevelObjective resource.
	ServiceLevelObjectiveArn *string `field:"required" json:"serviceLevelObjectiveArn" yaml:"serviceLevelObjectiveArn"`
}

