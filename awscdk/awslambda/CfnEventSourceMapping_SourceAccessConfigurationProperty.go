package awslambda


// The configuration used by AWS Lambda to access event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceAccessConfigurationProperty := &SourceAccessConfigurationProperty{
//   	Type: jsii.String("type"),
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html
//
type CfnEventSourceMapping_SourceAccessConfigurationProperty struct {
	// The type of source access configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html#cfn-lambda-eventsourcemapping-sourceaccessconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The URI for the source access configuration resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html#cfn-lambda-eventsourcemapping-sourceaccessconfiguration-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

