package awsroute53globalresolver


// Properties for CfnHostedZoneAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnHostedZoneAssociationMixinProps := &CfnHostedZoneAssociationMixinProps{
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	Name: jsii.String("name"),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-hostedzoneassociation.html
//
type CfnHostedZoneAssociationMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-hostedzoneassociation.html#cfn-route53globalresolver-hostedzoneassociation-hostedzoneid
	//
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-hostedzoneassociation.html#cfn-route53globalresolver-hostedzoneassociation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-hostedzoneassociation.html#cfn-route53globalresolver-hostedzoneassociation-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

