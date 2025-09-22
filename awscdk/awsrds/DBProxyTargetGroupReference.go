package awsrds


// A reference to a DBProxyTargetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBProxyTargetGroupReference := &DBProxyTargetGroupReference{
//   	TargetGroupArn: jsii.String("targetGroupArn"),
//   }
//
type DBProxyTargetGroupReference struct {
	// The TargetGroupArn of the DBProxyTargetGroup resource.
	TargetGroupArn *string `field:"required" json:"targetGroupArn" yaml:"targetGroupArn"`
}

