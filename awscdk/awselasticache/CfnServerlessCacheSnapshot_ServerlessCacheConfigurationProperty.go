package awselasticache


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessCacheConfigurationProperty := &ServerlessCacheConfigurationProperty{
//   	Engine: jsii.String("engine"),
//   	MajorEngineVersion: jsii.String("majorEngineVersion"),
//   	ServerlessCacheName: jsii.String("serverlessCacheName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscachesnapshot-serverlesscacheconfiguration.html
//
type CfnServerlessCacheSnapshot_ServerlessCacheConfigurationProperty struct {
	// The engine that the serverless cache is configured with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscachesnapshot-serverlesscacheconfiguration.html#cfn-elasticache-serverlesscachesnapshot-serverlesscacheconfiguration-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The engine version number that the serverless cache is configured with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscachesnapshot-serverlesscacheconfiguration.html#cfn-elasticache-serverlesscachesnapshot-serverlesscacheconfiguration-majorengineversion
	//
	MajorEngineVersion *string `field:"optional" json:"majorEngineVersion" yaml:"majorEngineVersion"`
	// The identifier of the serverless cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscachesnapshot-serverlesscacheconfiguration.html#cfn-elasticache-serverlesscachesnapshot-serverlesscacheconfiguration-serverlesscachename
	//
	ServerlessCacheName *string `field:"optional" json:"serverlessCacheName" yaml:"serverlessCacheName"`
}

