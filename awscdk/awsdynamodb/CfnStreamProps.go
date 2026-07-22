package awsdynamodb


// Properties for defining a `CfnStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamProps := &CfnStreamProps{
//   	StreamViewType: jsii.String("streamViewType"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-stream.html
//
type CfnStreamProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-stream.html#cfn-dynamodb-stream-streamviewtype
	//
	StreamViewType *string `field:"optional" json:"streamViewType" yaml:"streamViewType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-stream.html#cfn-dynamodb-stream-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

