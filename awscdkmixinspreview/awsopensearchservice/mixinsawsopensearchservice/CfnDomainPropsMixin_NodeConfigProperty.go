package mixinsawsopensearchservice


// Configuration options for defining the setup of any node type within the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   nodeConfigProperty := &NodeConfigProperty{
//   	Count: jsii.Number(123),
//   	Enabled: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodeconfig.html
//
type CfnDomainPropsMixin_NodeConfigProperty struct {
	// The number of nodes of a specific type within the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodeconfig.html#cfn-opensearchservice-domain-nodeconfig-count
	//
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// A boolean value indicating whether a specific node type is active or inactive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodeconfig.html#cfn-opensearchservice-domain-nodeconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The instance type of a particular node within the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodeconfig.html#cfn-opensearchservice-domain-nodeconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

