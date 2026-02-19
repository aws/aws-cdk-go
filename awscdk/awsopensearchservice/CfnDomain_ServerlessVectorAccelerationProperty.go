package awsopensearchservice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessVectorAccelerationProperty := &ServerlessVectorAccelerationProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-serverlessvectoracceleration.html
//
type CfnDomain_ServerlessVectorAccelerationProperty struct {
	// Whether to enable serverless vector acceleration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-serverlessvectoracceleration.html#cfn-opensearchservice-domain-serverlessvectoracceleration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

