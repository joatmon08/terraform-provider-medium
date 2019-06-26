build:
	go build -o terraform-provider-medium
	mkdir -p ~/.terraform.d/plugins
	cp terraform-provider-medium ~/.terraform.d/plugins/

demo:
	cd example && terraform init
	cd example && terraform apply

clean:
	cd example && terraform destroy --force
	cd example && rm -rf .terraform
	cd example && rm -f terraform.tfstate*
	rm -f terraform-provider-medium