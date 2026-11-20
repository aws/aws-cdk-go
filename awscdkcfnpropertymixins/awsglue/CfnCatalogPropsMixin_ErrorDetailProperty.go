package awsglue


// Details about an error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   errorDetailProperty := &ErrorDetailProperty{
//   	ErrorCode: jsii.String("errorCode"),
//   	ErrorMessage: jsii.String("errorMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-errordetail.html
//
type CfnCatalogPropsMixin_ErrorDetailProperty struct {
	// The code associated with this error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-errordetail.html#cfn-glue-catalog-errordetail-errorcode
	//
	ErrorCode *string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// A message describing the error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-errordetail.html#cfn-glue-catalog-errordetail-errormessage
	//
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
}

