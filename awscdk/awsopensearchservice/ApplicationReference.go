package awsopensearchservice


// A reference to a Application resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationReference := &ApplicationReference{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	ApplicationName: jsii.String("applicationName"),
//   }
//
type ApplicationReference struct {
	// The ARN of the Application resource.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The Name of the Application resource.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
}

