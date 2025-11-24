package mixinsawsemrserverless


// The configuration to use to enable the different types of interactive use cases in an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   interactiveConfigurationProperty := &InteractiveConfigurationProperty{
//   	LivyEndpointEnabled: jsii.Boolean(false),
//   	StudioEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-interactiveconfiguration.html
//
type CfnApplicationPropsMixin_InteractiveConfigurationProperty struct {
	// Enables an Apache Livy endpoint that you can connect to and run interactive jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-interactiveconfiguration.html#cfn-emrserverless-application-interactiveconfiguration-livyendpointenabled
	//
	// Default: - false.
	//
	LivyEndpointEnabled interface{} `field:"optional" json:"livyEndpointEnabled" yaml:"livyEndpointEnabled"`
	// Enables you to connect an application to Amazon EMR Studio to run interactive workloads in a notebook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-interactiveconfiguration.html#cfn-emrserverless-application-interactiveconfiguration-studioenabled
	//
	// Default: - false.
	//
	StudioEnabled interface{} `field:"optional" json:"studioEnabled" yaml:"studioEnabled"`
}

