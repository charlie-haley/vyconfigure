# VyConfigure

Declarative YAML configuration for VyOS

__Note: this project is far from production ready, use at your own risk!__

## Installation

You will need to enable the HTTP API on your VyOS instance, [refer to the upstream documentation for how to configure it.](https://docs.vyos.io/en/latest/configuration/service/https.html)

[The latest binary is available in releases](https://github.com/charlie-haley/vyconfigure/releases). [There's also a Docker Image available on GHCR.](https://github.com/charlie-haley/vyconfigure/pkgs/container/vyconfigure)

## Workflow
You should start by syncing your existing configuration to your local filesystem so you can begin using VyConfigure.
```bash
# This will sync your existing VyOS config to your current working directory
vyconfigure --host="https://<VyOS IP or Hostname>" --api-key="<VyOS HTTP API key>" sync
```

Once the configuration is on your local filesystem, you can preview the changes using
```bash
vyconfigure --host="https://<VyOS IP or Hostname>" --api-key="<VyOS HTTP API key>" plan
```

If you're happy with the changes, then you can apply them.
```bash
vyconfigure --host="https://<VyOS IP or Hostname>" --api-key="<VyOS HTTP API key>" apply
```

## How does VyConfigure work?
VyConfigure works by using [the VyOS HTTP API](https://docs.vyos.io/en/latest/configuration/service/https.html). It translates the configuration into YAML files and then back to a set of commands when you apply.

## What's the purpose of VyConfigure?
I created VyConfigure so I could manage my VyOS config the GitOps way. I found tools like Ansible too heavy and wanted something lightweight and simple that could be easily read.

## Unsupported features
Currently, configuring users with vyconfigure is explicity blocked due to complexities around encrypted passwords, for now it's recommended you configure these as usual.

[Please raise an issue](https://github.com/charlie-haley/vyconfigure/issues) for any issues or proposed features. Contributions also welcome 😊
