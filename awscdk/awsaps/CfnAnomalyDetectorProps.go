package awsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAnomalyDetector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalyDetectorProps := &CfnAnomalyDetectorProps{
//   	Alias: jsii.String("alias"),
//   	Configuration: &AnomalyDetectorConfigurationProperty{
//   		RandomCutForest: &RandomCutForestConfigurationProperty{
//   			Query: jsii.String("query"),
//
//   			// the properties below are optional
//   			IgnoreNearExpectedFromAbove: &IgnoreNearExpectedProperty{
//   				Amount: jsii.Number(123),
//   				Ratio: jsii.Number(123),
//   			},
//   			IgnoreNearExpectedFromBelow: &IgnoreNearExpectedProperty{
//   				Amount: jsii.Number(123),
//   				Ratio: jsii.Number(123),
//   			},
//   			SampleSize: jsii.Number(123),
//   			ShingleSize: jsii.Number(123),
//   		},
//   	},
//   	Workspace: jsii.String("workspace"),
//
//   	// the properties below are optional
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html
//
type CfnAnomalyDetectorProps struct {
	// The user-friendly name of the anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html#cfn-aps-anomalydetector-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The algorithm configuration of the anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html#cfn-aps-anomalydetector-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// An Amazon Managed Service for Prometheus workspace is a logical and isolated Prometheus server dedicated to ingesting, storing, and querying your Prometheus-compatible metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-anomalydetector.html#cfn-aps-anomalydetector-workspace
	//
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
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
}

