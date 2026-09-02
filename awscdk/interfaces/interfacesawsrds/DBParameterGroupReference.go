package interfacesawsrds


// A reference to a DBParameterGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBParameterGroupReference := &DBParameterGroupReference{
//   	DbParameterGroupArn: jsii.String("dbParameterGroupArn"),
//   	DbParameterGroupName: jsii.String("dbParameterGroupName"),
//   }
//
type DBParameterGroupReference struct {
	// The ARN of the DBParameterGroup resource.
	DbParameterGroupArn *string `field:"required" json:"dbParameterGroupArn" yaml:"dbParameterGroupArn"`
	// The DBParameterGroupName of the DBParameterGroup resource.
	DbParameterGroupName *string `field:"required" json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
}

