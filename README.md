# Terraform Mint Provider (Deprecated)

> [!WARNING]
> **This provider is deprecated and no longer maintained.**
> It has been superseded by the **RWX provider**, published as
> [`rwx-cloud/rwx`](https://registry.terraform.io/providers/rwx-cloud/rwx) on the
> Terraform Registry. Please migrate at your earliest convenience — this provider
> will receive no further releases, bug fixes, or dependency updates.

The Mint Provider enabled [Terraform](https://terraform.io) to manage
[RWX](https://www.rwx.com) resources (vault secrets and variables). Its
functionality now lives in the RWX provider.

## Where it moved

| | Mint (deprecated) | RWX (current) |
| --- | --- | --- |
| Registry | [`rwx-research/mint`](https://registry.terraform.io/providers/rwx-research/mint) | [`rwx-cloud/rwx`](https://registry.terraform.io/providers/rwx-cloud/rwx) |
| Source | `rwx-research/terraform-provider-mint` | [`rwx-cloud/terraform-provider-rwx`](https://github.com/rwx-cloud/terraform-provider-rwx) |
| Provider source address | `rwx-research/mint` | `rwx-cloud/rwx` |
| Resources | `mint_secret`, `mint_variable` | `rwx_secret`, `rwx_variable` |

## Migrating

The resource schemas are identical, so you can migrate existing Terraform state
in place — no resources are destroyed or recreated. In short:

1. Update `required_providers` to `source = "rwx-cloud/rwx"`, rename the
   `provider "mint"` block to `provider "rwx"`, and rename `mint_*` resource
   types to `rwx_*`.
2. `terraform init`
3. Re-point state and rename the resource types:
   ```sh
   terraform state replace-provider \
     registry.terraform.io/rwx-research/mint \
     registry.terraform.io/rwx-cloud/rwx
   terraform state mv mint_secret.example   rwx_secret.example
   terraform state mv mint_variable.example rwx_variable.example
   ```
4. `terraform plan` — it must show **no changes** before you apply.

**Full step-by-step guide (including `count`/`for_each` and caveats):**
https://registry.terraform.io/providers/rwx-cloud/rwx/latest/docs/guides/migrating-from-mint

## Questions

Reach out at [support@rwx.com](mailto:support@rwx.com).
