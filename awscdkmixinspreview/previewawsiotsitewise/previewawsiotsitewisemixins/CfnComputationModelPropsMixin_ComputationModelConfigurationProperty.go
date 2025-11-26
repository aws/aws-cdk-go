package previewawsiotsitewisemixins


// The configuration for the computation model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   computationModelConfigurationProperty := &ComputationModelConfigurationProperty{
//   	AnomalyDetection: &AnomalyDetectionComputationModelConfigurationProperty{
//   		InputProperties: jsii.String("inputProperties"),
//   		ResultProperty: jsii.String("resultProperty"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-computationmodelconfiguration.html
//
type CfnComputationModelPropsMixin_ComputationModelConfigurationProperty struct {
	// The configuration for the anomaly detection type of computation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-computationmodelconfiguration.html#cfn-iotsitewise-computationmodel-computationmodelconfiguration-anomalydetection
	//
	AnomalyDetection interface{} `field:"optional" json:"anomalyDetection" yaml:"anomalyDetection"`
}

