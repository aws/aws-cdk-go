package awsgameliftstreams


// Properties for defining a `CfnStreamGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamGroupProps := &CfnStreamGroupProps{
//   	Description: jsii.String("description"),
//   	LocationConfigurations: []interface{}{
//   		&LocationConfigurationProperty{
//   			LocationName: jsii.String("locationName"),
//
//   			// the properties below are optional
//   			AlwaysOnCapacity: jsii.Number(123),
//   			OnDemandCapacity: jsii.Number(123),
//   		},
//   	},
//   	StreamClass: jsii.String("streamClass"),
//
//   	// the properties below are optional
//   	DefaultApplication: &DefaultApplicationProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html
//
type CfnStreamGroupProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html#cfn-gameliftstreams-streamgroup-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html#cfn-gameliftstreams-streamgroup-locationconfigurations
	//
	LocationConfigurations interface{} `field:"required" json:"locationConfigurations" yaml:"locationConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html#cfn-gameliftstreams-streamgroup-streamclass
	//
	StreamClass *string `field:"required" json:"streamClass" yaml:"streamClass"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html#cfn-gameliftstreams-streamgroup-defaultapplication
	//
	DefaultApplication interface{} `field:"optional" json:"defaultApplication" yaml:"defaultApplication"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html#cfn-gameliftstreams-streamgroup-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

