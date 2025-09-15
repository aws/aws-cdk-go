package awsdocdb


// A reference to a DBClusterParameterGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBClusterParameterGroupReference := &DBClusterParameterGroupReference{
//   	DbClusterParameterGroupId: jsii.String("dbClusterParameterGroupId"),
//   }
//
type DBClusterParameterGroupReference struct {
	// The Id of the DBClusterParameterGroup resource.
	DbClusterParameterGroupId *string `field:"required" json:"dbClusterParameterGroupId" yaml:"dbClusterParameterGroupId"`
}

