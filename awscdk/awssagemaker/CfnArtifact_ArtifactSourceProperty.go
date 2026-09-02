package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactSourceProperty := &ArtifactSourceProperty{
//   	SourceUri: jsii.String("sourceUri"),
//
//   	// the properties below are optional
//   	SourceTypes: []interface{}{
//   		&ArtifactSourceTypeProperty{
//   			SourceIdType: jsii.String("sourceIdType"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-artifact-artifactsource.html
//
type CfnArtifact_ArtifactSourceProperty struct {
	// The URI of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-artifact-artifactsource.html#cfn-sagemaker-artifact-artifactsource-sourceuri
	//
	SourceUri *string `field:"required" json:"sourceUri" yaml:"sourceUri"`
	// A list of source types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-artifact-artifactsource.html#cfn-sagemaker-artifact-artifactsource-sourcetypes
	//
	SourceTypes interface{} `field:"optional" json:"sourceTypes" yaml:"sourceTypes"`
}

