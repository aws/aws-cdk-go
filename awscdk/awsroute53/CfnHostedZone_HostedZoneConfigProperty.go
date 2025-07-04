package awsroute53


// A complex type that contains an optional comment about your hosted zone.
//
// If you don't want to specify a comment, omit both the `HostedZoneConfig` and `Comment` elements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostedZoneConfigProperty := &HostedZoneConfigProperty{
//   	Comment: jsii.String("comment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html
//
type CfnHostedZone_HostedZoneConfigProperty struct {
	// Any comments that you want to include about the hosted zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html#cfn-route53-hostedzone-hostedzoneconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

