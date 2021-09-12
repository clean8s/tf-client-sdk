# tf-client-sdk

Creates a Terraform client from a command and/or a RPC port. It's used to access
a provider in a clean way.

```go
import "github.com/clean8s/tf-client-sdk"
func main() {
    c := tfclient.MakeClient(exec.Command("tf-provider-example.exe"), nil)
    c.ReadResource(...)
}
```

Contains code from `terraform/internal`, owned by Hashicorp. Licensed under
MPL. Derivative corresponding to version in`version.go`.

It doesn't completely track upstream and the fork not be automatically synced with
new TF versions until [https://github.com/hashicorp/terraform-plugin-go](https://github.com/hashicorp/terraform-plugin-go)
becomes stable and the plugin protocol (v6?) is finalized.
