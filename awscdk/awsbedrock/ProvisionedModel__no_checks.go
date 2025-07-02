//go:build no_runtime_type_checking

package awsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateProvisionedModel_FromProvisionedModelArnParameters(_scope constructs.Construct, _id *string, provisionedModelArn *string) error {
	return nil
}

