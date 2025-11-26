package previewawselasticachemixins


// The data storage limit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStorageProperty := &DataStorageProperty{
//   	Maximum: jsii.Number(123),
//   	Minimum: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html
//
type CfnServerlessCachePropsMixin_DataStorageProperty struct {
	// The upper limit for data storage the cache is set to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html#cfn-elasticache-serverlesscache-datastorage-maximum
	//
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// The lower limit for data storage the cache is set to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html#cfn-elasticache-serverlesscache-datastorage-minimum
	//
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
	// The unit that the storage is measured in, in GB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-datastorage.html#cfn-elasticache-serverlesscache-datastorage-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

