package awsmediatailor


// Configuration for custom output functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customOutputConfigurationProperty := &CustomOutputConfigurationProperty{
//   	Runtime: jsii.String("runtime"),
//
//   	// the properties below are optional
//   	Output: map[string]*string{
//   		"outputKey": jsii.String("output"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-customoutputconfiguration.html
//
type CfnFunction_CustomOutputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-customoutputconfiguration.html#cfn-mediatailor-function-customoutputconfiguration-runtime
	//
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-customoutputconfiguration.html#cfn-mediatailor-function-customoutputconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
}

