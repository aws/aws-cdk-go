package awsrds


// The `ProcessorFeature` property type specifies the processor features of a DB instance class status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorFeatureProperty := &ProcessorFeatureProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-processorfeature.html
//
type CfnDBInstance_ProcessorFeatureProperty struct {
	// The name of the processor feature.
	//
	// Valid names are `coreCount` and `threadsPerCore` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-processorfeature.html#cfn-rds-dbinstance-processorfeature-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of a processor feature name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-processorfeature.html#cfn-rds-dbinstance-processorfeature-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

