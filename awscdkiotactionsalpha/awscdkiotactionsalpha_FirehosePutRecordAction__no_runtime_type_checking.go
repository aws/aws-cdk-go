//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateNewFirehosePutRecordActionParameters(stream awscdkkinesisfirehosealpha.IDeliveryStream, props *FirehosePutRecordActionProps) error {
	return nil
}

