package awstimestream


// The cluster configuration of the resource at the time of backup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   clusterConfigurationProperty := &ClusterConfigurationProperty{
//   	DedicatedCompactor: jsii.Boolean(false),
//   	IngestQueryInstances: jsii.Number(123),
//   	QueryOnlyInstances: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbbackup-clusterconfiguration.html
//
type CfnInfluxDBBackupPropsMixin_ClusterConfigurationProperty struct {
	// Indicates if the compactor instance is a standalone instance or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbbackup-clusterconfiguration.html#cfn-timestream-influxdbbackup-clusterconfiguration-dedicatedcompactor
	//
	DedicatedCompactor interface{} `field:"optional" json:"dedicatedCompactor" yaml:"dedicatedCompactor"`
	// The number of ingest and query instances in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbbackup-clusterconfiguration.html#cfn-timestream-influxdbbackup-clusterconfiguration-ingestqueryinstances
	//
	IngestQueryInstances *float64 `field:"optional" json:"ingestQueryInstances" yaml:"ingestQueryInstances"`
	// The number of query-only instances in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbbackup-clusterconfiguration.html#cfn-timestream-influxdbbackup-clusterconfiguration-queryonlyinstances
	//
	QueryOnlyInstances *float64 `field:"optional" json:"queryOnlyInstances" yaml:"queryOnlyInstances"`
}

