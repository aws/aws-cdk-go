package awssam


// Properties for defining a `CfnSimpleTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSimpleTableProps := &CfnSimpleTableProps{
//   	PrimaryKey: &PrimaryKeyProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Name: jsii.String("name"),
//   	},
//   	ProvisionedThroughput: &ProvisionedThroughputProperty{
//   		WriteCapacityUnits: jsii.Number(123),
//
//   		// the properties below are optional
//   		ReadCapacityUnits: jsii.Number(123),
//   	},
//   	SseSpecification: &SSESpecificationProperty{
//   		SseEnabled: jsii.Boolean(false),
//   	},
//   	TableName: jsii.String("tableName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-simpletable.html
//
type CfnSimpleTableProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-simpletable.html#cfn-serverless-simpletable-primarykey
	//
	PrimaryKey interface{} `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-simpletable.html#cfn-serverless-simpletable-provisionedthroughput
	//
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-simpletable.html#cfn-serverless-simpletable-ssespecification
	//
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-simpletable.html#cfn-serverless-simpletable-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-simpletable.html#cfn-serverless-simpletable-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

