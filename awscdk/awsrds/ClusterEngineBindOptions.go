package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
)

// The extra options passed to the `IClusterEngine.bindToCluster` method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterGroup ParameterGroup
//   var roleRef IRoleRef
//
//   clusterEngineBindOptions := &ClusterEngineBindOptions{
//   	ParameterGroup: parameterGroup,
//   	S3ExportRole: roleRef,
//   	S3ImportRole: roleRef,
//   }
//
type ClusterEngineBindOptions struct {
	// The customer-provided ParameterGroup.
	// Default: - none.
	//
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The role used for S3 exporting.
	// Default: - none.
	//
	S3ExportRole interfacesawsiam.IRoleRef `field:"optional" json:"s3ExportRole" yaml:"s3ExportRole"`
	// The role used for S3 importing.
	// Default: - none.
	//
	S3ImportRole interfacesawsiam.IRoleRef `field:"optional" json:"s3ImportRole" yaml:"s3ImportRole"`
}

