package awsiotwireless


// A reference to a ServiceProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceProfileReference := &ServiceProfileReference{
//   	ServiceProfileArn: jsii.String("serviceProfileArn"),
//   	ServiceProfileId: jsii.String("serviceProfileId"),
//   }
//
type ServiceProfileReference struct {
	// The ARN of the ServiceProfile resource.
	ServiceProfileArn *string `field:"required" json:"serviceProfileArn" yaml:"serviceProfileArn"`
	// The Id of the ServiceProfile resource.
	ServiceProfileId *string `field:"required" json:"serviceProfileId" yaml:"serviceProfileId"`
}

