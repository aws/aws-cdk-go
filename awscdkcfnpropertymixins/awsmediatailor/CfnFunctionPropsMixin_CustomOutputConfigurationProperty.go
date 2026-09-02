package awsmediatailor


// Configuration for custom output functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   customOutputConfigurationProperty := &CustomOutputConfigurationProperty{
//   	Output: map[string]*string{
//   		"outputKey": jsii.String("output"),
//   	},
//   	Runtime: jsii.String("runtime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-customoutputconfiguration.html
//
type CfnFunctionPropsMixin_CustomOutputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-customoutputconfiguration.html#cfn-mediatailor-function-customoutputconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-customoutputconfiguration.html#cfn-mediatailor-function-customoutputconfiguration-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
}

