package awsopensearchservice


// Options for configuring service software updates for a domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   softwareUpdateOptionsProperty := &SoftwareUpdateOptionsProperty{
//   	AutoSoftwareUpdateEnabled: jsii.Boolean(false),
//   }
//
type CfnDomain_SoftwareUpdateOptionsProperty struct {
	// Specifies whether automatic service software updates are enabled for the domain.
	AutoSoftwareUpdateEnabled interface{} `field:"optional" json:"autoSoftwareUpdateEnabled" yaml:"autoSoftwareUpdateEnabled"`
}

