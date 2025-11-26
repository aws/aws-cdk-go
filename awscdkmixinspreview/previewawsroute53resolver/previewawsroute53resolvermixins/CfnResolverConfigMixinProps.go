package previewawsroute53resolvermixins


// Properties for CfnResolverConfigPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResolverConfigMixinProps := &CfnResolverConfigMixinProps{
//   	AutodefinedReverseFlag: jsii.String("autodefinedReverseFlag"),
//   	ResourceId: jsii.String("resourceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverconfig.html
//
type CfnResolverConfigMixinProps struct {
	// Represents the desired status of `AutodefinedReverse` .
	//
	// The only supported value on creation is `DISABLE` . Deletion of this resource will return `AutodefinedReverse` to its default value of `ENABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverconfig.html#cfn-route53resolver-resolverconfig-autodefinedreverseflag
	//
	AutodefinedReverseFlag *string `field:"optional" json:"autodefinedReverseFlag" yaml:"autodefinedReverseFlag"`
	// The ID of the Amazon Virtual Private Cloud VPC or a Route 53 Profile that you're configuring Resolver for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverconfig.html#cfn-route53resolver-resolverconfig-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

