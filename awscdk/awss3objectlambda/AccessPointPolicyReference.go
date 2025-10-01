package awss3objectlambda


// A reference to a AccessPointPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPointPolicyReference := &AccessPointPolicyReference{
//   	ObjectLambdaAccessPoint: jsii.String("objectLambdaAccessPoint"),
//   }
//
type AccessPointPolicyReference struct {
	// The ObjectLambdaAccessPoint of the AccessPointPolicy resource.
	ObjectLambdaAccessPoint *string `field:"required" json:"objectLambdaAccessPoint" yaml:"objectLambdaAccessPoint"`
}

