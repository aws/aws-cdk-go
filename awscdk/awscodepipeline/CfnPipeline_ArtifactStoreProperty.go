package awscodepipeline


// The S3 bucket where artifacts for the pipeline are stored.
//
// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactStoreProperty := &ArtifactStoreProperty{
//   	Location: jsii.String("location"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	EncryptionKey: &EncryptionKeyProperty{
//   		Id: jsii.String("id"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html
//
type CfnPipeline_ArtifactStoreProperty struct {
	// The S3 bucket used for storing the artifacts for a pipeline.
	//
	// You can specify the name of an S3 bucket but not a folder in the bucket. A folder to contain the pipeline artifacts is created for you based on the name of the pipeline. You can use any S3 bucket in the same AWS Region as the pipeline to store your pipeline artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-location
	//
	Location *string `field:"required" json:"location" yaml:"location"`
	// The type of the artifact store, such as S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The encryption key used to encrypt the data in the artifact store, such as an AWS Key Management Service ( AWS KMS) key.
	//
	// If this is undefined, the default key for Amazon S3 is used. To see an example artifact store encryption key field, see the example structure here: [AWS::CodePipeline::Pipeline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey
	//
	EncryptionKey interface{} `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

