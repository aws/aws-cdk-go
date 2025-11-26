package previewawscleanroomsmixins


// The analysis template artifact metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisTemplateArtifactMetadataProperty := &AnalysisTemplateArtifactMetadataProperty{
//   	AdditionalArtifactHashes: []interface{}{
//   		&HashProperty{
//   			Sha256: jsii.String("sha256"),
//   		},
//   	},
//   	EntryPointHash: &HashProperty{
//   		Sha256: jsii.String("sha256"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifactmetadata.html
//
type CfnAnalysisTemplatePropsMixin_AnalysisTemplateArtifactMetadataProperty struct {
	// Additional artifact hashes for the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifactmetadata.html#cfn-cleanrooms-analysistemplate-analysistemplateartifactmetadata-additionalartifacthashes
	//
	AdditionalArtifactHashes interface{} `field:"optional" json:"additionalArtifactHashes" yaml:"additionalArtifactHashes"`
	// The hash of the entry point for the analysis template artifact metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysistemplateartifactmetadata.html#cfn-cleanrooms-analysistemplate-analysistemplateartifactmetadata-entrypointhash
	//
	EntryPointHash interface{} `field:"optional" json:"entryPointHash" yaml:"entryPointHash"`
}

