package awssagemaker


// TTL configuration of the feature group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ttlDurationProperty := &TtlDurationProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-ttlduration.html
//
type CfnFeatureGroup_TtlDurationProperty struct {
	// Unit of ttl configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-ttlduration.html#cfn-sagemaker-featuregroup-ttlduration-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// Value of ttl configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-ttlduration.html#cfn-sagemaker-featuregroup-ttlduration-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

