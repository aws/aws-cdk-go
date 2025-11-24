package mixinsawsbedrock


// Granularity settings for documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   documentExtractionGranularityProperty := &DocumentExtractionGranularityProperty{
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentextractiongranularity.html
//
type CfnDataAutomationProjectPropsMixin_DocumentExtractionGranularityProperty struct {
	// Granularity settings for documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentextractiongranularity.html#cfn-bedrock-dataautomationproject-documentextractiongranularity-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

