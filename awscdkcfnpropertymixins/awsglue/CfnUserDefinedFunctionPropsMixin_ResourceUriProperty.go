package awsglue


// The URIs for function resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   resourceUriProperty := &ResourceUriProperty{
//   	ResourceType: jsii.String("resourceType"),
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-userdefinedfunction-resourceuri.html
//
type CfnUserDefinedFunctionPropsMixin_ResourceUriProperty struct {
	// The type of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-userdefinedfunction-resourceuri.html#cfn-glue-userdefinedfunction-resourceuri-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The URI for accessing the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-userdefinedfunction-resourceuri.html#cfn-glue-userdefinedfunction-resourceuri-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

