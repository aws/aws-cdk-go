package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyInferenceConfigProperty := &clarifyInferenceConfigProperty{
//   	contentTemplate: jsii.String("contentTemplate"),
//   	featureHeaders: []*string{
//   		jsii.String("featureHeaders"),
//   	},
//   	featuresAttribute: jsii.String("featuresAttribute"),
//   	featureTypes: []*string{
//   		jsii.String("featureTypes"),
//   	},
//   	labelAttribute: jsii.String("labelAttribute"),
//   	labelHeaders: []*string{
//   		jsii.String("labelHeaders"),
//   	},
//   	labelIndex: jsii.Number(123),
//   	maxPayloadInMb: jsii.Number(123),
//   	maxRecordCount: jsii.Number(123),
//   	probabilityAttribute: jsii.String("probabilityAttribute"),
//   	probabilityIndex: jsii.Number(123),
//   }
//
type CfnEndpointConfig_ClarifyInferenceConfigProperty struct {
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.ContentTemplate`.
	ContentTemplate *string `field:"optional" json:"contentTemplate" yaml:"contentTemplate"`
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.FeatureHeaders`.
	FeatureHeaders *[]*string `field:"optional" json:"featureHeaders" yaml:"featureHeaders"`
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.FeaturesAttribute`.
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.FeatureTypes`.
	FeatureTypes *[]*string `field:"optional" json:"featureTypes" yaml:"featureTypes"`
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.LabelAttribute`.
	LabelAttribute *string `field:"optional" json:"labelAttribute" yaml:"labelAttribute"`
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.LabelHeaders`.
	LabelHeaders *[]*string `field:"optional" json:"labelHeaders" yaml:"labelHeaders"`
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.LabelIndex`.
	LabelIndex *float64 `field:"optional" json:"labelIndex" yaml:"labelIndex"`
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.MaxPayloadInMB`.
	MaxPayloadInMb *float64 `field:"optional" json:"maxPayloadInMb" yaml:"maxPayloadInMb"`
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.MaxRecordCount`.
	MaxRecordCount *float64 `field:"optional" json:"maxRecordCount" yaml:"maxRecordCount"`
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.ProbabilityAttribute`.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// `CfnEndpointConfig.ClarifyInferenceConfigProperty.ProbabilityIndex`.
	ProbabilityIndex *float64 `field:"optional" json:"probabilityIndex" yaml:"probabilityIndex"`
}

