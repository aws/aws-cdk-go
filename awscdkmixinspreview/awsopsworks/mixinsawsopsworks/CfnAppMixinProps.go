package mixinsawsopsworks


// Properties for CfnAppPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppMixinProps := &CfnAppMixinProps{
//   	AppSource: &SourceProperty{
//   		Password: jsii.String("password"),
//   		Revision: jsii.String("revision"),
//   		SshKey: jsii.String("sshKey"),
//   		Type: jsii.String("type"),
//   		Url: jsii.String("url"),
//   		Username: jsii.String("username"),
//   	},
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	DataSources: []interface{}{
//   		&DataSourceProperty{
//   			Arn: jsii.String("arn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	EnableSsl: jsii.Boolean(false),
//   	Environment: []interface{}{
//   		&EnvironmentVariableProperty{
//   			Key: jsii.String("key"),
//   			Secure: jsii.Boolean(false),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Shortname: jsii.String("shortname"),
//   	SslConfiguration: &SslConfigurationProperty{
//   		Certificate: jsii.String("certificate"),
//   		Chain: jsii.String("chain"),
//   		PrivateKey: jsii.String("privateKey"),
//   	},
//   	StackId: jsii.String("stackId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html
//
type CfnAppMixinProps struct {
	// A `Source` object that specifies the app repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-appsource
	//
	AppSource interface{} `field:"optional" json:"appSource" yaml:"appSource"`
	// One or more user-defined key/value pairs to be added to the stack attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The app's data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-datasources
	//
	DataSources interface{} `field:"optional" json:"dataSources" yaml:"dataSources"`
	// A description of the app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The app virtual host settings, with multiple domains separated by commas.
	//
	// For example: `'www.example.com, example.com'`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-domains
	//
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Whether to enable SSL for the app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-enablessl
	//
	EnableSsl interface{} `field:"optional" json:"enableSsl" yaml:"enableSsl"`
	// An array of `EnvironmentVariable` objects that specify environment variables to be associated with the app.
	//
	// After you deploy the app, these variables are defined on the associated app server instance. For more information, see [Environment Variables](https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html#workingapps-creating-environment) .
	//
	// There is no specific limit on the number of environment variables. However, the size of the associated data structure - which includes the variables' names, values, and protected flag values - cannot exceed 20 KB. This limit should accommodate most if not all use cases. Exceeding it will cause an exception with the message, "Environment: is too large (maximum is 20KB)."
	//
	// > If you have specified one or more environment variables, you cannot modify the stack's Chef version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The app name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The app's short name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-shortname
	//
	Shortname *string `field:"optional" json:"shortname" yaml:"shortname"`
	// An `SslConfiguration` object with the SSL configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-sslconfiguration
	//
	SslConfiguration interface{} `field:"optional" json:"sslConfiguration" yaml:"sslConfiguration"`
	// The stack ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-stackid
	//
	StackId *string `field:"optional" json:"stackId" yaml:"stackId"`
	// The app type.
	//
	// Each supported type is associated with a particular layer. For example, PHP applications are associated with a PHP layer. OpsWorks Stacks deploys an application to those instances that are members of the corresponding layer. If your app isn't one of the standard types, or you prefer to implement your own Deploy recipes, specify `other` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

