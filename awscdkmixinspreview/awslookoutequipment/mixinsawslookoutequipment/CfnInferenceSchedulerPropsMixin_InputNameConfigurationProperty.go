package mixinsawslookoutequipment


// Specifies configuration information for the input data for the inference, including timestamp format and delimiter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputNameConfigurationProperty := &InputNameConfigurationProperty{
//   	ComponentTimestampDelimiter: jsii.String("componentTimestampDelimiter"),
//   	TimestampFormat: jsii.String("timestampFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-inputnameconfiguration.html
//
type CfnInferenceSchedulerPropsMixin_InputNameConfigurationProperty struct {
	// Indicates the delimiter character used between items in the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-inputnameconfiguration.html#cfn-lookoutequipment-inferencescheduler-inputnameconfiguration-componenttimestampdelimiter
	//
	ComponentTimestampDelimiter *string `field:"optional" json:"componentTimestampDelimiter" yaml:"componentTimestampDelimiter"`
	// The format of the timestamp, whether Epoch time, or standard, with or without hyphens (-).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-inputnameconfiguration.html#cfn-lookoutequipment-inferencescheduler-inputnameconfiguration-timestampformat
	//
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
}

