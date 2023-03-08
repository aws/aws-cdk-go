package awskinesisanalytics


// Describes execution properties for a Flink-based Kinesis Data Analytics application.
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
type CfnApplicationV2_EnvironmentPropertiesProperty struct {
	// Describes the execution property groups.
	PropertyGroups interface{} `field:"optional" json:"propertyGroups" yaml:"propertyGroups"`
}

