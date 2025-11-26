package previewawsemrmixins


// `KeyValue` is a subproperty of the `HadoopJarStepConfig` property type.
//
// `KeyValue` is used to pass parameters to a step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyValueProperty := &KeyValueProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-keyvalue.html
//
type CfnClusterPropsMixin_KeyValueProperty struct {
	// The unique identifier of a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-keyvalue.html#cfn-emr-cluster-keyvalue-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value part of the identified key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-keyvalue.html#cfn-emr-cluster-keyvalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

