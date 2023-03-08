package awseks


// The cluster control plane logging configuration for your cluster.
//
// > When updating a resource, you must include this `ClusterLogging` property if the previous CloudFormation template of the resource had it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterLoggingProperty := &clusterLoggingProperty{
//   	enabledTypes: []interface{}{
//   		&loggingTypeConfigProperty{
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnCluster_ClusterLoggingProperty struct {
	// The enabled control plane logs for your cluster. All log types are disabled if the array is empty.
	//
	// > When updating a resource, you must include this `EnabledTypes` property if the previous CloudFormation template of the resource had it.
	EnabledTypes interface{} `field:"optional" json:"enabledTypes" yaml:"enabledTypes"`
}

