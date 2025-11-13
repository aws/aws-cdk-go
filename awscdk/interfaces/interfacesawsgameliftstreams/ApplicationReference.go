package interfacesawsgameliftstreams


// A reference to a Application resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationReference := &ApplicationReference{
//   	ApplicationArn: jsii.String("applicationArn"),
//   }
//
type ApplicationReference struct {
	// The Arn of the Application resource.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
}

