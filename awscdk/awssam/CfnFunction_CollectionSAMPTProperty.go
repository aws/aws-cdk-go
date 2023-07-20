package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   collectionSAMPTProperty := &CollectionSAMPTProperty{
//   	CollectionId: jsii.String("collectionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-collectionsampt.html
//
type CfnFunction_CollectionSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-collectionsampt.html#cfn-serverless-function-collectionsampt-collectionid
	//
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
}

