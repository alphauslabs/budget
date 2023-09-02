`tucp` is the command line client for the **TrueUnblended Control Plane** service [(tucpd)](https://github.com/mobingilabs/ouchan/tree/master/cloudrun/tucpd).

To install using [HomeBrew](https://brew.sh/), run the following command:

```bash
$ brew install alphauslabs/tap/tucp
```

To setup authentication, set your `GOOGLE_APPLICATION_CREDENTIALS` env variable using your credentials file. You also need to give your credentials file access to the `tucpd-[next|prod]` service. To do so, try the following commands:

```bash
# Install the `iam` tool:
$ brew install alphauslabs/tap/iam

# Validate `iam` credentials:
$ iam whoami

# Request access to our `tucpd-[next|prod]` service (once only):
$ iam allow-me tucpd-prod
```

Explore more available subcommands and flags though:

```bash
$ tucp -h
# or
$ tucp <subcmd> -h
```
