package awskinesisanalytics


// Describes execution properties for a Managed Service for Apache Flink application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentPropertiesProperty := &EnvironmentPropertiesProperty{
//   	PropertyGroups: []interface{}{
//   		&PropertyGroupProperty{
//   			PropertyGroupId: jsii.String("propertyGroupId"),
//   			PropertyMap: map[string]*string{
//   				"propertyMapKey": jsii.String("propertyMap"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-environmentproperties.html
//
type CfnApplicationV2_EnvironmentPropertiesProperty struct {
	// Describes the execution property groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-environmentproperties.html#cfn-kinesisanalyticsv2-application-environmentproperties-propertygroups
	//
	PropertyGroups interface{} `field:"optional" json:"propertyGroups" yaml:"propertyGroups"`
}

