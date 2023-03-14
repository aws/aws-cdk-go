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
//   hostedZoneConfigProperty := &hostedZoneConfigProperty{
//   	comment: jsii.String("comment"),
//   }
//
type CfnHostedZone_HostedZoneConfigProperty struct {
	// Any comments that you want to include about the hosted zone.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

