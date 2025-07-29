package awsglue


// Information about a serialization/deserialization program (SerDe) that serves as an extractor and loader.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   serdeInfoProperty := &SerdeInfoProperty{
//   	Name: jsii.String("name"),
//   	Parameters: parameters,
//   	SerializationLibrary: jsii.String("serializationLibrary"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-serdeinfo.html
//
type CfnPartition_SerdeInfoProperty struct {
	// Name of the SerDe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-serdeinfo.html#cfn-glue-partition-serdeinfo-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// These key-value pairs define initialization parameters for the SerDe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-serdeinfo.html#cfn-glue-partition-serdeinfo-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Usually the class that implements the SerDe.
	//
	// An example is `org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-serdeinfo.html#cfn-glue-partition-serdeinfo-serializationlibrary
	//
	SerializationLibrary *string `field:"optional" json:"serializationLibrary" yaml:"serializationLibrary"`
}

