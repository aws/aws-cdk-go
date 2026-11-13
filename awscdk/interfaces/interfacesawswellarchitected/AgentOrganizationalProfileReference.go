package interfacesawswellarchitected


// A reference to a AgentOrganizationalProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentOrganizationalProfileReference := &AgentOrganizationalProfileReference{
//   	AgentOrganizationalProfileArn: jsii.String("agentOrganizationalProfileArn"),
//   }
//
type AgentOrganizationalProfileReference struct {
	// The Arn of the AgentOrganizationalProfile resource.
	AgentOrganizationalProfileArn *string `field:"required" json:"agentOrganizationalProfileArn" yaml:"agentOrganizationalProfileArn"`
}

