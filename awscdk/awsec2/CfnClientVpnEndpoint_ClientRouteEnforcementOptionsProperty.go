package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientRouteEnforcementOptionsProperty := &ClientRouteEnforcementOptionsProperty{
//   	Enforced: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientrouteenforcementoptions.html
//
type CfnClientVpnEndpoint_ClientRouteEnforcementOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientrouteenforcementoptions.html#cfn-ec2-clientvpnendpoint-clientrouteenforcementoptions-enforced
	//
	Enforced interface{} `field:"optional" json:"enforced" yaml:"enforced"`
}

