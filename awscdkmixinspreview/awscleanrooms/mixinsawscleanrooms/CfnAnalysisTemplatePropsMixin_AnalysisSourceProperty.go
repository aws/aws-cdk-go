package mixinsawscleanrooms


// The structure that defines the body of the analysis template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisSourceProperty := &AnalysisSourceProperty{
//   	Artifacts: &AnalysisTemplateArtifactsProperty{
//   		AdditionalArtifacts: []interface{}{
//   			&AnalysisTemplateArtifactProperty{
//   				Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   		},
//   		EntryPoint: &AnalysisTemplateArtifactProperty{
//   			Location: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysissource.html
//
type CfnAnalysisTemplatePropsMixin_AnalysisSourceProperty struct {
	// The artifacts of the analysis source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysissource.html#cfn-cleanrooms-analysistemplate-analysissource-artifacts
	//
	Artifacts interface{} `field:"optional" json:"artifacts" yaml:"artifacts"`
	// The query text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysissource.html#cfn-cleanrooms-analysistemplate-analysissource-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
}

