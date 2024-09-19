package awsmsk


// Properties for defining a `CfnConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationProps := &CfnConfigurationProps{
//   	Name: jsii.String("name"),
//   	ServerProperties: jsii.String("serverProperties"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KafkaVersionsList: []*string{
//   		jsii.String("kafkaVersionsList"),
//   	},
//   	LatestRevision: &LatestRevisionProperty{
//   		CreationTime: jsii.String("creationTime"),
//   		Description: jsii.String("description"),
//   		Revision: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html
//
type CfnConfigurationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html#cfn-msk-configuration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html#cfn-msk-configuration-serverproperties
	//
	ServerProperties *string `field:"required" json:"serverProperties" yaml:"serverProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html#cfn-msk-configuration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html#cfn-msk-configuration-kafkaversionslist
	//
	KafkaVersionsList *[]*string `field:"optional" json:"kafkaVersionsList" yaml:"kafkaVersionsList"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-configuration.html#cfn-msk-configuration-latestrevision
	//
	LatestRevision interface{} `field:"optional" json:"latestRevision" yaml:"latestRevision"`
}

