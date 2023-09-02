`budget` is the command line client for our internal budget service [(budgetd)](https://github.com/mobingilabs/ouchan/tree/master/cloudrun/budgetd).

To install using [HomeBrew](https://brew.sh/), run the following command:

```bash
$ brew install alphauslabs/tap/budget
```

To setup authentication, set your `GOOGLE_APPLICATION_CREDENTIALS` env variable using your credentials file.

Explore more available subcommands and flags though:

```bash
$ budget -h
# or
$ budget <subcmd> -h
```
