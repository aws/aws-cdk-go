package mixinsawscleanrooms


// The analysis template artifacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisTemplateArtifactsProperty := &AnalysisTemplateArtifactsProperty{
//   	AdditionalArtifacts: []interface{}{
//   		&AnalysisTemplateArtifactProperty{
//   			Location: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   	},
//   	EntryPoint: &AnalysisTemplateArtifactProperty{
//   		Location: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifacts.html
//
type CfnAnalysisTemplatePropsMixin_AnalysisTemplateArtifactsProperty struct {
	// Additional artifacts for the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifacts.html#cfn-cleanrooms-analysistemplate-analysistemplateartifacts-additionalartifacts
	//
	AdditionalArtifacts interface{} `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// The entry point for the analysis template artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifacts.html#cfn-cleanrooms-analysistemplate-analysistemplateartifacts-entrypoint
	//
	EntryPoint interface{} `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The role ARN for the analysis template artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifacts.html#cfn-cleanrooms-analysistemplate-analysistemplateartifacts-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

