package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestStorageAzure(t *testing.T) {
	t.Parallel()

	//for expected := range testCases {
	// Specify the test case folder and "-var" options
	tfOptions := &terraform.Options{
		TerraformDir: "./fixtures/storage-account-name",
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, tfOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, tfOptions)

	// Run `terraform output` to get the values of output variables
	actualTextExample := terraform.Output(t, tfOptions, "storage-name")
	assert.Equal(t, "storagedemopacifico20", actualTextExample)

}
