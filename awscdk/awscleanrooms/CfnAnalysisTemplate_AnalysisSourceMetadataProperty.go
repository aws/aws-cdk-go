package awscleanrooms


// The analysis source metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisSourceMetadataProperty := &AnalysisSourceMetadataProperty{
//   	Artifacts: &AnalysisTemplateArtifactMetadataProperty{
//   		EntryPointHash: &HashProperty{
//   			Sha256: jsii.String("sha256"),
//   		},
//
//   		// the properties below are optional
//   		AdditionalArtifactHashes: []interface{}{
//   			&HashProperty{
//   				Sha256: jsii.String("sha256"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysissourcemetadata.html
//
type CfnAnalysisTemplate_AnalysisSourceMetadataProperty struct {
	// The artifacts of the analysis source metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysissourcemetadata.html#cfn-cleanrooms-analysistemplate-analysissourcemetadata-artifacts
	//
	Artifacts interface{} `field:"required" json:"artifacts" yaml:"artifacts"`
}

