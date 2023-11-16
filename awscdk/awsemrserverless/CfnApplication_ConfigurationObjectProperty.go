package awsemrserverless


// Configuration for a JobRun.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationObjectProperty_ configurationObjectProperty
//
//   configurationObjectProperty := &configurationObjectProperty{
//   	Classification: jsii.String("classification"),
//
//   	// the properties below are optional
//   	Configurations: []interface{}{
//   		configurationObjectProperty_,
//   	},
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html
//
type CfnApplication_ConfigurationObjectProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html#cfn-emrserverless-application-configurationobject-classification
	//
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html#cfn-emrserverless-application-configurationobject-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html#cfn-emrserverless-application-configurationobject-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

