package awsecs


// Configuration returned by AlternateTargetConfiguration.bind().
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alternateTargetConfig := &AlternateTargetConfig{
//   	AlternateTargetGroupArn: jsii.String("alternateTargetGroupArn"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	ProductionListenerRule: jsii.String("productionListenerRule"),
//   	TestListenerRule: jsii.String("testListenerRule"),
//   }
//
type AlternateTargetConfig struct {
	// The ARN of the alternate target group.
	AlternateTargetGroupArn *string `field:"required" json:"alternateTargetGroupArn" yaml:"alternateTargetGroupArn"`
	// The IAM role ARN for the configuration.
	// Default: - a new role will be created.
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The production listener rule ARN (ALB) or listener ARN (NLB).
	// Default: - none.
	//
	ProductionListenerRule *string `field:"optional" json:"productionListenerRule" yaml:"productionListenerRule"`
	// The test listener rule ARN (ALB) or listener ARN (NLB).
	// Default: - none.
	//
	TestListenerRule *string `field:"optional" json:"testListenerRule" yaml:"testListenerRule"`
}

