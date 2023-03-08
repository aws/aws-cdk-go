package awsroute53resolver


// Properties for defining a `CfnResolverConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverConfigProps := &cfnResolverConfigProps{
//   	autodefinedReverseFlag: jsii.String("autodefinedReverseFlag"),
//   	resourceId: jsii.String("resourceId"),
//   }
//
type CfnResolverConfigProps struct {
	// Represents the desired status of `AutodefinedReverse` .
	//
	// The only supported value on creation is `DISABLE` . Deletion of this resource will return `AutodefinedReverse` to its default value of `ENABLED` .
	AutodefinedReverseFlag *string `field:"required" json:"autodefinedReverseFlag" yaml:"autodefinedReverseFlag"`
	// The ID of the Amazon Virtual Private Cloud VPC that you're configuring Resolver for.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

