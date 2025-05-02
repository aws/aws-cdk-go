package awsbedrock


// Settings for generating data from documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentStandardExtractionProperty := &DocumentStandardExtractionProperty{
//   	BoundingBox: &DocumentBoundingBoxProperty{
//   		State: jsii.String("state"),
//   	},
//   	Granularity: &DocumentExtractionGranularityProperty{
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentstandardextraction.html
//
type CfnDataAutomationProject_DocumentStandardExtractionProperty struct {
	// Whether to generate bounding boxes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentstandardextraction.html#cfn-bedrock-dataautomationproject-documentstandardextraction-boundingbox
	//
	BoundingBox interface{} `field:"required" json:"boundingBox" yaml:"boundingBox"`
	// Which granularities to generate data for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentstandardextraction.html#cfn-bedrock-dataautomationproject-documentstandardextraction-granularity
	//
	Granularity interface{} `field:"required" json:"granularity" yaml:"granularity"`
}

