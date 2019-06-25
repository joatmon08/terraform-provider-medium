build:
	go build -o terraform-provider-medium
	terraform init

clean:
	terraform destroy --force
	rm -rf .terraform
	rm -f terraform-provider-medium
	rm -f terraform.tfstate*