package awsqbusiness


// A reference to a DataAccessor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataAccessorReference := &DataAccessorReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	DataAccessorArn: jsii.String("dataAccessorArn"),
//   	DataAccessorId: jsii.String("dataAccessorId"),
//   }
//
type DataAccessorReference struct {
	// The ApplicationId of the DataAccessor resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ARN of the DataAccessor resource.
	DataAccessorArn *string `field:"required" json:"dataAccessorArn" yaml:"dataAccessorArn"`
	// The DataAccessorId of the DataAccessor resource.
	DataAccessorId *string `field:"required" json:"dataAccessorId" yaml:"dataAccessorId"`
}

