package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTopicV2PropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTopicV2MixinProps := &CfnTopicV2MixinProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	CustomInstructions: &CustomInstructionsProperty{
//   		CustomInstructionsString: jsii.String("customInstructionsString"),
//   	},
//   	DataSetRelations: []interface{}{
//   		&DataSetRelationProperty{
//   			Left: &DataSetRelationEndpointProperty{
//   				ColumnNames: []*string{
//   					jsii.String("columnNames"),
//   				},
//   				DataSetArn: jsii.String("dataSetArn"),
//   			},
//   			Right: &DataSetRelationEndpointProperty{
//   				ColumnNames: []*string{
//   					jsii.String("columnNames"),
//   				},
//   				DataSetArn: jsii.String("dataSetArn"),
//   			},
//   		},
//   	},
//   	DataSets: []interface{}{
//   		&DataSetReferenceProperty{
//   			DataSetArn: jsii.String("dataSetArn"),
//   			DataSetName: jsii.String("dataSetName"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	FolderArns: []*string{
//   		jsii.String("folderArns"),
//   	},
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TopicId: jsii.String("topicId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html
//
type CfnTopicV2MixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html#cfn-quicksight-topicv2-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html#cfn-quicksight-topicv2-custominstructions
	//
	CustomInstructions interface{} `field:"optional" json:"customInstructions" yaml:"customInstructions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html#cfn-quicksight-topicv2-datasetrelations
	//
	DataSetRelations interface{} `field:"optional" json:"dataSetRelations" yaml:"dataSetRelations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html#cfn-quicksight-topicv2-datasets
	//
	DataSets interface{} `field:"optional" json:"dataSets" yaml:"dataSets"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html#cfn-quicksight-topicv2-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html#cfn-quicksight-topicv2-folderarns
	//
	FolderArns *[]*string `field:"optional" json:"folderArns" yaml:"folderArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html#cfn-quicksight-topicv2-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html#cfn-quicksight-topicv2-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html#cfn-quicksight-topicv2-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topicv2.html#cfn-quicksight-topicv2-topicid
	//
	TopicId *string `field:"optional" json:"topicId" yaml:"topicId"`
}

