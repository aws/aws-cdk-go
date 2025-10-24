package awseventschemas


// Properties for defining a `CfnRegistry`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRegistryProps := &CfnRegistryProps{
//   	Description: jsii.String("description"),
//   	RegistryName: jsii.String("registryName"),
//   	Tags: []TagsEntryProperty{
//   		&TagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-registry.html
//
type CfnRegistryProps struct {
	// A description of the registry to be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-registry.html#cfn-eventschemas-registry-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the schema registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-registry.html#cfn-eventschemas-registry-registryname
	//
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// Tags to associate with the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-registry.html#cfn-eventschemas-registry-tags
	//
	Tags *[]*CfnRegistry_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

