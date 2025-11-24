package mixinsawscleanrooms


// The analysis source metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisSourceMetadataProperty := &AnalysisSourceMetadataProperty{
//   	Artifacts: &AnalysisTemplateArtifactMetadataProperty{
//   		AdditionalArtifactHashes: []interface{}{
//   			&HashProperty{
//   				Sha256: jsii.String("sha256"),
//   			},
//   		},
//   		EntryPointHash: &HashProperty{
//   			Sha256: jsii.String("sha256"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysissourcemetadata.html
//
type CfnAnalysisTemplatePropsMixin_AnalysisSourceMetadataProperty struct {
	// The artifacts of the analysis source metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysissourcemetadata.html#cfn-cleanrooms-analysistemplate-analysissourcemetadata-artifacts
	//
	Artifacts interface{} `field:"optional" json:"artifacts" yaml:"artifacts"`
}

