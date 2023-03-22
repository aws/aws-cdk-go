package awskinesisanalyticsv2


// The configuration of connectors and user-defined functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customArtifactConfigurationProperty := &customArtifactConfigurationProperty{
//   	artifactType: jsii.String("artifactType"),
//
//   	// the properties below are optional
//   	mavenReference: &mavenReferenceProperty{
//   		artifactId: jsii.String("artifactId"),
//   		groupId: jsii.String("groupId"),
//   		version: jsii.String("version"),
//   	},
//   	s3ContentLocation: &s3ContentLocationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		fileKey: jsii.String("fileKey"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
type CfnApplication_CustomArtifactConfigurationProperty struct {
	// Set this to either `UDF` or `DEPENDENCY_JAR` .
	//
	// `UDF` stands for user-defined functions. This type of artifact must be in an S3 bucket. A `DEPENDENCY_JAR` can be in either Maven or an S3 bucket.
	ArtifactType *string `field:"required" json:"artifactType" yaml:"artifactType"`
	// The parameters required to fully specify a Maven reference.
	MavenReference interface{} `field:"optional" json:"mavenReference" yaml:"mavenReference"`
	// The location of the custom artifacts.
	S3ContentLocation interface{} `field:"optional" json:"s3ContentLocation" yaml:"s3ContentLocation"`
}

