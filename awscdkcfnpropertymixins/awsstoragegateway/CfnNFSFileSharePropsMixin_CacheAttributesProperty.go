package awsstoragegateway


// Refresh cache information for the file share.
//
// Removing this property from a template leaves the share's current value: the service documents no value that resets it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cacheAttributesProperty := &CacheAttributesProperty{
//   	CacheStaleTimeoutInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-storagegateway-nfsfileshare-cacheattributes.html
//
type CfnNFSFileSharePropsMixin_CacheAttributesProperty struct {
	// How long the directory contents cached by the gateway are considered valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-storagegateway-nfsfileshare-cacheattributes.html#cfn-storagegateway-nfsfileshare-cacheattributes-cachestaletimeoutinseconds
	//
	CacheStaleTimeoutInSeconds *float64 `field:"optional" json:"cacheStaleTimeoutInSeconds" yaml:"cacheStaleTimeoutInSeconds"`
}

