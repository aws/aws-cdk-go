//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validateStreamDeliveryResource_KinesisParameters(stream awskinesis.IStream, options *KinesisStreamDeliveryOptions) error {
	return nil
}

