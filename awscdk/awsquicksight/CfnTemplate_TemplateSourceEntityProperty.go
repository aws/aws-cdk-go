package awsquicksight


// The source entity of the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateSourceEntityProperty := &TemplateSourceEntityProperty{
//   	SourceAnalysis: &TemplateSourceAnalysisProperty{
//   		Arn: jsii.String("arn"),
//   		DataSetReferences: []interface{}{
//   			&DataSetReferenceProperty{
//   				DataSetArn: jsii.String("dataSetArn"),
//   				DataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   			},
//   		},
//   	},
//   	SourceTemplate: &TemplateSourceTemplateProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templatesourceentity.html
//
type CfnTemplate_TemplateSourceEntityProperty struct {
	// The source analysis, if it is based on an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templatesourceentity.html#cfn-quicksight-template-templatesourceentity-sourceanalysis
	//
	SourceAnalysis interface{} `field:"optional" json:"sourceAnalysis" yaml:"sourceAnalysis"`
	// The source template, if it is based on an template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templatesourceentity.html#cfn-quicksight-template-templatesourceentity-sourcetemplate
	//
	SourceTemplate interface{} `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

