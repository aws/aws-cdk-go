// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha


// Attributes for the App Runner Service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   serviceAttributes := &serviceAttributes{
//   	serviceArn: jsii.String("serviceArn"),
//   	serviceName: jsii.String("serviceName"),
//   	serviceStatus: jsii.String("serviceStatus"),
//   	serviceUrl: jsii.String("serviceUrl"),
//   }
//
// Experimental.
type ServiceAttributes struct {
	// The ARN of the service.
	// Experimental.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// The name of the service.
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The status of the service.
	// Experimental.
	ServiceStatus *string `field:"required" json:"serviceStatus" yaml:"serviceStatus"`
	// The URL of the service.
	// Experimental.
	ServiceUrl *string `field:"required" json:"serviceUrl" yaml:"serviceUrl"`
}

