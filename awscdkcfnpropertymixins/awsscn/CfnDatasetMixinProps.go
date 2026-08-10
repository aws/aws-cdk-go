package awsscn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDatasetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDatasetMixinProps := &CfnDatasetMixinProps{
//   	Description: jsii.String("description"),
//   	InstanceId: jsii.String("instanceId"),
//   	Name: jsii.String("name"),
//   	Namespace: jsii.String("namespace"),
//   	PartitionSpec: &PartitionSpecProperty{
//   		Fields: []interface{}{
//   			&DataLakeDatasetPartitionFieldProperty{
//   				Name: jsii.String("name"),
//   				Transform: &TransformProperty{
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	Schema: &SchemaProperty{
//   		Fields: []interface{}{
//   			&DataLakeDatasetSchemaFieldProperty{
//   				IsRequired: jsii.Boolean(false),
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		PrimaryKeys: []interface{}{
//   			&DataLakeDatasetPrimaryKeyFieldProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataset.html
//
type CfnDatasetMixinProps struct {
	// The description of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataset.html#cfn-scn-dataset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Web Services Supply Chain instance identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataset.html#cfn-scn-dataset-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The name of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataset.html#cfn-scn-dataset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataset.html#cfn-scn-dataset-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The partition specification of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataset.html#cfn-scn-dataset-partitionspec
	//
	PartitionSpec interface{} `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// The schema of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataset.html#cfn-scn-dataset-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// The tags for the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataset.html#cfn-scn-dataset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

