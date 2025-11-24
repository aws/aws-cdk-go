package mixinsawsodb


// Information about the data collection options enabled for a VM cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataCollectionOptionsProperty := &DataCollectionOptionsProperty{
//   	IsDiagnosticsEventsEnabled: jsii.Boolean(false),
//   	IsHealthMonitoringEnabled: jsii.Boolean(false),
//   	IsIncidentLogsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-datacollectionoptions.html
//
type CfnCloudVmClusterPropsMixin_DataCollectionOptionsProperty struct {
	// Specifies whether diagnostic collection is enabled for the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-datacollectionoptions.html#cfn-odb-cloudvmcluster-datacollectionoptions-isdiagnosticseventsenabled
	//
	IsDiagnosticsEventsEnabled interface{} `field:"optional" json:"isDiagnosticsEventsEnabled" yaml:"isDiagnosticsEventsEnabled"`
	// Specifies whether health monitoring is enabled for the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-datacollectionoptions.html#cfn-odb-cloudvmcluster-datacollectionoptions-ishealthmonitoringenabled
	//
	IsHealthMonitoringEnabled interface{} `field:"optional" json:"isHealthMonitoringEnabled" yaml:"isHealthMonitoringEnabled"`
	// Specifies whether incident logs are enabled for the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-datacollectionoptions.html#cfn-odb-cloudvmcluster-datacollectionoptions-isincidentlogsenabled
	//
	IsIncidentLogsEnabled interface{} `field:"optional" json:"isIncidentLogsEnabled" yaml:"isIncidentLogsEnabled"`
}

