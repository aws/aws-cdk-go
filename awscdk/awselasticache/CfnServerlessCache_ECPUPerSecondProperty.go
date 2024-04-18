package awselasticache


// The configuration for the number of ElastiCache Processing Units (ECPU) the cache can consume per second.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eCPUPerSecondProperty := &ECPUPerSecondProperty{
//   	Maximum: jsii.Number(123),
//   	Minimum: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-ecpupersecond.html
//
type CfnServerlessCache_ECPUPerSecondProperty struct {
	// The configuration for the maximum number of ECPUs the cache can consume per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-ecpupersecond.html#cfn-elasticache-serverlesscache-ecpupersecond-maximum
	//
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// The configuration for the minimum number of ECPUs the cache should be able consume per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-ecpupersecond.html#cfn-elasticache-serverlesscache-ecpupersecond-minimum
	//
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

