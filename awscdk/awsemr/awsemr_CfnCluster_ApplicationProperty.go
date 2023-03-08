package awsemr


// `Application` is a property of `AWS::EMR::Cluster` .
//
// The `Application` property type defines the open-source big data applications for EMR to install and configure when a cluster is created.
//
// With Amazon EMR release version 4.0 and later, the only accepted parameter is the application `Name` . To pass arguments to these applications, you use configuration classifications specified using JSON objects in a `Configuration` property. For more information, see [Configuring Applications](https://docs.aws.amazon.com//emr/latest/ReleaseGuide/emr-configure-apps.html) .
//
// With earlier Amazon EMR releases, the application is any AWS or third-party software that you can add to the cluster. You can specify the version of the application and arguments to pass to it. Amazon EMR accepts and forwards the argument list to the corresponding installation script as a bootstrap action argument.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationProperty := &applicationProperty{
//   	additionalInfo: map[string]*string{
//   		"additionalInfoKey": jsii.String("additionalInfo"),
//   	},
//   	args: []*string{
//   		jsii.String("args"),
//   	},
//   	name: jsii.String("name"),
//   	version: jsii.String("version"),
//   }
//
type CfnCluster_ApplicationProperty struct {
	// This option is for advanced users only.
	//
	// This is meta information about clusters and applications that are used for testing and troubleshooting.
	AdditionalInfo interface{} `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// Arguments for Amazon EMR to pass to the application.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The name of the application.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The version of the application.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

