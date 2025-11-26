package previewawscleanroomsmixins


// The analysis template artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisTemplateArtifactProperty := &AnalysisTemplateArtifactProperty{
//   	Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifact.html
//
type CfnAnalysisTemplatePropsMixin_AnalysisTemplateArtifactProperty struct {
	// The artifact location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifact.html#cfn-cleanrooms-analysistemplate-analysistemplateartifact-location
	//
	Location interface{} `field:"optional" json:"location" yaml:"location"`
}

