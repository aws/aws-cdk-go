package awsvpclattice


// Describes a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetProperty := &TargetProperty{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	Port: jsii.Number(123),
//   }
//
type CfnTargetGroup_TargetProperty struct {
	// The ID of the target.
	//
	// If the target type of the target group is `INSTANCE` , this is an instance ID. If the target type is `IP` , this is an IP address. If the target type is `LAMBDA` , this is the ARN of the Lambda function. If the target type is `ALB` , this is the ARN of the Application Load Balancer.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The port on which the target is listening.
	//
	// For HTTP, the default is `80` . For HTTPS, the default is `443` .
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

