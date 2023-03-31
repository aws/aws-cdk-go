package awsdatasync


// A list of Amazon Resource Names (ARNs) of agents to use for a Network File System (NFS) location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onPremConfigProperty := &onPremConfigProperty{
//   	agentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   }
//
type CfnLocationNFS_OnPremConfigProperty struct {
	// ARNs of the agents to use for an NFS location.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
}

