package awsrds


// The type returned from the `IClusterEngine.bindToCluster` method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterGroup parameterGroup
//
//   clusterEngineConfig := &ClusterEngineConfig{
//   	Features: &ClusterEngineFeatures{
//   		S3Export: jsii.String("s3Export"),
//   		S3Import: jsii.String("s3Import"),
//   	},
//   	ParameterGroup: parameterGroup,
//   	Port: jsii.Number(123),
//   }
//
type ClusterEngineConfig struct {
	// Features supported by the database engine.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html
	//
	// Default: - no features.
	//
	Features *ClusterEngineFeatures `field:"optional" json:"features" yaml:"features"`
	// The ParameterGroup to use for the cluster.
	// Default: - no ParameterGroup will be used.
	//
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The port to use for this cluster, unless the customer specified the port directly.
	// Default: - use the default port for clusters (3306).
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

