package awsemrserverless


// A reference to a Application resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationReference := &ApplicationReference{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	ApplicationId: jsii.String("applicationId"),
//   }
//
type ApplicationReference struct {
	// The ARN of the Application resource.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The ApplicationId of the Application resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
}

