package awsrds


// A reference to a DBClusterParameterGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBClusterParameterGroupReference := &DBClusterParameterGroupReference{
//   	DbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   }
//
type DBClusterParameterGroupReference struct {
	// The DBClusterParameterGroupName of the DBClusterParameterGroup resource.
	DbClusterParameterGroupName *string `field:"required" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
}

