package awsquicksight


// The source entity of the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateSourceEntityProperty := &templateSourceEntityProperty{
//   	sourceAnalysis: &templateSourceAnalysisProperty{
//   		arn: jsii.String("arn"),
//   		dataSetReferences: []interface{}{
//   			&dataSetReferenceProperty{
//   				dataSetArn: jsii.String("dataSetArn"),
//   				dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   			},
//   		},
//   	},
//   	sourceTemplate: &templateSourceTemplateProperty{
//   		arn: jsii.String("arn"),
//   	},
//   }
//
type CfnTemplate_TemplateSourceEntityProperty struct {
	// The source analysis, if it is based on an analysis.
	SourceAnalysis interface{} `field:"optional" json:"sourceAnalysis" yaml:"sourceAnalysis"`
	// The source template, if it is based on an template.
	SourceTemplate interface{} `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

