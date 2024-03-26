package awsiotsitewise


// Represents one level between a composite model and the root of the asset model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propertyPathDefinitionProperty := &PropertyPathDefinitionProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-propertypathdefinition.html
//
type CfnAssetModel_PropertyPathDefinitionProperty struct {
	// The name of the path segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-propertypathdefinition.html#cfn-iotsitewise-assetmodel-propertypathdefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

