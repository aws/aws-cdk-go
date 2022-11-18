package awsrum


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDestinationProperty := &metricDestinationProperty{
//   	destination: jsii.String("destination"),
//
//   	// the properties below are optional
//   	destinationArn: jsii.String("destinationArn"),
//   	iamRoleArn: jsii.String("iamRoleArn"),
//   	metricDefinitions: []interface{}{
//   		&metricDefinitionProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			dimensionKeys: map[string]*string{
//   				"dimensionKeysKey": jsii.String("dimensionKeys"),
//   			},
//   			eventPattern: jsii.String("eventPattern"),
//   			unitLabel: jsii.String("unitLabel"),
//   			valueKey: jsii.String("valueKey"),
//   		},
//   	},
//   }
//
type CfnAppMonitor_MetricDestinationProperty struct {
	// `CfnAppMonitor.MetricDestinationProperty.Destination`.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// `CfnAppMonitor.MetricDestinationProperty.DestinationArn`.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// `CfnAppMonitor.MetricDestinationProperty.IamRoleArn`.
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// `CfnAppMonitor.MetricDestinationProperty.MetricDefinitions`.
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
}

