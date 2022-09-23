package awsrds


// The type returned from the {@link IClusterEngine.bindToCluster} method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterGroup parameterGroup
//
//   clusterEngineConfig := &clusterEngineConfig{
//   	features: &clusterEngineFeatures{
//   		s3Export: jsii.String("s3Export"),
//   		s3Import: jsii.String("s3Import"),
//   	},
//   	parameterGroup: parameterGroup,
//   	port: jsii.Number(123),
//   }
//
type ClusterEngineConfig struct {
	// Features supported by the database engine.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html
	//
	Features *ClusterEngineFeatures `field:"optional" json:"features" yaml:"features"`
	// The ParameterGroup to use for the cluster.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The port to use for this cluster, unless the customer specified the port directly.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

