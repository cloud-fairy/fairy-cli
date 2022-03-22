# Cloudfairy cli 🧰

CLI helper to create cloud-fairy modules and connectors.

# Installation


[Go to release page](../../releases)
Download release for your OS and architecture.

# Usage
```
# Generate a component
fairy generate component [username]/[module-name]

# Generate a connector
fairy generate connector [username]/[module-name]
```

# Cloudfairy Terraform API
Expect your terraform module/connector to receive the following inputs of object types
- config
- dependency

Provide a single output named "cfout".

Additional outputs are optional, but cloud-fairy terragrunt files collect only **cfout** data as part of a cloudfairy project.

Your dependencies will be populated according to the **cloudfairy.yaml** descriptor file.

# Documentation WIP
