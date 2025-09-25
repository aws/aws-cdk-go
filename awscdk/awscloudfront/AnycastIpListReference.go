package awscloudfront


// A reference to a AnycastIpList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anycastIpListReference := &AnycastIpListReference{
//   	AnycastIpListId: jsii.String("anycastIpListId"),
//   }
//
type AnycastIpListReference struct {
	// The Id of the AnycastIpList resource.
	AnycastIpListId *string `field:"required" json:"anycastIpListId" yaml:"anycastIpListId"`
}

