# aws-creds-check

Attempts to get AWS caller identity.

## Install

```
#> go get github.com/jimrazmus/aws-creds-check
```

## Run

```
#> aws-creds-check
```

## Note

Set the AWS_SDK_LOAD_CONFIG environment variable to 1 if you use an aws config file.

```
By default NewSession will only load credentials from the shared credentials file (~/.aws/credentials). If the AWS_SDK_LOAD_CONFIG environment variable is set to a truthy value the Session will be created from the configuration values from the shared config (~/.aws/config) and shared credentials (~/.aws/credentials) files.
```

[Sessions options from Shared Config](https://docs.aws.amazon.com/sdk-for-go/api/aws/session/)
