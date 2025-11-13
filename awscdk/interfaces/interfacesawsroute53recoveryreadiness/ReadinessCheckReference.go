package interfacesawsroute53recoveryreadiness


// A reference to a ReadinessCheck resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   readinessCheckReference := &ReadinessCheckReference{
//   	ReadinessCheckArn: jsii.String("readinessCheckArn"),
//   	ReadinessCheckName: jsii.String("readinessCheckName"),
//   }
//
type ReadinessCheckReference struct {
	// The ARN of the ReadinessCheck resource.
	ReadinessCheckArn *string `field:"required" json:"readinessCheckArn" yaml:"readinessCheckArn"`
	// The ReadinessCheckName of the ReadinessCheck resource.
	ReadinessCheckName *string `field:"required" json:"readinessCheckName" yaml:"readinessCheckName"`
}

