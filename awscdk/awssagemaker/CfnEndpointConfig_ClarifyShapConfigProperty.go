package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyShapConfigProperty := &ClarifyShapConfigProperty{
//   	ShapBaselineConfig: &ClarifyShapBaselineConfigProperty{
//   		MimeType: jsii.String("mimeType"),
//   		ShapBaseline: jsii.String("shapBaseline"),
//   		ShapBaselineUri: jsii.String("shapBaselineUri"),
//   	},
//
//   	// the properties below are optional
//   	NumberOfSamples: jsii.Number(123),
//   	Seed: jsii.Number(123),
//   	TextConfig: &ClarifyTextConfigProperty{
//   		Granularity: jsii.String("granularity"),
//   		Language: jsii.String("language"),
//   	},
//   	UseLogit: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html
//
type CfnEndpointConfig_ClarifyShapConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html#cfn-sagemaker-endpointconfig-clarifyshapconfig-shapbaselineconfig
	//
	ShapBaselineConfig interface{} `field:"required" json:"shapBaselineConfig" yaml:"shapBaselineConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html#cfn-sagemaker-endpointconfig-clarifyshapconfig-numberofsamples
	//
	NumberOfSamples *float64 `field:"optional" json:"numberOfSamples" yaml:"numberOfSamples"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html#cfn-sagemaker-endpointconfig-clarifyshapconfig-seed
	//
	Seed *float64 `field:"optional" json:"seed" yaml:"seed"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html#cfn-sagemaker-endpointconfig-clarifyshapconfig-textconfig
	//
	TextConfig interface{} `field:"optional" json:"textConfig" yaml:"textConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html#cfn-sagemaker-endpointconfig-clarifyshapconfig-uselogit
	//
	UseLogit interface{} `field:"optional" json:"useLogit" yaml:"useLogit"`
}

