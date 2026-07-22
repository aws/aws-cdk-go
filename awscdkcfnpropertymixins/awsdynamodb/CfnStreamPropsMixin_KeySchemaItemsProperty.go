package awsdynamodb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   keySchemaItemsProperty := &KeySchemaItemsProperty{
//   	AttributeName: jsii.String("attributeName"),
//   	KeyType: jsii.String("keyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-stream-keyschemaitems.html
//
type CfnStreamPropsMixin_KeySchemaItemsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-stream-keyschemaitems.html#cfn-dynamodb-stream-keyschemaitems-attributename
	//
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-stream-keyschemaitems.html#cfn-dynamodb-stream-keyschemaitems-keytype
	//
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

