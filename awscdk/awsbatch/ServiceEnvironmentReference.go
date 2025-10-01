package awsbatch


// A reference to a ServiceEnvironment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceEnvironmentReference := &ServiceEnvironmentReference{
//   	ServiceEnvironmentArn: jsii.String("serviceEnvironmentArn"),
//   }
//
type ServiceEnvironmentReference struct {
	// The ServiceEnvironmentArn of the ServiceEnvironment resource.
	ServiceEnvironmentArn *string `field:"required" json:"serviceEnvironmentArn" yaml:"serviceEnvironmentArn"`
}

