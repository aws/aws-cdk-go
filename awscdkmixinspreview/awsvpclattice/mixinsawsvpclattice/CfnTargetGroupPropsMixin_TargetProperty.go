package mixinsawsvpclattice


// Describes a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetProperty := &TargetProperty{
//   	Id: jsii.String("id"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-target.html
//
type CfnTargetGroupPropsMixin_TargetProperty struct {
	// The ID of the target.
	//
	// If the target group type is `INSTANCE` , this is an instance ID. If the target group type is `IP` , this is an IP address. If the target group type is `LAMBDA` , this is the ARN of a Lambda function. If the target group type is `ALB` , this is the ARN of an Application Load Balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-target.html#cfn-vpclattice-targetgroup-target-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The port on which the target is listening.
	//
	// For HTTP, the default is 80. For HTTPS, the default is 443.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-target.html#cfn-vpclattice-targetgroup-target-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

