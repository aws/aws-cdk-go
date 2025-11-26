package previewawsbatchmixins


// The FireLens configuration for the container.
//
// This is used to specify and configure a log router for container logs. For more information, see [Custom log](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) routing in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   firelensConfigurationProperty := &FirelensConfigurationProperty{
//   	Options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-firelensconfiguration.html
//
type CfnJobDefinitionPropsMixin_FirelensConfigurationProperty struct {
	// The options to use when configuring the log router.
	//
	// This field is optional and can be used to specify a custom configuration file or to add additional metadata, such as the task, task definition, cluster, and container instance details to the log event. If specified, the syntax to use is `"options":{"enable-ecs-log-metadata":"true|false","config-file-type:"s3|file","config-file-value":"arn:aws:s3:::mybucket/fluent.conf|filepath"}` . For more information, see [Creating a task definition that uses a FireLens configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-firelensconfiguration.html#cfn-batch-jobdefinition-firelensconfiguration-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The log router to use.
	//
	// The valid values are `fluentd` or `fluentbit` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-firelensconfiguration.html#cfn-batch-jobdefinition-firelensconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

