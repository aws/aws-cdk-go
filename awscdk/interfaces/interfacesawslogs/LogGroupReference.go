package interfacesawslogs


// A reference to a LogGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logGroupReference := &LogGroupReference{
//   	LogGroupArn: jsii.String("logGroupArn"),
//   	LogGroupName: jsii.String("logGroupName"),
//   }
//
type LogGroupReference struct {
	// The ARN of the LogGroup resource.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
	// The LogGroupName of the LogGroup resource.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

