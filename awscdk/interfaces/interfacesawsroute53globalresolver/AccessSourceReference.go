package interfacesawsroute53globalresolver


// A reference to a AccessSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessSourceReference := &AccessSourceReference{
//   	AccessSourceArn: jsii.String("accessSourceArn"),
//   	AccessSourceId: jsii.String("accessSourceId"),
//   }
//
type AccessSourceReference struct {
	// The ARN of the AccessSource resource.
	AccessSourceArn *string `field:"required" json:"accessSourceArn" yaml:"accessSourceArn"`
	// The AccessSourceId of the AccessSource resource.
	AccessSourceId *string `field:"required" json:"accessSourceId" yaml:"accessSourceId"`
}

