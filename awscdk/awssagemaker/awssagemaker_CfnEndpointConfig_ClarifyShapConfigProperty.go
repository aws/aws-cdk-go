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

