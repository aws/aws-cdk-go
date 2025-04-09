package awsevidently


// A set of key-value pairs that specify users who should always be served a specific variation of a feature.
//
// Each key specifies a user using their user ID, account ID, or some other identifier. The value specifies the name of the variation that the user is to be served.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityOverrideProperty := &EntityOverrideProperty{
//   	EntityId: jsii.String("entityId"),
//   	Variation: jsii.String("variation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-feature-entityoverride.html
//
type CfnFeature_EntityOverrideProperty struct {
	// The entity ID to be served the variation specified in `Variation` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-feature-entityoverride.html#cfn-evidently-feature-entityoverride-entityid
	//
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// The name of the variation to serve to the user session that matches the `EntityId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-feature-entityoverride.html#cfn-evidently-feature-entityoverride-variation
	//
	Variation *string `field:"optional" json:"variation" yaml:"variation"`
}

