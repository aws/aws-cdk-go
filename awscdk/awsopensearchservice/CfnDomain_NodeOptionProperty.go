package awsopensearchservice


// Container for specifying node type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeOptionProperty := &NodeOptionProperty{
//   	NodeConfig: &NodeConfigProperty{
//   		Count: jsii.Number(123),
//   		Enabled: jsii.Boolean(false),
//   		Type: jsii.String("type"),
//   	},
//   	NodeType: jsii.String("nodeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodeoption.html
//
type CfnDomain_NodeOptionProperty struct {
	// Container for specifying configuration of any node type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodeoption.html#cfn-opensearchservice-domain-nodeoption-nodeconfig
	//
	NodeConfig interface{} `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// Container for node type like coordinating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodeoption.html#cfn-opensearchservice-domain-nodeoption-nodetype
	//
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
}

