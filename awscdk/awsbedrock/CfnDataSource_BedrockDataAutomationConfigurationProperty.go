package awsbedrock


// Contains configurations for using Amazon Bedrock Data Automation as the parser for ingesting your data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bedrockDataAutomationConfigurationProperty := &BedrockDataAutomationConfigurationProperty{
//   	ParsingModality: jsii.String("parsingModality"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockdataautomationconfiguration.html
//
type CfnDataSource_BedrockDataAutomationConfigurationProperty struct {
	// Specifies whether to enable parsing of multimodal data, including both text and/or images.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockdataautomationconfiguration.html#cfn-bedrock-datasource-bedrockdataautomationconfiguration-parsingmodality
	//
	ParsingModality *string `field:"optional" json:"parsingModality" yaml:"parsingModality"`
}

