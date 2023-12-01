package awselasticache


// The data storage limit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataStorageProperty := &DataStorageProperty{
//   	Maximum: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html
//
type CfnServerlessCache_DataStorageProperty struct {
	// The upper limit for data storage the cache is set to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html#cfn-elasticache-serverlesscache-datastorage-maximum
	//
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
	// The unit that the storage is measured in, in GB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html#cfn-elasticache-serverlesscache-datastorage-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}

