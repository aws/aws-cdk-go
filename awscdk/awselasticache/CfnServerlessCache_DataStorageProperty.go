package awselasticache


// The data storage limit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataStorageProperty := &DataStorageProperty{
//   	Unit: jsii.String("unit"),
//
//   	// the properties below are optional
//   	Maximum: jsii.Number(123),
//   	Minimum: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html
//
type CfnServerlessCache_DataStorageProperty struct {
	// The unit that the storage is measured in, in GB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html#cfn-elasticache-serverlesscache-datastorage-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// The upper limit for data storage the cache is set to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html#cfn-elasticache-serverlesscache-datastorage-maximum
	//
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// The lower limit for data storage the cache is set to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html#cfn-elasticache-serverlesscache-datastorage-minimum
	//
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

