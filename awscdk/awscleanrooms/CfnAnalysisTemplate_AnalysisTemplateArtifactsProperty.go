package awscleanrooms


// The analysis template artifacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisTemplateArtifactsProperty := &AnalysisTemplateArtifactsProperty{
//   	EntryPoint: &AnalysisTemplateArtifactProperty{
//   		Location: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AdditionalArtifacts: []interface{}{
//   		&AnalysisTemplateArtifactProperty{
//   			Location: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifacts.html
//
type CfnAnalysisTemplate_AnalysisTemplateArtifactsProperty struct {
	// The entry point for the analysis template artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifacts.html#cfn-cleanrooms-analysistemplate-analysistemplateartifacts-entrypoint
	//
	EntryPoint interface{} `field:"required" json:"entryPoint" yaml:"entryPoint"`
	// The role ARN for the analysis template artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifacts.html#cfn-cleanrooms-analysistemplate-analysistemplateartifacts-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Additional artifacts for the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifacts.html#cfn-cleanrooms-analysistemplate-analysistemplateartifacts-additionalartifacts
	//
	AdditionalArtifacts interface{} `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
}

