package awsscn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataIntegrationFlow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataIntegrationFlowProps := &CfnDataIntegrationFlowProps{
//   	InstanceId: jsii.String("instanceId"),
//   	Name: jsii.String("name"),
//   	Sources: []interface{}{
//   		&DataIntegrationFlowSourceProperty{
//   			SourceName: jsii.String("sourceName"),
//   			SourceType: jsii.String("sourceType"),
//
//   			// the properties below are optional
//   			DatasetSource: &DatasetSourceConfigurationProperty{
//   				DatasetIdentifier: jsii.String("datasetIdentifier"),
//
//   				// the properties below are optional
//   				Options: &DatasetOptionsProperty{
//   					DedupeRecords: jsii.Boolean(false),
//   					DedupeStrategy: &DedupeStrategyProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   							Fields: []interface{}{
//   								&FieldPriorityDedupeFieldProperty{
//   									Name: jsii.String("name"),
//   									SortOrder: jsii.String("sortOrder"),
//   								},
//   							},
//   						},
//   					},
//   					LoadType: jsii.String("loadType"),
//   				},
//   			},
//   			S3Source: &S3SourceConfigurationProperty{
//   				BucketName: jsii.String("bucketName"),
//   				Prefix: jsii.String("prefix"),
//
//   				// the properties below are optional
//   				Options: &S3OptionsProperty{
//   					FileType: jsii.String("fileType"),
//   				},
//   			},
//   		},
//   	},
//   	Target: &DataIntegrationFlowTargetProperty{
//   		TargetType: jsii.String("targetType"),
//
//   		// the properties below are optional
//   		DatasetTarget: &DatasetTargetConfigurationProperty{
//   			DatasetIdentifier: jsii.String("datasetIdentifier"),
//
//   			// the properties below are optional
//   			Options: &DatasetOptionsProperty{
//   				DedupeRecords: jsii.Boolean(false),
//   				DedupeStrategy: &DedupeStrategyProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   						Fields: []interface{}{
//   							&FieldPriorityDedupeFieldProperty{
//   								Name: jsii.String("name"),
//   								SortOrder: jsii.String("sortOrder"),
//   							},
//   						},
//   					},
//   				},
//   				LoadType: jsii.String("loadType"),
//   			},
//   		},
//   	},
//   	Transformation: &DataIntegrationFlowTransformationProperty{
//   		TransformationType: jsii.String("transformationType"),
//
//   		// the properties below are optional
//   		SqlTransformation: &SqlTransformationConfigurationProperty{
//   			Query: jsii.String("query"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html
//
type CfnDataIntegrationFlowProps struct {
	// The Amazon Web Services Supply Chain instance identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-instanceid
	//
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The name of the DataIntegrationFlow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source configurations for the DataIntegrationFlow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-sources
	//
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// The DataIntegrationFlow target parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-target
	//
	Target interface{} `field:"required" json:"target" yaml:"target"`
	// The DataIntegrationFlow transformation parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-transformation
	//
	Transformation interface{} `field:"required" json:"transformation" yaml:"transformation"`
	// The tags for the DataIntegrationFlow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

