# learn-crud

## Requirements

1. Go 1.14
2. Access to https://github.com/tiket.
3. `ntaps`

This project has dependency to other repository in github.com/tiket (a private repository), so we have to make sure
that `go get` can access that repository either via ssh or https.

- Via SSH:
    1. Add your public key as ssh-key in
       github. (https://help.github.com/en/github/authenticating-to-github/adding-a-new-ssh-key-to-your-github-account)
    2. Force git to use ssh instead of https when access github.com
   ```
   git config --global url."git@github.com:".insteadOf "https://github.com/"
   ```
- Via HTTPS:
    1. Set credentials to access github.com via https (https://git-scm.com/docs/gitcredentials).

## How to start in development

### Clone Project

1. Clone project either using ssh or https to your preferred directory.

```bash
git clone git@github.com:tiket/learn-crud.git
```

2. Move to project directory.

```bash
cd learn-crud
```

### Set Environment Variable

Create your .env file from .env.dist and adjust the values ​​for each variable.

```bash
cp .env.dist .env
```

when .env file exists on the same binary folder, then it will load automatically or if you want to load custom env file
or properties then you can execute with

```bash
CONFIG_FILE=/path/to/config/file ntaps run
```

if no env file or properties supplied then it will use default config below

```bash
EXAMPLE_CONFIG="some value"
```

### Run Program

1. Run unit test

```bash
ntaps test
```

2. Build binary

```bash
ntaps build
```

3. Execute binary

```bash
ntaps run
```

For more commands please check https://github.com/tiket/ntaps/blob/master/README.md.

# Project Layout

- **application** directory contains entity, repository and bussiness logic
- **config** directory contains configuration for your application and it must have default values for local
  development. it accepts value from env variable and will override default value.
- **di** contains all dependecy injection configuration, all component must be registered here so it can be used by
  other components
- **infrastructure** contains code to the outside world
- **interfaces** glued for application layer
- **shared** shared code and functionality
- **shared/dto** put dto here
- **mocks** contains mocked file

## how to generate mockgen

```bash
ntaps mock
```

## good

1. write code that easy to remove, not easy to maintain
2. code through interface (SOLID)
3. all single unit of business code lives in application directory
4. infrastructure job :
    - bind req to dto
    - dto validation
    - authentication and authorization
    - calling interfaces view services
    - send response
5. application handles business logic
6. repo handles persistent store (either db, file, cache etc ...)
7. entity represent something unique
8. dto bind request into model (model is not unique)
9. keep data structure simple
