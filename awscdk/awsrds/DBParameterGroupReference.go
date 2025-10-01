package awsrds


// A reference to a DBParameterGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBParameterGroupReference := &DBParameterGroupReference{
//   	DbParameterGroupName: jsii.String("dbParameterGroupName"),
//   }
//
type DBParameterGroupReference struct {
	// The DBParameterGroupName of the DBParameterGroup resource.
	DbParameterGroupName *string `field:"required" json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
}

