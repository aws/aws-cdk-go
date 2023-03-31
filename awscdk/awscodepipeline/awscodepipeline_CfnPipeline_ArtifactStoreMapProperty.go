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
//   artifactStoreMapProperty := &artifactStoreMapProperty{
//   	artifactStore: &artifactStoreProperty{
//   		location: jsii.String("location"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		encryptionKey: &encryptionKeyProperty{
//   			id: jsii.String("id"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	region: jsii.String("region"),
//   }
//
type CfnPipeline_ArtifactStoreMapProperty struct {
	// Represents information about the S3 bucket where artifacts are stored for the pipeline.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStore interface{} `field:"required" json:"artifactStore" yaml:"artifactStore"`
	// The action declaration's AWS Region, such as us-east-1.
	Region *string `field:"required" json:"region" yaml:"region"`
}

