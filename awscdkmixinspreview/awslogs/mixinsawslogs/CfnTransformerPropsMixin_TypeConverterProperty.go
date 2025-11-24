package mixinsawslogs


// Use this processor to convert a value type associated with the specified key to the specified type.
//
// It's a casting processor that changes the types of the specified fields. Values can be converted into one of the following datatypes: `integer` , `double` , `string` and `boolean` .
//
// For more information about this processor including examples, see [trimString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-trimString) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   typeConverterProperty := &TypeConverterProperty{
//   	Entries: []interface{}{
//   		&TypeConverterEntryProperty{
//   			Key: jsii.String("key"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-typeconverter.html
//
type CfnTransformerPropsMixin_TypeConverterProperty struct {
	// An array of `TypeConverterEntry` objects, where each object contains the information about one field to change the type of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-typeconverter.html#cfn-logs-transformer-typeconverter-entries
	//
	Entries interface{} `field:"optional" json:"entries" yaml:"entries"`
}

