package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnArtifactPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnArtifactMixinProps := &CfnArtifactMixinProps{
//   	ArtifactName: jsii.String("artifactName"),
//   	ArtifactType: jsii.String("artifactType"),
//   	MetadataProperties: &MetadataPropertiesProperty{
//   		CommitId: jsii.String("commitId"),
//   		GeneratedBy: jsii.String("generatedBy"),
//   		ProjectId: jsii.String("projectId"),
//   		Repository: jsii.String("repository"),
//   	},
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   	Source: &ArtifactSourceProperty{
//   		SourceTypes: []interface{}{
//   			&ArtifactSourceTypeProperty{
//   				SourceIdType: jsii.String("sourceIdType"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		SourceUri: jsii.String("sourceUri"),
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
type CfnArtifactMixinProps struct {
	// The name of the artifact.
	//
	// Must be unique to your account in an AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-artifactname
	//
	ArtifactName *string `field:"optional" json:"artifactName" yaml:"artifactName"`
	// The artifact type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-artifacttype
	//
	ArtifactType *string `field:"optional" json:"artifactType" yaml:"artifactType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-metadataproperties
	//
	MetadataProperties interface{} `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// A list of properties to add to the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// A list of tags to apply to the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-artifact.html#cfn-sagemaker-artifact-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

