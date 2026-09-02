package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnArtifact`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnArtifactProps := &CfnArtifactProps{
//   	ArtifactType: jsii.String("artifactType"),
//   	Source: &ArtifactSourceProperty{
//   		SourceUri: jsii.String("sourceUri"),
//
//   		// the properties below are optional
//   		SourceTypes: []interface{}{
//   			&ArtifactSourceTypeProperty{
//   				SourceIdType: jsii.String("sourceIdType"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ArtifactName: jsii.String("artifactName"),
//   	MetadataProperties: &MetadataPropertiesProperty{
//   		CommitId: jsii.String("commitId"),
//   		GeneratedBy: jsii.String("generatedBy"),
//   		ProjectId: jsii.String("projectId"),
//   		Repository: jsii.String("repository"),
//   	},
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html
//
type CfnArtifactProps struct {
	// The artifact type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-artifacttype
	//
	ArtifactType *string `field:"required" json:"artifactType" yaml:"artifactType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// The name of the artifact.
	//
	// Must be unique to your account in an AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-artifactname
	//
	ArtifactName *string `field:"optional" json:"artifactName" yaml:"artifactName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-metadataproperties
	//
	MetadataProperties interface{} `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// A list of properties to add to the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// A list of tags to apply to the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

