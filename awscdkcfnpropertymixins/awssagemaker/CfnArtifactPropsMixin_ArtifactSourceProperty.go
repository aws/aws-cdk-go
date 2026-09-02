package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   artifactSourceProperty := &ArtifactSourceProperty{
//   	SourceTypes: []interface{}{
//   		&ArtifactSourceTypeProperty{
//   			SourceIdType: jsii.String("sourceIdType"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	SourceUri: jsii.String("sourceUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-artifact-artifactsource.html
//
type CfnArtifactPropsMixin_ArtifactSourceProperty struct {
	// A list of source types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-artifact-artifactsource.html#cfn-sagemaker-artifact-artifactsource-sourcetypes
	//
	SourceTypes interface{} `field:"optional" json:"sourceTypes" yaml:"sourceTypes"`
	// The URI of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-artifact-artifactsource.html#cfn-sagemaker-artifact-artifactsource-sourceuri
	//
	SourceUri *string `field:"optional" json:"sourceUri" yaml:"sourceUri"`
}

