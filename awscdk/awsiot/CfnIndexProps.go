package awsiot


// Properties for defining a `CfnIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIndexProps := &CfnIndexProps{
//   	IndexName: jsii.String("indexName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-index.html
//
type CfnIndexProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-index.html#cfn-iot-index-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
}

