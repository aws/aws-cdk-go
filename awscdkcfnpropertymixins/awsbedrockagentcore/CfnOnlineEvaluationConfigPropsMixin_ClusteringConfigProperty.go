package awsbedrockagentcore


// The configuration for clustering analysis of evaluation results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   clusteringConfigProperty := &ClusteringConfigProperty{
//   	Frequencies: []*string{
//   		jsii.String("frequencies"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-clusteringconfig.html
//
type CfnOnlineEvaluationConfigPropsMixin_ClusteringConfigProperty struct {
	// The list of frequencies at which clustering reports are generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-clusteringconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-clusteringconfig-frequencies
	//
	Frequencies *[]*string `field:"optional" json:"frequencies" yaml:"frequencies"`
}

