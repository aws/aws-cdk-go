package awsopsworks


// Properties for defining a `CfnApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppProps := &CfnAppProps{
//   	Name: jsii.String("name"),
//   	StackId: jsii.String("stackId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
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
//   			Value: jsii.String("value"),
//
//   			// the properties below are optional
//   			Secure: jsii.Boolean(false),
//   		},
//   	},
//   	Shortname: jsii.String("shortname"),
//   	SslConfiguration: &SslConfigurationProperty{
//   		Certificate: jsii.String("certificate"),
//   		Chain: jsii.String("chain"),
//   		PrivateKey: jsii.String("privateKey"),
//   	},
//   }
//
type CfnAppProps struct {
	// The app name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The stack ID.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// The app type.
	//
	// Each supported type is associated with a particular layer. For example, PHP applications are associated with a PHP layer. AWS OpsWorks Stacks deploys an application to those instances that are members of the corresponding layer. If your app isn't one of the standard types, or you prefer to implement your own Deploy recipes, specify `other` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// A `Source` object that specifies the app repository.
	AppSource interface{} `field:"optional" json:"appSource" yaml:"appSource"`
	// One or more user-defined key/value pairs to be added to the stack attributes.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The app's data source.
	DataSources interface{} `field:"optional" json:"dataSources" yaml:"dataSources"`
	// A description of the app.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The app virtual host settings, with multiple domains separated by commas.
	//
	// For example: `'www.example.com, example.com'`
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Whether to enable SSL for the app.
	EnableSsl interface{} `field:"optional" json:"enableSsl" yaml:"enableSsl"`
	// An array of `EnvironmentVariable` objects that specify environment variables to be associated with the app.
	//
	// After you deploy the app, these variables are defined on the associated app server instance. For more information, see [Environment Variables](https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html#workingapps-creating-environment) .
	//
	// There is no specific limit on the number of environment variables. However, the size of the associated data structure - which includes the variables' names, values, and protected flag values - cannot exceed 20 KB. This limit should accommodate most if not all use cases. Exceeding it will cause an exception with the message, "Environment: is too large (maximum is 20KB)."
	//
	// > If you have specified one or more environment variables, you cannot modify the stack's Chef version.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The app's short name.
	Shortname *string `field:"optional" json:"shortname" yaml:"shortname"`
	// An `SslConfiguration` object with the SSL configuration.
	SslConfiguration interface{} `field:"optional" json:"sslConfiguration" yaml:"sslConfiguration"`
}

