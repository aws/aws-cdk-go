package previewawsquicksightmixins


// The static data configuration of the reference line data configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceLineStaticDataConfigurationProperty := &ReferenceLineStaticDataConfigurationProperty{
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinestaticdataconfiguration.html
//
type CfnTemplatePropsMixin_ReferenceLineStaticDataConfigurationProperty struct {
	// The double input of the static data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinestaticdataconfiguration.html#cfn-quicksight-template-referencelinestaticdataconfiguration-value
	//
	// Default: - 0.
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

