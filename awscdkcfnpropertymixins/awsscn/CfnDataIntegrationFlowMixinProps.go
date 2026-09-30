package awsscn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDataIntegrationFlowPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDataIntegrationFlowMixinProps := &CfnDataIntegrationFlowMixinProps{
//   	InstanceId: jsii.String("instanceId"),
//   	Name: jsii.String("name"),
//   	Sources: []interface{}{
//   		&DataIntegrationFlowSourceProperty{
//   			DatasetSource: &DatasetSourceConfigurationProperty{
//   				DatasetIdentifier: jsii.String("datasetIdentifier"),
//   				Options: &DatasetOptionsProperty{
//   					DedupeRecords: jsii.Boolean(false),
//   					DedupeStrategy: &DedupeStrategyProperty{
//   						FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   							Fields: []interface{}{
//   								&FieldPriorityDedupeFieldProperty{
//   									Name: jsii.String("name"),
//   									SortOrder: jsii.String("sortOrder"),
//   								},
//   							},
//   						},
//   						Type: jsii.String("type"),
//   					},
//   					LoadType: jsii.String("loadType"),
//   				},
//   			},
//   			S3Source: &S3SourceConfigurationProperty{
//   				BucketName: jsii.String("bucketName"),
//   				Options: &S3OptionsProperty{
//   					FileType: jsii.String("fileType"),
//   				},
//   				Prefix: jsii.String("prefix"),
//   			},
//   			SourceName: jsii.String("sourceName"),
//   			SourceType: jsii.String("sourceType"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Target: &DataIntegrationFlowTargetProperty{
//   		DatasetTarget: &DatasetTargetConfigurationProperty{
//   			DatasetIdentifier: jsii.String("datasetIdentifier"),
//   			Options: &DatasetOptionsProperty{
//   				DedupeRecords: jsii.Boolean(false),
//   				DedupeStrategy: &DedupeStrategyProperty{
//   					FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   						Fields: []interface{}{
//   							&FieldPriorityDedupeFieldProperty{
//   								Name: jsii.String("name"),
//   								SortOrder: jsii.String("sortOrder"),
//   							},
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   				LoadType: jsii.String("loadType"),
//   			},
//   		},
//   		TargetType: jsii.String("targetType"),
//   	},
//   	Transformation: &DataIntegrationFlowTransformationProperty{
//   		SqlTransformation: &SqlTransformationConfigurationProperty{
//   			Query: jsii.String("query"),
//   		},
//   		TransformationType: jsii.String("transformationType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html
//
type CfnDataIntegrationFlowMixinProps struct {
	// The Amazon Web Services Supply Chain instance identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The name of the DataIntegrationFlow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The source configurations for the DataIntegrationFlow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// The tags for the DataIntegrationFlow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The DataIntegrationFlow target parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-target
	//
	Target interface{} `field:"optional" json:"target" yaml:"target"`
	// The DataIntegrationFlow transformation parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html#cfn-scn-dataintegrationflow-transformation
	//
	Transformation interface{} `field:"optional" json:"transformation" yaml:"transformation"`
}

