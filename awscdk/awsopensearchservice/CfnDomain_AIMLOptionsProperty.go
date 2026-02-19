package awsopensearchservice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIMLOptionsProperty := &AIMLOptionsProperty{
//   	S3VectorsEngine: &S3VectorsEngineProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	ServerlessVectorAcceleration: &ServerlessVectorAccelerationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-aimloptions.html
//
type CfnDomain_AIMLOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-aimloptions.html#cfn-opensearchservice-domain-aimloptions-s3vectorsengine
	//
	S3VectorsEngine interface{} `field:"optional" json:"s3VectorsEngine" yaml:"s3VectorsEngine"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-aimloptions.html#cfn-opensearchservice-domain-aimloptions-serverlessvectoracceleration
	//
	ServerlessVectorAcceleration interface{} `field:"optional" json:"serverlessVectorAcceleration" yaml:"serverlessVectorAcceleration"`
}

