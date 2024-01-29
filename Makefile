.PHONY: build

build:
	sam build

clean-deploy:
	rm -rf .aws-sam/cache
	sam build
	sam deploy --no-confirm-changeset