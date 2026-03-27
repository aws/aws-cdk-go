package awsopensearchserverless


// Vector search configuration options for the collection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   vectorOptionsProperty := &VectorOptionsProperty{
//   	ServerlessVectorAcceleration: jsii.String("serverlessVectorAcceleration"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-vectoroptions.html
//
type CfnCollectionPropsMixin_VectorOptionsProperty struct {
	// Indicates whether GPU acceleration is enabled for vector indexing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-vectoroptions.html#cfn-opensearchserverless-collection-vectoroptions-serverlessvectoracceleration
	//
	ServerlessVectorAcceleration *string `field:"optional" json:"serverlessVectorAcceleration" yaml:"serverlessVectorAcceleration"`
}

