package awshealthlake


// The natural language processing (NLP) configuration for the Data Store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   nlpConfigurationProperty := &NlpConfigurationProperty{
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-nlpconfiguration.html
//
type CfnFHIRDatastorePropsMixin_NlpConfigurationProperty struct {
	// The NLP status for the Data Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-nlpconfiguration.html#cfn-healthlake-fhirdatastore-nlpconfiguration-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

