package mixinsawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAnomalyDetectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnomalyDetectorMixinProps := &CfnAnomalyDetectorMixinProps{
//   	Alias: jsii.String("alias"),
//   	Configuration: &AnomalyDetectorConfigurationProperty{
//   		RandomCutForest: &RandomCutForestConfigurationProperty{
//   			IgnoreNearExpectedFromAbove: &IgnoreNearExpectedProperty{
//   				Amount: jsii.Number(123),
//   				Ratio: jsii.Number(123),
//   			},
//   			IgnoreNearExpectedFromBelow: &IgnoreNearExpectedProperty{
//   				Amount: jsii.Number(123),
//   				Ratio: jsii.Number(123),
//   			},
//   			Query: jsii.String("query"),
//   			SampleSize: jsii.Number(123),
//   			ShingleSize: jsii.Number(123),
//   		},
//   	},
//   	EvaluationIntervalInSeconds: jsii.Number(123),
//   	Labels: []interface{}{
//   		&LabelProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MissingDataAction: &MissingDataActionProperty{
//   		MarkAsAnomaly: jsii.Boolean(false),
//   		Skip: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Workspace: jsii.String("workspace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html
//
type CfnAnomalyDetectorMixinProps struct {
	// The user-friendly name of the anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html#cfn-aps-anomalydetector-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The algorithm configuration of the anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html#cfn-aps-anomalydetector-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The frequency, in seconds, at which the anomaly detector evaluates metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html#cfn-aps-anomalydetector-evaluationintervalinseconds
	//
	// Default: - 60.
	//
	EvaluationIntervalInSeconds *float64 `field:"optional" json:"evaluationIntervalInSeconds" yaml:"evaluationIntervalInSeconds"`
	// The Amazon Managed Service for Prometheus metric labels associated with the anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html#cfn-aps-anomalydetector-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// The action taken when data is missing during evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html#cfn-aps-anomalydetector-missingdataaction
	//
	MissingDataAction interface{} `field:"optional" json:"missingDataAction" yaml:"missingDataAction"`
	// The tags applied to the anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html#cfn-aps-anomalydetector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An Amazon Managed Service for Prometheus workspace is a logical and isolated Prometheus server dedicated to ingesting, storing, and querying your Prometheus-compatible metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html#cfn-aps-anomalydetector-workspace
	//
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

