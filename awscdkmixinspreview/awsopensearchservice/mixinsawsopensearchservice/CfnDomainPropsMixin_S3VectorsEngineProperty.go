package mixinsawsopensearchservice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3VectorsEngineProperty := &S3VectorsEngineProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-s3vectorsengine.html
//
type CfnDomainPropsMixin_S3VectorsEngineProperty struct {
	// Whether to enable S3 vectors engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-s3vectorsengine.html#cfn-opensearchservice-domain-s3vectorsengine-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

