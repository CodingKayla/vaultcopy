# vaultcopy

I had a need for local copies of live vault secrets with certain keys set to different values for testing. I wrote this to copy paths and their K/V pairs and override given K/V pairs while retaining the values of the rest.

## Notes
* This only works with version 1 of the K/V secret engine
* If your app does token renewal you may need to create a policy and generate a renewable token.
  * `path "secret/*" {
  capabilities = ["read"]
}`
* Don't mix up the local and external addresses.

## Setup
1. Start a local development vault server
  * `vault server -dev -dev-root-token-id="derp"`
2. Set local variables for vault command
  * `export VAULT_ADDR=http://127.0.0.1:8200`
  * `export VAULT_TOKEN=derp`
3. Enable v1 of the engine
  * `vault secrets disable secret`
  * `vault secrets enable -path=secret -version=1 kv`
4. Set local and external vault variables
  * `export LOCAL_ADDR=http://127.0.0.1:8200`
  * `export LOCAL_TOKEN=derp`
  * `export EXTERNAL_ADDR=https://your.live.vault.address:8200`
  * `export EXTERNAL_TOKEN=<external token>`
5. Setup paths.json
  * Refer to the paths.json in the repo for an example
6. Run it!
  * `go run main.go`
