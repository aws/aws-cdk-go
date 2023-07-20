package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyInferenceConfigProperty := &ClarifyInferenceConfigProperty{
//   	ContentTemplate: jsii.String("contentTemplate"),
//   	FeatureHeaders: []*string{
//   		jsii.String("featureHeaders"),
//   	},
//   	FeaturesAttribute: jsii.String("featuresAttribute"),
//   	FeatureTypes: []*string{
//   		jsii.String("featureTypes"),
//   	},
//   	LabelAttribute: jsii.String("labelAttribute"),
//   	LabelHeaders: []*string{
//   		jsii.String("labelHeaders"),
//   	},
//   	LabelIndex: jsii.Number(123),
//   	MaxPayloadInMb: jsii.Number(123),
//   	MaxRecordCount: jsii.Number(123),
//   	ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   	ProbabilityIndex: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html
//
type CfnEndpointConfig_ClarifyInferenceConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-contenttemplate
	//
	ContentTemplate *string `field:"optional" json:"contentTemplate" yaml:"contentTemplate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-featureheaders
	//
	FeatureHeaders *[]*string `field:"optional" json:"featureHeaders" yaml:"featureHeaders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-featuresattribute
	//
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-featuretypes
	//
	FeatureTypes *[]*string `field:"optional" json:"featureTypes" yaml:"featureTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-labelattribute
	//
	LabelAttribute *string `field:"optional" json:"labelAttribute" yaml:"labelAttribute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-labelheaders
	//
	LabelHeaders *[]*string `field:"optional" json:"labelHeaders" yaml:"labelHeaders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-labelindex
	//
	LabelIndex *float64 `field:"optional" json:"labelIndex" yaml:"labelIndex"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-maxpayloadinmb
	//
	MaxPayloadInMb *float64 `field:"optional" json:"maxPayloadInMb" yaml:"maxPayloadInMb"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-maxrecordcount
	//
	MaxRecordCount *float64 `field:"optional" json:"maxRecordCount" yaml:"maxRecordCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-probabilityattribute
	//
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.html#cfn-sagemaker-endpointconfig-clarifyinferenceconfig-probabilityindex
	//
	ProbabilityIndex *float64 `field:"optional" json:"probabilityIndex" yaml:"probabilityIndex"`
}

