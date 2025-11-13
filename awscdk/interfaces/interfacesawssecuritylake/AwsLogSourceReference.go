package interfacesawssecuritylake


// A reference to a AwsLogSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsLogSourceReference := &AwsLogSourceReference{
//   	SourceName: jsii.String("sourceName"),
//   	SourceVersion: jsii.String("sourceVersion"),
//   }
//
type AwsLogSourceReference struct {
	// The SourceName of the AwsLogSource resource.
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// The SourceVersion of the AwsLogSource resource.
	SourceVersion *string `field:"required" json:"sourceVersion" yaml:"sourceVersion"`
}

