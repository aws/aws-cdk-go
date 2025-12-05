package previewawsdevopsagentmixins


// AWS resource definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var resourceMetadata interface{}
//
//   aWSResourceProperty := &AWSResourceProperty{
//   	ResourceArn: jsii.String("resourceArn"),
//   	ResourceMetadata: resourceMetadata,
//   	ResourceType: jsii.String("resourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsresource.html
//
type CfnAssociationPropsMixin_AWSResourceProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsresource.html#cfn-devopsagent-association-awsresource-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// Additional metadata for the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsresource.html#cfn-devopsagent-association-awsresource-resourcemetadata
	//
	ResourceMetadata interface{} `field:"optional" json:"resourceMetadata" yaml:"resourceMetadata"`
	// Resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsresource.html#cfn-devopsagent-association-awsresource-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

