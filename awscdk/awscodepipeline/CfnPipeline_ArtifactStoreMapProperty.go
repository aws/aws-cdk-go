package awscodepipeline


// A mapping of `artifactStore` objects and their corresponding AWS Regions.
//
// There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.
//
// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactStoreMapProperty := &ArtifactStoreMapProperty{
//   	ArtifactStore: &ArtifactStoreProperty{
//   		Location: jsii.String("location"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		EncryptionKey: &EncryptionKeyProperty{
//   			Id: jsii.String("id"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstoremap.html
//
type CfnPipeline_ArtifactStoreMapProperty struct {
	// Represents information about the S3 bucket where artifacts are stored for the pipeline.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstoremap.html#cfn-codepipeline-pipeline-artifactstoremap-artifactstore
	//
	ArtifactStore interface{} `field:"required" json:"artifactStore" yaml:"artifactStore"`
	// The action declaration's AWS Region, such as us-east-1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstoremap.html#cfn-codepipeline-pipeline-artifactstoremap-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
}

