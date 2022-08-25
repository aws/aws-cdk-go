package awssam


// Properties for defining a `CfnSimpleTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSimpleTableProps := &cfnSimpleTableProps{
//   	primaryKey: &primaryKeyProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   	},
//   	provisionedThroughput: &provisionedThroughputProperty{
//   		writeCapacityUnits: jsii.Number(123),
//
//   		// the properties below are optional
//   		readCapacityUnits: jsii.Number(123),
//   	},
//   	sseSpecification: &sSESpecificationProperty{
//   		sseEnabled: jsii.Boolean(false),
//   	},
//   	tableName: jsii.String("tableName"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSimpleTableProps struct {
	// `AWS::Serverless::SimpleTable.PrimaryKey`.
	PrimaryKey interface{} `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// `AWS::Serverless::SimpleTable.ProvisionedThroughput`.
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// `AWS::Serverless::SimpleTable.SSESpecification`.
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// `AWS::Serverless::SimpleTable.TableName`.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// `AWS::Serverless::SimpleTable.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

