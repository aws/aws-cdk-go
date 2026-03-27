package awscleanroomsml

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConfiguredModelAlgorithm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfiguredModelAlgorithmProps := &CfnConfiguredModelAlgorithmProps{
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	InferenceContainerConfig: &InferenceContainerConfigProperty{
//   		ImageUri: jsii.String("imageUri"),
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrainingContainerConfig: &ContainerConfigProperty{
//   		ImageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		Arguments: []*string{
//   			jsii.String("arguments"),
//   		},
//   		Entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		MetricDefinitions: []interface{}{
//   			&MetricDefinitionProperty{
//   				Name: jsii.String("name"),
//   				Regex: jsii.String("regex"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithm.html
//
type CfnConfiguredModelAlgorithmProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithm.html#cfn-cleanroomsml-configuredmodelalgorithm-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithm.html#cfn-cleanroomsml-configuredmodelalgorithm-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithm.html#cfn-cleanroomsml-configuredmodelalgorithm-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithm.html#cfn-cleanroomsml-configuredmodelalgorithm-inferencecontainerconfig
	//
	InferenceContainerConfig interface{} `field:"optional" json:"inferenceContainerConfig" yaml:"inferenceContainerConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithm.html#cfn-cleanroomsml-configuredmodelalgorithm-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms-ml configured model algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithm.html#cfn-cleanroomsml-configuredmodelalgorithm-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithm.html#cfn-cleanroomsml-configuredmodelalgorithm-trainingcontainerconfig
	//
	TrainingContainerConfig interface{} `field:"optional" json:"trainingContainerConfig" yaml:"trainingContainerConfig"`
}

