package awsiot


// Properties for CfnIndexPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnIndexMixinProps := &CfnIndexMixinProps{
//   	IndexName: jsii.String("indexName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-index.html
//
type CfnIndexMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-index.html#cfn-iot-index-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
}

