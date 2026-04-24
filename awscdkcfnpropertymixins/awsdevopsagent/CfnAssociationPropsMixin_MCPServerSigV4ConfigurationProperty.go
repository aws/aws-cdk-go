package awsdevopsagent


// SigV4-authenticated MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mCPServerSigV4ConfigurationProperty := &MCPServerSigV4ConfigurationProperty{
//   	Tools: []*string{
//   		jsii.String("tools"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserversigv4configuration.html
//
type CfnAssociationPropsMixin_MCPServerSigV4ConfigurationProperty struct {
	// List of MCP tools available for the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserversigv4configuration.html#cfn-devopsagent-association-mcpserversigv4configuration-tools
	//
	Tools *[]*string `field:"optional" json:"tools" yaml:"tools"`
}

