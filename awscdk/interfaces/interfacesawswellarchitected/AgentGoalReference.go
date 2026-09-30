package interfacesawswellarchitected


// A reference to a AgentGoal resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentGoalReference := &AgentGoalReference{
//   	AgentGoalArn: jsii.String("agentGoalArn"),
//   }
//
type AgentGoalReference struct {
	// The Arn of the AgentGoal resource.
	AgentGoalArn *string `field:"required" json:"agentGoalArn" yaml:"agentGoalArn"`
}

