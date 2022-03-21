# Optional, only if exists

variable "config" {

  # type = any

  type = object({
    someArg1 = string
    # Configure your parameters
  })
}

variable "dependency" {

  # type = object({
  #   from_module = any
  # })

  type = object({
    from_module = object({
      someArg1 = string
      # Configure your parameters
    })
  })
}


# Expected single output

output "cfout" {
  # whatever value required by supporting connected modules
  value = {
    someOutput = "Happy days"
  }
}