package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyShapConfigProperty := &clarifyShapConfigProperty{
//   	shapBaselineConfig: &clarifyShapBaselineConfigProperty{
//   		mimeType: jsii.String("mimeType"),
//   		shapBaseline: jsii.String("shapBaseline"),
//   		shapBaselineUri: jsii.String("shapBaselineUri"),
//   	},
//
//   	// the properties below are optional
//   	numberOfSamples: jsii.Number(123),
//   	seed: jsii.Number(123),
//   	textConfig: &clarifyTextConfigProperty{
//   		granularity: jsii.String("granularity"),
//   		language: jsii.String("language"),
//   	},
//   	useLogit: jsii.Boolean(false),
//   }
//
type CfnEndpointConfig_ClarifyShapConfigProperty struct {
	// `CfnEndpointConfig.ClarifyShapConfigProperty.ShapBaselineConfig`.
	ShapBaselineConfig interface{} `field:"required" json:"shapBaselineConfig" yaml:"shapBaselineConfig"`
	// `CfnEndpointConfig.ClarifyShapConfigProperty.NumberOfSamples`.
	NumberOfSamples *float64 `field:"optional" json:"numberOfSamples" yaml:"numberOfSamples"`
	// `CfnEndpointConfig.ClarifyShapConfigProperty.Seed`.
	Seed *float64 `field:"optional" json:"seed" yaml:"seed"`
	// `CfnEndpointConfig.ClarifyShapConfigProperty.TextConfig`.
	TextConfig interface{} `field:"optional" json:"textConfig" yaml:"textConfig"`
	// `CfnEndpointConfig.ClarifyShapConfigProperty.UseLogit`.
	UseLogit interface{} `field:"optional" json:"useLogit" yaml:"useLogit"`
}

