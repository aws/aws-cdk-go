package awsiotsitewise


// Contains the configuration for anomaly detection computation models.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anomalyDetectionComputationModelConfigurationProperty := &AnomalyDetectionComputationModelConfigurationProperty{
//   	InputProperties: jsii.String("inputProperties"),
//   	ResultProperty: jsii.String("resultProperty"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-anomalydetectioncomputationmodelconfiguration.html
//
type CfnComputationModel_AnomalyDetectionComputationModelConfigurationProperty struct {
	// The list of input properties for the anomaly detection model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-anomalydetectioncomputationmodelconfiguration.html#cfn-iotsitewise-computationmodel-anomalydetectioncomputationmodelconfiguration-inputproperties
	//
	InputProperties *string `field:"required" json:"inputProperties" yaml:"inputProperties"`
	// The property where the anomaly detection results will be stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-anomalydetectioncomputationmodelconfiguration.html#cfn-iotsitewise-computationmodel-anomalydetectioncomputationmodelconfiguration-resultproperty
	//
	ResultProperty *string `field:"required" json:"resultProperty" yaml:"resultProperty"`
}

