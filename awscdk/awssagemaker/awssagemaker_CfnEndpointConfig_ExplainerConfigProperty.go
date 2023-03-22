package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   explainerConfigProperty := &explainerConfigProperty{
//   	clarifyExplainerConfig: &clarifyExplainerConfigProperty{
//   		shapConfig: &clarifyShapConfigProperty{
//   			shapBaselineConfig: &clarifyShapBaselineConfigProperty{
//   				mimeType: jsii.String("mimeType"),
//   				shapBaseline: jsii.String("shapBaseline"),
//   				shapBaselineUri: jsii.String("shapBaselineUri"),
//   			},
//
//   			// the properties below are optional
//   			numberOfSamples: jsii.Number(123),
//   			seed: jsii.Number(123),
//   			textConfig: &clarifyTextConfigProperty{
//   				granularity: jsii.String("granularity"),
//   				language: jsii.String("language"),
//   			},
//   			useLogit: jsii.Boolean(false),
//   		},
//
//   		// the properties below are optional
//   		enableExplanations: jsii.String("enableExplanations"),
//   		inferenceConfig: &clarifyInferenceConfigProperty{
//   			contentTemplate: jsii.String("contentTemplate"),
//   			featureHeaders: []*string{
//   				jsii.String("featureHeaders"),
//   			},
//   			featuresAttribute: jsii.String("featuresAttribute"),
//   			featureTypes: []*string{
//   				jsii.String("featureTypes"),
//   			},
//   			labelAttribute: jsii.String("labelAttribute"),
//   			labelHeaders: []*string{
//   				jsii.String("labelHeaders"),
//   			},
//   			labelIndex: jsii.Number(123),
//   			maxPayloadInMb: jsii.Number(123),
//   			maxRecordCount: jsii.Number(123),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			probabilityIndex: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnEndpointConfig_ExplainerConfigProperty struct {
	// `CfnEndpointConfig.ExplainerConfigProperty.ClarifyExplainerConfig`.
	ClarifyExplainerConfig interface{} `field:"optional" json:"clarifyExplainerConfig" yaml:"clarifyExplainerConfig"`
}

