# eth-watcher

> Watch EVM compatible chain's native token transfer & ERC 20 Token Transfer events.

## /!\ Warning

This project is still in very early stage, it might be unstable and / or might have major changes in future releases. Use at your own risk.

## Get started

1. Copy `config.yml.example` and rename to `config.yml` , edit details as you wish.
2. Run `docker-compose pull` to pull prebuilt images, you can also build your own version. Feel free to modify!
3. Run `docker-compose up -d` to start.

## Webhook

### Request info

- Method: `POST`
- Content Type: `application/json`

### Structure

Native token transfer events:

```json
{
  "ts": "2024-01-14T21:36:12+08:00",
  "chain_id": 11155111,
  "sender": "0xD3E8ce4841ed658Ec8dcb99B7a74beFC377253EA",
  "receiver": "0x9C8a0A9B5d5b178D73e775a2dC4D52711758C388",
  "is_native": true,
  "amount": 0.0001,
  "tx": "0xfbf2b6d0b8ad824e5abc7db2a9852a0f63fadafd13fd31147b82c5dcec2aa48d"
}
```

ERC20 token transfer events:

```json
{
  "ts": "2024-01-14T21:51:00+08:00",
  "chain_id": 11155111,
  "sender": "0xD3E8ce4841ed658Ec8dcb99B7a74beFC377253EA",
  "receiver": "0x9C8a0A9B5d5b178D73e775a2dC4D52711758C388",
  "is_native": false,
  "contract": {
    "address": "0xcb7729f2B44Ae7B86D58Bb8068f0EAD8fcF9378c",
    "name": "TestERC20Coin",
    "symbol": "TEC",
    "decimals": 18
  },
  "amount": 1,
  "tx": "0xcff756b6dd2e58ffed5860c1872f50c1e630d59cce0a4a430332277089fd4184"
}
```

## Troubleshooting

1. Change `system.production` to false in config file.
2. Restart service, read logs and see what's wrong.
3. Or just add breakpoints and run debugger. Done right!

## Config Demo

The config powering [NyaOne Sponsor Bot](https://nya.one/@sponsor) looks like below:

```yaml
system:
  redis: "redis://redis:6379/0"
  production: true
receiver:
  - "0x9C8a0A9B5d5b178D73e775a2dC4D52711758C388"
chain:
  - id: 1 # Mainnet
    rpc: "https://ethereum.publicnode.com"
    interval: 30s
    includeNative: true
    includeERC20: true
  - id: 137 # Polygon
    rpc: "https://polygon.llamarpc.com"
    interval: 10s
    includeNative: true
    includeERC20: true
  - id: 56 # BSC
    rpc: "https://bscrpc.com"
    interval: 10s
    includeNative: true
    includeERC20: true
webhooks:
  - "<REDACTED>"
```
