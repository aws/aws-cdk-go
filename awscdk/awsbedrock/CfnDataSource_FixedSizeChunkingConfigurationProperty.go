package awsbedrock


// Configurations for when you choose fixed-size chunking.
//
// If you set the `chunkingStrategy` as `NONE` , exclude this field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fixedSizeChunkingConfigurationProperty := &FixedSizeChunkingConfigurationProperty{
//   	MaxTokens: jsii.Number(123),
//   	OverlapPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-fixedsizechunkingconfiguration.html
//
type CfnDataSource_FixedSizeChunkingConfigurationProperty struct {
	// The maximum number of tokens to include in a chunk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-fixedsizechunkingconfiguration.html#cfn-bedrock-datasource-fixedsizechunkingconfiguration-maxtokens
	//
	MaxTokens *float64 `field:"required" json:"maxTokens" yaml:"maxTokens"`
	// The percentage of overlap between adjacent chunks of a data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-fixedsizechunkingconfiguration.html#cfn-bedrock-datasource-fixedsizechunkingconfiguration-overlappercentage
	//
	OverlapPercentage *float64 `field:"required" json:"overlapPercentage" yaml:"overlapPercentage"`
}

