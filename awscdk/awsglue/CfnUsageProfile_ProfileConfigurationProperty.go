package awsglue


// Specifies the job and session values that an admin configures in an AWS Glue usage profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profileConfigurationProperty := &ProfileConfigurationProperty{
//   	JobConfiguration: map[string]interface{}{
//   		"jobConfigurationKey": &ConfigurationObjectProperty{
//   			"allowedValues": []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			"defaultValue": jsii.String("defaultValue"),
//   			"maxValue": jsii.String("maxValue"),
//   			"minValue": jsii.String("minValue"),
//   		},
//   	},
//   	SessionConfiguration: map[string]interface{}{
//   		"sessionConfigurationKey": &ConfigurationObjectProperty{
//   			"allowedValues": []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			"defaultValue": jsii.String("defaultValue"),
//   			"maxValue": jsii.String("maxValue"),
//   			"minValue": jsii.String("minValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-usageprofile-profileconfiguration.html
//
type CfnUsageProfile_ProfileConfigurationProperty struct {
	// A key-value map of configuration parameters for AWS Glue jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-usageprofile-profileconfiguration.html#cfn-glue-usageprofile-profileconfiguration-jobconfiguration
	//
	JobConfiguration interface{} `field:"optional" json:"jobConfiguration" yaml:"jobConfiguration"`
	// A key-value map of configuration parameters for AWS Glue sessions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-usageprofile-profileconfiguration.html#cfn-glue-usageprofile-profileconfiguration-sessionconfiguration
	//
	SessionConfiguration interface{} `field:"optional" json:"sessionConfiguration" yaml:"sessionConfiguration"`
}

